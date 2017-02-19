package main

import (
	"crawler"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal(fmt.Sprintf("%s [episodes overview url]", os.Args[0]))
	}

	episodesOverviewURL := os.Args[1]

	episodesOverviewCrawler, err := crawler.NewEpisodesOverview(episodesOverviewURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(episodesOverviewCrawler)
}
