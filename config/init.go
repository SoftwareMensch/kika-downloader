package config

import (
	"github.com/sarulabs/di"
	cliServices "rkl.io/kika-downloader/cli/services"
	coreServices "rkl.io/kika-downloader/core/services"
	coreConfig "rkl.io/kika-downloader/core/config"
)

// InitApp basic app set up
func InitApp(socksProxyUrl string) (di.Context, error) {
	appBuilder, err := di.NewBuilder("app")
	if err != nil {
		return nil, err
	}

	if err := appBuilder.Set("socks_proxy_url", socksProxyUrl); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("regexp_video_id", coreConfig.RegExpVideoId); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("regexp_xml_data_url", coreConfig.RegExpXmlDataUrl); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_episodes_overview_page_page_items", coreConfig.XPathEpisodesOverviewPagePageItems); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_episodes_items", coreConfig.XPathEpisodesItems); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_video_page_video_tags", coreConfig.XPathVideoPageVideoTags); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_video_page_xml_data_tag", coreConfig.XPathVideoPageXmlDataTag); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_xml_series_title", coreConfig.XPathXmlSeriesTitle); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_xml_episode_title", coreConfig.XPathXmlEpisodeTitle); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_xml_episode_language", coreConfig.XPathXmlEpisodeLanguageCode); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_xml_episode_description", coreConfig.XPathXmlEpisodeDescription); err != nil {
		return nil, err
	}

	if err := appBuilder.Set("xpath_xml_assets", coreConfig.XPathXmlAssets); err != nil {
		return nil, err
	}

	// http client service
	if err := coreServices.AssignHttpClient(appBuilder); err != nil {
		return nil, err
	}

	// episodes overview crawler service
	if err := coreServices.AssignEpisodesOverviewUrlIterator(appBuilder); err != nil {
		return nil, err
	}

	// episodes items crawler service
	if err := coreServices.AssignEpisodesOverviewItemsUrlIterator(appBuilder); err != nil {
		return nil, err
	}

	// video meta data extractor service
	if err := coreServices.AssignExtractor(appBuilder); err != nil {
		return nil, err
	}

	// video downloader
	if err := coreServices.AssignVideoDownloader(appBuilder); err != nil {
		return nil, err
	}

	// fetch all command handler
	if err := cliServices.AssignFetchAllCommandHandler(appBuilder); err != nil {
		return nil, err
	}

	// print csv command handler
	if err := cliServices.AssignPrintCsvCommandHandler(appBuilder); err != nil {
		return nil, err
	}

	return appBuilder.Build(), nil
}
