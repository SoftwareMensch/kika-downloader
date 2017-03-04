package main

import (
	"flag"
	"fmt"
	"kika-downloader/commands"
	"kika-downloader/config"
	"kika-downloader/contract"
	"log"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s <command> [<args>]\n", os.Args[0])
		os.Exit(1)
	}

	sockProxyUrl := flag.String("socks-proxy-url", "", "url of socks proxy")

	flag.Parse()

	appContext, err := config.InitApp(*sockProxyUrl)
	if err != nil {
		log.Fatal(err)
	}

	// fetch-all command
	fetchAllCommandFlagSet := flag.NewFlagSet("fetch-all", flag.ExitOnError)
	fetchAllUrl := fetchAllCommandFlagSet.String("url", "", "entry url")
	fetchAllOutputDir := fetchAllCommandFlagSet.String("output-dir", "", "download directory")
	fetchAllCommandFlagSet.Parse(os.Args[2:])

	// print-csv command
	printCsvCommandFlagSet := flag.NewFlagSet("print-csv", flag.ExitOnError)
	printCsvEntryUrl := printCsvCommandFlagSet.String("url", "", "entry url")
	printCsvCommandFlagSet.Parse(os.Args[2:])

	switch os.Args[1] {
	case "fetch-all":
		if *fetchAllUrl == "" {
			log.Fatal("please provide entry url")
		}

		entryUrl, err := url.Parse(*fetchAllUrl)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := os.Stat(*fetchAllOutputDir); err != nil {
			log.Fatal(err)
		}

		command := commands.NewFetchAll(entryUrl, *fetchAllOutputDir)

		service, err := appContext.SafeGet("command_handler.fetch_all")
		if err != nil {
			log.Fatal(err)
		}

		handler := service.(contract.CommandHandlerInterface)

		if _, err := handler.Handle(command); err != nil {
			log.Fatal(err)
		}

	case "print-csv":
		if *printCsvEntryUrl == "" {
			log.Fatal("please provice entry url")
		}

		entryUrl, err := url.Parse(*printCsvEntryUrl)
		if err != nil {
			log.Fatal(err)
		}

		command := commands.NewPrintCsvCommand(entryUrl)

		service, err := appContext.SafeGet("command_handler.print_csv")
		if err != nil {
			log.Fatal(err)
		}

		handler := service.(contract.CommandHandlerInterface)

		if _, err := handler.Handle(command); err != nil {
			log.Fatal(err)
		}

	default:
		fmt.Printf("%s is not a valid command\n", os.Args[1])
	}
}
