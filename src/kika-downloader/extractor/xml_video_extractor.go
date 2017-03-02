package extractor

import (
	"fmt"
	"gopkg.in/xmlpath.v2"
	"kika-downloader/contract"
	"kika-downloader/http"
	"kika-downloader/model"
	"kika-downloader/utils"
	"kika-downloader/vo"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type xmlVideoExtractor struct {
	httpClient http.ClientInterface

	xPathVideoPageVideoTags  string
	xPathVideoPageXmlDataTag string

	xPathXmlSeriesTitle         string
	xPathXmlEpisodeTitle        string
	xPathXmlEpisodeLanguageCode string
	xPathXmlEpisodeDescription  string
	xPathXmlAssets              string

	regExpVideoId    string
	regExpXmlDataUrl string
}

type videoAsset struct {
	Width    int
	Height   int
	FileSize int64
	VideoURL *url.URL
}

// NewXmlVideoExtractor return xml meta data extractor
func NewXmlVideoExtractor(
	httpClient http.ClientInterface,

	xPathVideoPageVideoTags string,
	xPathVideoPageXmlDataTag string,

	xPathXmlSeriesTitle string,
	xPathXmlEpisodeTitle string,
	xPathXmlEpisodeLanguageCode string,
	xPathXmlEpisodeDescription string,
	xPathXmlAssets string,

	regExpVideoId string,
	regExpXmlDataUrl string,
) contract.VideoExtractorInterface {
	return &xmlVideoExtractor{
		httpClient: httpClient,

		xPathVideoPageVideoTags:  xPathVideoPageVideoTags,
		xPathVideoPageXmlDataTag: xPathVideoPageXmlDataTag,

		xPathXmlSeriesTitle:         xPathXmlSeriesTitle,
		xPathXmlEpisodeTitle:        xPathXmlEpisodeTitle,
		xPathXmlEpisodeLanguageCode: xPathXmlEpisodeLanguageCode,
		xPathXmlEpisodeDescription:  xPathXmlEpisodeDescription,
		xPathXmlAssets:              xPathXmlAssets,

		regExpVideoId:    regExpVideoId,
		regExpXmlDataUrl: regExpXmlDataUrl,
	}
}

// ExtractVideoFromURL extract video meta data from url
func (e *xmlVideoExtractor) ExtractVideoFromURL(rawURL string) (contract.VideoInterface, error) {
	_, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	videoPageResponse, err := e.httpClient.Get(rawURL)
	if err != nil {
		return nil, err
	}

	rootDoc, err := utils.NodeFromHtmlBody(videoPageResponse.Body)
	if err != nil {
		return nil, err
	}

	xmlElementIdId, err := e.findIdOfXmlElement(rootDoc)
	if err != nil {
		return nil, err
	}

	xmlUrl, err := e.findXmlDataUrl(rootDoc, xmlElementIdId)
	if err != nil {
		return nil, err
	}

	xmlResponse, err := e.httpClient.Get(xmlUrl.String())
	if err != nil {
		return nil, err
	}

	xmlRoot, err := xmlpath.Parse(xmlResponse.Body)
	if err != nil {
		return nil, err
	}

	// Because of the leading 'p', xmlElementIdId is not a valid UUID, so
	// we remove it and validate it.
	videoId := xmlElementIdId[1:]

	if !utils.IsUUID4(videoId) {
		return nil, fmt.Errorf("invalid video id (%s)", videoId)
	}

	return e.makeVideoFromXmlRoot(videoId, xmlRoot)
}

func (e *xmlVideoExtractor) makeVideoFromXmlRoot(videoId string, xml *xmlpath.Node) (contract.VideoInterface, error) {
	seriesTitlePath, err := xmlpath.Compile(e.xPathXmlSeriesTitle)
	if err != nil {
		return nil, err
	}

	fullEpisodeTitlePath, err := xmlpath.Compile(e.xPathXmlEpisodeTitle)
	if err != nil {
		return nil, err
	}

	episodeLanguageCodePath, err := xmlpath.Compile(e.xPathXmlEpisodeLanguageCode)
	if err != nil {
		return nil, err
	}

	episodeDescriptionPath, err := xmlpath.Compile(e.xPathXmlEpisodeDescription)
	if err != nil {
		return nil, err
	}

	seriesTitle, ok := seriesTitlePath.String(xml)
	if !ok {
		return nil, fmt.Errorf("no series title found in xml document")
	}

	fullEpisodesTitle, ok := fullEpisodeTitlePath.String(xml)
	if !ok {
		return nil, fmt.Errorf("no episodes title found in xml document")
	}

	episodeLanguageCode, ok := episodeLanguageCodePath.String(xml)
	if !ok {
		return nil, fmt.Errorf("no language code found in xml document")
	}

	episodeDescription, ok := episodeDescriptionPath.String(xml)
	if !ok {
		return nil, fmt.Errorf("no episode description found in xml document")
	}

	// split full episode title into number and text
	episodeNo, episodeTitle, err := e.splitFullEpisodeTitle(fullEpisodesTitle)
	if err != nil {
		return nil, err
	}

	// extract real video
	asset, err := e.extractVideoInformation(xml)
	if err != nil {
		return nil, err
	}

	// assemble video
	video := model.NewVideo(
		videoId,
		seriesTitle,
		episodeTitle,
		episodeDescription,
		episodeLanguageCode,
		episodeNo,
		vo.NewResolution(asset.Width, asset.Height),
		asset.VideoURL,
		asset.FileSize,
	)

	return video, nil
}

// return width, height, origin url, error
func (e *xmlVideoExtractor) extractVideoInformation(doc *xmlpath.Node) (*videoAsset, error) {
	assetsPath, err := xmlpath.Compile(e.xPathXmlAssets)
	if err != nil {
		return nil, err
	}

	assetsIter := assetsPath.Iter(doc)

	lastPixels := 0

	pos := 0
	var bestAsset *videoAsset

	for assetsIter.Next() {
		pos++
		assetNode := assetsIter.Node()

		mediaTypeString, ok := xmlpath.MustCompile(fmt.Sprintf("%s[%d]/mediaType", e.xPathXmlAssets, pos)).String(assetNode)
		if !ok || mediaTypeString != "MP4" {
			continue
		}

		frameWidthString, ok := xmlpath.MustCompile(fmt.Sprintf("%s[%d]/frameWidth", e.xPathXmlAssets, pos)).String(assetNode)
		if !ok {
			continue
		}

		frameHeightString, ok := xmlpath.MustCompile(fmt.Sprintf("%s[%d]/frameHeight", e.xPathXmlAssets, pos)).String(assetNode)
		if !ok {
			continue
		}

		fileSizeString, ok := xmlpath.MustCompile(fmt.Sprintf("%s[%d]/fileSize", e.xPathXmlAssets, pos)).String(assetNode)
		if !ok {
			continue
		}

		progressiveDownloadUrlString, ok := xmlpath.MustCompile(fmt.Sprintf("%s[%d]/progressiveDownloadUrl", e.xPathXmlAssets, pos)).String(assetNode)
		if !ok {
			continue
		}

		originURL, err := url.Parse(progressiveDownloadUrlString)
		if err != nil {
			continue
		}

		frameWidthInt, err := strconv.Atoi(frameWidthString)
		if err != nil {
			continue
		}

		fileSizeInt, err := strconv.Atoi(fileSizeString)
		if err != nil {
			continue
		}

		frameHeightInt, err := strconv.Atoi(frameHeightString)
		if err != nil {
			continue
		}

		// we want just the best
		currentPixels := frameWidthInt * frameHeightInt

		if currentPixels > lastPixels {
			bestAsset = &videoAsset{
				Width:    frameWidthInt,
				Height:   frameHeightInt,
				FileSize: int64(fileSizeInt),
				VideoURL: originURL,
			}

			lastPixels = currentPixels
		}
	}

	if bestAsset == nil {
		return nil, fmt.Errorf("couldn't find any valid bestAsset")
	}

	return bestAsset, nil
}

func (e *xmlVideoExtractor) splitFullEpisodeTitle(fullTitle string) (int, string, error) {
	r, err := regexp.Compile("^(\\d+). +(.*)$")
	if err != nil {
		return -1, "", err
	}

	episodeNo := -1
	episodeTitle := ""

	if r.MatchString(fullTitle) && r.NumSubexp() == 2 {
		subs := r.FindAllStringSubmatch(fullTitle, -1)
		episodeNo, err = strconv.Atoi(subs[0][1])
		if err != nil {
			return -1, "", err
		}

		episodeTitle = subs[0][2]
	}

	if episodeNo < 1 {
		return -1, "", fmt.Errorf("couldn't find episode number imn xml document")
	}

	if episodeTitle == "" {
		return -1, "", fmt.Errorf("couldn't find episode title in xml document")
	}

	return episodeNo, episodeTitle, nil
}

func (e *xmlVideoExtractor) findIdOfXmlElement(node *xmlpath.Node) (string, error) {
	videoId := ""

	videoNodesPath, err := xmlpath.Compile(e.xPathVideoPageVideoTags)
	if err != nil {
		return "", err
	}

	videoNodesIter := videoNodesPath.Iter(node)

	for videoNodesIter.Next() {
		videoNode := videoNodesIter.Node()

		rawId := videoNode.String()

		r, err := regexp.Compile(e.regExpVideoId)
		if err != nil {
			return "", err
		}

		if r.MatchString(rawId) {
			videoId = strings.Replace(rawId, "html5-", "", 1)
			break
		}
	}

	if videoId == "" {
		return "", fmt.Errorf("couldn't find valid video id")
	}

	return videoId, nil
}

func (e *xmlVideoExtractor) findXmlDataUrl(node *xmlpath.Node, videoId string) (*url.URL, error) {
	xmlDataUrl := ""

	xpath := fmt.Sprintf(e.xPathVideoPageXmlDataTag, videoId)

	linkNodePath, err := xmlpath.Compile(xpath)
	if err != nil {
		return nil, err
	}

	linkNodeIter := linkNodePath.Iter(node)

	for linkNodeIter.Next() {
		value := linkNodeIter.Node().String()

		r, err := regexp.Compile(e.regExpXmlDataUrl)
		if err != nil {
			return nil, err
		}

		if r.MatchString(value) && r.NumSubexp() == 1 {
			subs := r.FindAllStringSubmatch(value, -1)
			xmlDataUrl = subs[0][1]
		}
	}

	if xmlDataUrl == "" {
		return nil, fmt.Errorf("couldn't find xml data url")
	}

	parsedXmlDataUrl, err := url.Parse(xmlDataUrl)
	if err != nil {
		return nil, err
	}

	return parsedXmlDataUrl, nil
}
