package tests

import (
	"rkl.io/kika-downloader/cli/dto"
	"rkl.io/kika-downloader/cli/view"
	"testing"
)

func TestUpdateEpisodeDownloadProgress(t *testing.T) {
	terminalView := view.NewTerminal()

	if terminalView.GetCurrentIndex() != -1 {
		t.Errorf("initial index should be -1")
	}

	videoProgress01 := []dto.EpisodeDownloadProgress{
		dto.NewEpisodeDownloadProgress("10.00", "Series Title", "Episode Title"),
		dto.NewEpisodeDownloadProgress("20.00", "Series Title", "Episode Title"),
		dto.NewEpisodeDownloadProgress("30.00", "Series Title", "Episode Title"),
		dto.NewEpisodeDownloadProgress("40.00", "Series Title", "Episode Title"),
		dto.NewEpisodeDownloadProgress("50.00", "Series Title", "Episode Title"),
		dto.NewEpisodeDownloadProgress("60.00", "Series Title", "Episode Title"),
		dto.NewEpisodeDownloadProgress("70.00", "Series Title", "Episode Title"),
	}

	videoProgress02 := []dto.EpisodeDownloadProgress{
		dto.NewEpisodeDownloadProgress("10.00", "Series Title", "Other Title"),
		dto.NewEpisodeDownloadProgress("20.00", "Series Title", "Other Title"),
		dto.NewEpisodeDownloadProgress("30.00", "Series Title", "Other Title"),
	}

	for _, p := range videoProgress01 {
		terminalView.UpdateEpisodeDownloadProgress(p)
	}

	// test index
	if terminalView.GetCurrentIndex() != 0 {
		t.Errorf("current index should be 0")
	}

	// test latest progress
	lp := terminalView.GetLatestProgressByIndex(0)
	if lp == nil {
		t.Errorf("latest progress should not be nil")
	}

	if lp.GetPercentage() != "70.00" {
		t.Errorf("expected 70.00 as latest progress, got %s", lp.GetPercentage())
	}

	for _, p := range videoProgress02 {
		terminalView.UpdateEpisodeDownloadProgress(p)
	}

	// test index
	if terminalView.GetCurrentIndex() != 1 {
		t.Errorf("current index should be 0")
	}

	// test latest progress
	lp = terminalView.GetLatestProgressByIndex(0)
	if lp == nil {
		t.Errorf("latest progress should not be nil")
	}

	lp = terminalView.GetLatestProgressByIndex(1)
	if lp == nil {
		t.Errorf("latest progress should not be nil")
	}

	if lp.GetPercentage() != "30.00" {
		t.Errorf("expected 30.00 as latest progress, got %s", lp.GetPercentage())
	}

	terminalView.Render()
}
