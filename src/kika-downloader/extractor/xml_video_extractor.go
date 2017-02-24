package extractor

import (
	"kika-downloader/http"
	"kika-downloader/contract"
	"net/url"
	"kika-downloader/utils"
	"gopkg.in/xmlpath.v2"
	"fmt"
	"strings"
	"regexp"
	"log"
)

type xmlVideoExtractor struct{
	httpClient http.ClientInterface

	xPathVideoPageVideoTags string
	xPathVideoPageXmlDataTag string

	regExpVideoId string
	regExpXmlDataUrl string
}

// NewXmlVideoExtractor return xml meta data extractor
func NewXmlVideoExtractor(
	httpClient http.ClientInterface,
	xPathVideoPageVideoTags string,
	xPathVideoPageXmlDataTag string,
	regExpVideoId string,
	regExpXmlDataUrl string,
) contract.VideoExtractorInterface {
	return &xmlVideoExtractor{
		httpClient: httpClient,
		xPathVideoPageVideoTags: xPathVideoPageVideoTags,
		xPathVideoPageXmlDataTag: xPathVideoPageXmlDataTag,
		regExpVideoId: regExpVideoId,
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

	videoId, err := e.findVideoId(rootDoc)
	if err != nil {
		return nil, err
	}

	xmlUrl, err := e.findXmlDataUrl(rootDoc, videoId)
	if err != nil {
		return nil, err
	}

	// TODO, go ahead
	xmlUrl = xmlUrl

	return nil, nil
}

func (e *xmlVideoExtractor) findVideoId(node *xmlpath.Node) (string, error) {
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
