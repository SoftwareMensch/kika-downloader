package view

import (
	"crypto/md5"
	"fmt"
	tm "github.com/buger/goterm"
	"rkl.io/kika-downloader/cli/dto"
	"sort"
)

type terminal struct {
	currentIndex     int
	progressIndexMap map[string]int
	dtoProgressMap   map[int]dto.EpisodeDownloadProgress
}

// NewTerminal return new terminal view
func NewTerminal() *terminal {
	return &terminal{
		currentIndex: -1,
	}
}

// GetCurrentIndex get current index
func (t *terminal) GetCurrentIndex() int {
	return t.currentIndex
}

// GetLatestProgressByIndex get latest progress by index
func (t *terminal) GetLatestProgressByIndex(index int) *dto.EpisodeDownloadProgress {
	if _, ok := t.dtoProgressMap[index]; !ok {
		return nil
	}

	p := t.dtoProgressMap[index]

	return &p
}

// UpdateEpisodeDownloadProgress update proress of episode download
func (t *terminal) UpdateEpisodeDownloadProgress(p dto.EpisodeDownloadProgress) {
	if t.progressIndexMap == nil {
		t.progressIndexMap = make(map[string]int)
	}

	if t.dtoProgressMap == nil {
		t.dtoProgressMap = make(map[int]dto.EpisodeDownloadProgress)
	}

	indexId := fmt.Sprintf("%x", md5.Sum([]byte(p.GetSeriesTitle()+p.GetEpisodeTitle())))

	if _, ok := t.progressIndexMap[indexId]; !ok {
		t.currentIndex++
		t.progressIndexMap[indexId] = t.currentIndex
	}

	t.dtoProgressMap[t.progressIndexMap[indexId]] = p
}

// Render render output
func (t *terminal) Render() {
	if t.dtoProgressMap == nil {
		return
	}

	var indexes []int

	for i := range t.dtoProgressMap {
		indexes = append(indexes, i)
	}

	sort.Ints(indexes)

	tm.Clear()
	tm.MoveCursor(1, 1)

	for _, i := range indexes {
		p := t.dtoProgressMap[i]

		tm.Println(
			fmt.Sprintf(
				"%-9s of \"%s - %s\" done",
				"["+p.GetPercentage()+"%]",
				p.GetSeriesTitle(),
				p.GetEpisodeTitle(),
			),
		)
	}

	tm.Flush()
}
