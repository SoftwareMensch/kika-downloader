package config

import (
	"github.com/sarulabs/di"
	coreServices "rkl.io/kika-downloader/core/services"
)

// InitCoreBuilder init core
func InitCoreBuilder(socksProxyUrl string) (*di.Builder, error) {
	coreBuilder, err := di.NewBuilder(di.App)
	if err != nil {
		return nil, err
	}

	if err := coreBuilder.Set("socks_proxy_url", socksProxyUrl); err != nil {
		return nil, err
	}

	if err := coreBuilder.Set("regexp_video_id", RegExpVideoId); err != nil {
		return nil, err
	}

	if err := coreBuilder.Set("regexp_xml_data_url", RegExpXmlDataUrl); err != nil {
		return nil, err
	}

	if err := coreBuilder.Set("xpath_episodes_overview_page_page_items", XPathEpisodesOverviewPagePageItems); err != nil {
		return nil, err
	}

	if err := coreBuilder.Set("xpath_episodes_items", XPathEpisodesItems); err != nil {
		return nil, err
	}

	if err := coreBuilder.Set("xpath_video_page_video_tags", XPathVideoPageVideoTags); err != nil {
		return nil, err
	}

	if err := coreBuilder.Set("xpath_video_page_xml_data_tag", XPathVideoPageXmlDataTag); err != nil {
		return nil, err
	}

	if err := coreBuilder.Set("xpath_xml_series_title", XPathXmlSeriesTitle); err != nil {
		return nil, err
	}

	if err := coreBuilder.Set("xpath_xml_episode_title", XPathXmlEpisodeTitle); err != nil {
		return nil, err
	}

	if err := coreBuilder.Set("xpath_xml_episode_language", XPathXmlEpisodeLanguageCode); err != nil {
		return nil, err
	}

	if err := coreBuilder.Set("xpath_xml_episode_description", XPathXmlEpisodeDescription); err != nil {
		return nil, err
	}

	if err := coreBuilder.Set("xpath_xml_assets", XPathXmlAssets); err != nil {
		return nil, err
	}

	// http client service
	if err := coreServices.AssignHttpClient(coreBuilder); err != nil {
		return nil, err
	}

	// episodes overview crawler service
	if err := coreServices.AssignEpisodesOverviewUrlIterator(coreBuilder); err != nil {
		return nil, err
	}

	// episodes items crawler service
	if err := coreServices.AssignEpisodesOverviewItemsUrlIterator(coreBuilder); err != nil {
		return nil, err
	}

	// video meta data extractor service
	if err := coreServices.AssignExtractor(coreBuilder); err != nil {
		return nil, err
	}

	// video downloader
	if err := coreServices.AssignVideoDownloader(coreBuilder); err != nil {
		return nil, err
	}

	return coreBuilder, nil
}
