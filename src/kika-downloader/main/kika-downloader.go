package main

import (
	"fmt"
	"log"
	"os"

	"kika-downloader/config"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal(fmt.Sprintf("%s [episodes overview url]", os.Args[0]))
		os.Exit(-1)
	}

	torSocksURL := os.Args[1]
	episodesOverviewURL := os.Args[2]

	appContext, err := config.SetupApp(torSocksURL, episodesOverviewURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(appContext)
}
