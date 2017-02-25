package extractor

import (
	"fmt"
	"gopkg.in/xmlpath.v2"
	"kika-downloader/contract"
	"kika-downloader/http"
	"kika-downloader/utils"
	"log"
	"net/url"
	"regexp"
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

	regExpVideoId    string
	regExpXmlDataUrl string
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

	return e.makeVideoFromXmlRoot(xmlRoot)
}

func (e *xmlVideoExtractor) makeVideoFromXmlRoot(xml *xmlpath.Node) (contract.VideoInterface, error) {
	seriesTitlePath, err := xmlpath.Compile(e.xPathXmlSeriesTitle)
	if err != nil {
		return nil, err
	}

	fullEposideTitlePath, err := xmlpath.Compile(e.xPathXmlEpisodeTitle)
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
	log.Printf("found series title: %s\n", seriesTitle)

	fullEpisodesTitle, ok := fullEposideTitlePath.String(xml)
	if !ok {
		return nil, fmt.Errorf("no episodes title found in xml document")
	}
	log.Printf("found full episode title: %s\n", fullEpisodesTitle)

	episodeLanguageCode, ok := episodeLanguageCodePath.String(xml)
	if !ok {
		return nil, fmt.Errorf("no language code found in xml document")
	}
	log.Printf("found language code: %s\n", episodeLanguageCode)

	episodeDescription, ok := episodeDescriptionPath.String(xml)
	if !ok {
		return nil, fmt.Errorf("no episode description found in xml document")
	}
	log.Println("--- episode description ---")
	log.Println(episodeDescription)
	log.Println("--- episode description ---")


	return nil, nil
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
			log.Printf("found video id: %s\n", videoId)
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
			log.Printf("found xml data url: %s\n", xmlDataUrl)
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
