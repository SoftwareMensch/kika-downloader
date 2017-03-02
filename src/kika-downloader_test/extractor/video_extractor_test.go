package extractor

import (
	"fmt"
	"kika-downloader/config"
	"kika-downloader/contract"
	"kika-downloader/utils"
	testConfig "kika-downloader_test/config"
	"testing"
)

func TestVideoExtraction(t *testing.T) {
	appContext, err := config.InitApp(testConfig.TorSocksProxyURL)
	if err != nil {
		t.Error(err)
	}

	service, err := appContext.SafeGet("video_extractor")
	if err != nil {
		t.Error(err)
	}

	metaDataExtractor := service.(contract.VideoExtractorInterface)

	video, err := metaDataExtractor.ExtractVideoFromURL(testConfig.ExtractorTestURL)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("[T] Extracted video %s\n", video.GetId())

	if !utils.IsUUID4(video.GetId()) {
		t.Errorf("video if \"%s\" is no valid uuid version 4", video.GetId())
	}

	if video.GetSeriesTitle() != "Super Wings" {
		t.Errorf("\"%s\" does not match series title \"Super Wings\"", video.GetSeriesTitle())
	}

	if video.GetEpisodeTitle() != "Schwingt die Pinsel!" {
		t.Errorf("\"%s\" does not match episode title \"Schwingt die Pinsel!\"", video.GetEpisodeTitle())
	}

	if video.GetLanguageCode() != "de" {
		t.Errorf("\"%s\" does not match language code \"de\"", video.GetLanguageCode())
	}

	if video.GetEpisodeNumber() != 23 {
		t.Errorf("\"%d\" does not match episode numer \"23\"", video.GetEpisodeNumber())
	}

	if video.GetFileSize() != 339953956 {
		t.Errorf("\"%d\" does not match file size \"339953956\"", video.GetFileSize())
	}

	if video.GetVideoOriginURL().String() != "http://pmdonline.kika.de/mp4dyn/7/FCMS-7bdd1753-dd4d-404e-a652-d94ec2183363-5a2c8da1cdb7_7b.mp4" {
		t.Errorf("expected url does not match")
	}
}
