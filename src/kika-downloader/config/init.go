package config

import (
	"github.com/sarulabs/di"
	"kika-downloader/services"
	"net/url"
)

// InitApp basic app set up
func InitApp(socksProxyUrl, episodesOverviewURL string) (di.Context, error) {
	appBuilder, err := di.NewBuilder("app")
	if err != nil {
		return nil, err
	}

	if err := appBuilder.Set("socks_proxy_url", socksProxyUrl); err != nil {
		return nil, err
	}

	parsedEpisodesOverviewURL, err := url.Parse(episodesOverviewURL)
	if err != nil {
		return nil, err
	}

	if err := appBuilder.Set("episodes_overview_url", parsedEpisodesOverviewURL); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("regexp_video_id", RegExpVideoId); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("regexp_xml_data_url", RegExpXmlDataUrl); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_episodes_overview_page_page_items", XPathEpisodesOverviewPagePageItems); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_episodes_items", XPathEpisodesItems); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_video_page_video_tags", XPathVideoPageVideoTags); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_video_page_xml_data_tag", XPathVideoPageXmlDataTag); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_xml_series_title", XPathXmlSeriesTitle); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_xml_episode_title", XPathXmlEpisodeTitle); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_xml_episode_language", XPathXmlEpisodeLanguageCode); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_xml_episode_description", XPathXmlEpisodeDescription); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_xml_assets", XPathXmlAssets); err != nil {
		return nil, err
	}

	// http client service
	if err := services.AssignHttpClient(appBuilder); err != nil {
		return nil, err
	}

	// episodes overview crawler service
	if err := services.AssignEpisodesOverviewUrlIterator(appBuilder); err != nil {
		return nil, err
	}

	// episodes items crawler service
	if err := services.AssignEpisodesOverviewItemsUrlIterator(appBuilder); err != nil {
		return nil, err
	}

	// video meta data extractor service
	if err := services.AssignExtractor(appBuilder); err != nil {
		return nil, err
	}

	// video downloader
	if err := services.AssignVideoDownloader(appBuilder); err != nil {
		return nil, err
	}

	return appBuilder.Build(), nil
}
