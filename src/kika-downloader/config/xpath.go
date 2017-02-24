package config

const (
	// XPathEpisodesOverviewPagePageItems to retrieve pages where videos are schon
	XPathEpisodesOverviewPagePageItems = "//a[@class='pageItem']/@href"

	// XPathEpisodesItems to retrieve video urls from a  page
	XPathEpisodesItems = "//div[@class='modCon']/div[@class='mod modD modMini']/div[@class='boxCon']/div[contains(@class, 'box')]/div[contains(@class, 'teaser')]/a[@class='linkAll']/@href"

	// XPathVideoPageVideoTags to find all video tags on final video page
	XPathVideoPageVideoTags = "//video/@id"

	// XPathVideoPageXmlDataTag to find the concrete element where xml url can be extracted, %s is placeholder for video id
	XPathVideoPageXmlDataTag = "//a[@id='%s']/@onclick"
)
