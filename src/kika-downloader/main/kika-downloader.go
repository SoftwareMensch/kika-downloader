package main

import (
	"fmt"
	"log"

	"flag"
	"kika-downloader/commands"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s <command> [<args>]\n", os.Args[0])
		os.Exit(1)
	}

	//	socksProxyURL := flag.String("socks-proxy-url", "", "url of socks proxy")

	fetchAllCommandFlagSet := flag.NewFlagSet("fetch-all", flag.ExitOnError)

	switch os.Args[1] {
	case "fetch-all":
		fetchAllCommand, err := makeFetchAllCommand(os.Args[2:], fetchAllCommandFlagSet)
		if err != nil {
			log.Fatal(err)
		}

		_, err = commands.NewFetchAllHandler().Handle(fetchAllCommand)
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Printf("%s is not a valid command\n", os.Args[1])
	}

	flag.Parse()
}

func makeFetchAllCommand(args []string, flagSet *flag.FlagSet) (*commands.FetchAll, error) {
	if err := flagSet.Parse(args); err != nil {
		log.Fatal(err)
	}

	urlFlag := flagSet.String("url", "", "entry url")
	if *urlFlag == "" {
		return nil, fmt.Errorf("please supply the entry url with -url")
	}

	parsedURL, err := url.Parse(*urlFlag)
	if err != nil {
		return nil, err
	}

	return commands.NewFetchAll(parsedURL), nil
}
