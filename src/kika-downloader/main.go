package main

import (
	"flag"
	"fmt"
	"kika-downloader/commands"
	"kika-downloader/config"
	"kika-downloader/contract"
	"kika-downloader/dto"
	"net/url"
	"os"
	"path/filepath"
)

func main() {
	var commandError error

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  -socks-proxy-url=<socks5://127.0.0.1:9050>               optional socks proxy (i.e. TOR)\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  fetch-all -url=<entry url> -output-dir=<download dir>    download all videos to given directory\n")
		fmt.Fprintf(os.Stderr, "  print-csv -url=<entry url>                               print csv like output of all videos\n")
		fmt.Fprintf(os.Stderr, "\n")
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	fetchAllFlagSet := flag.NewFlagSet("fetch-all", flag.ExitOnError)
	printCsvFlagSet := flag.NewFlagSet("print-csv", flag.ExitOnError)

	flag.Parse()

	switch os.Args[1] {
	case "fetch-all":
		commandError = runFetchAllCommand(fetchAllFlagSet, os.Args)

	case "print-csv":
		commandError = runPrintCsvCommand(printCsvFlagSet, os.Args)

	default:
		commandError = fmt.Errorf(fmt.Sprintf("%s is not a valid command\n", os.Args[1]))
	}

	if commandError != nil {
		fmt.Fprintf(os.Stderr, "Unknown command \"%s\":\n\n", os.Args[1])
		flag.Usage()
		os.Exit(2)
	}
}

func runFetchAllCommand(flagSet *flag.FlagSet, args []string) error {
	sockProxyUrl := flagSet.String("socks-proxy-url", "", "url of socks proxy")
	fetchAllUrl := flagSet.String("url", "", "entry url")
	fetchAllOutputDir := flagSet.String("output-dir", "", "download directory")

	flagSet.Parse(args[2:])

	appContext, err := config.InitApp(*sockProxyUrl)
	if err != nil {
		return err
	}

	if *fetchAllUrl == "" {
		return fmt.Errorf("please provide entry url")
	}

	entryUrl, err := url.Parse(*fetchAllUrl)
	if err != nil {
		return err
	}

	if _, err := os.Stat(*fetchAllOutputDir); err != nil {
		return err
	}

	command := commands.NewFetchAll(entryUrl, *fetchAllOutputDir)

	service, err := appContext.SafeGet("command_handler.fetch_all")
	if err != nil {
		return err
	}

	handler := service.(contract.CommandHandlerInterface)

	go func() {
		for progressDtoInterface := range handler.GetDtoOutputChannel() {
			switch progressDtoInterface.(type) {
			case dto.EpisodeDownloadProgress:
				progressDto := progressDtoInterface.(dto.EpisodeDownloadProgress)
				fmt.Printf("\r%s%% of \"%s - %s\" done",
					progressDto.GetPercentage(),
					progressDto.GetSeriesTitle(),
					progressDto.GetEpisodeTitle(),
				)

				break
			default:
				fmt.Print("\n")
			}
		}
	}()

	if _, err := handler.Handle(command); err != nil {
		return err
	}

	return nil
}

func runPrintCsvCommand(flagSet *flag.FlagSet, args []string) error {
	sockProxyUrl := flagSet.String("socks-proxy-url", "", "url of socks proxy")
	printCsvEntryUrl := flagSet.String("url", "", "entry url")

	flagSet.Parse(args[2:])

	appContext, err := config.InitApp(*sockProxyUrl)
	if err != nil {
		return err
	}

	if *printCsvEntryUrl == "" {
		return fmt.Errorf("please provice entry url")
	}

	entryUrl, err := url.Parse(*printCsvEntryUrl)
	if err != nil {
		return err
	}

	command := commands.NewPrintCsvCommand(entryUrl)

	service, err := appContext.SafeGet("command_handler.print_csv")
	if err != nil {
		return err
	}

	handler := service.(contract.CommandHandlerInterface)

	if _, err := handler.Handle(command); err != nil {
		return err
	}

	return nil
}
