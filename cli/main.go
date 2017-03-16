package main

import (
	"flag"
	"fmt"
	"github.com/buger/goterm"
	"net/url"
	"os"
	"path/filepath"
	"rkl.io/kika-downloader/cli/commands"
	"rkl.io/kika-downloader/cli/config"
	cliContract "rkl.io/kika-downloader/cli/contract"
	"rkl.io/kika-downloader/cli/dto"
	"time"
	"sort"
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
		fmt.Fprintf(os.Stderr, "  fetch-all : download all videos to given directory\n")
		fmt.Fprintf(os.Stderr, "    -url=<entry url>\n")
		fmt.Fprintf(os.Stderr, "    -output-dir=<download dir>\n")
		fmt.Fprintf(os.Stderr, "    -max-simultaneous-downloads=[number, default 3]\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  print-csv : print csv like output of all videos\n")
		fmt.Fprintf(os.Stderr, "    -url=<entry url>\n")
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
	maxSimultaneousDownloads := flagSet.Int("max-simultaneous-downloads", 3, "maximum simultaneous downloads")

	flagSet.Parse(args[2:])

	appContext, err := config.InitCliContext(*sockProxyUrl)
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

	command := commands.NewFetchAll(entryUrl, *fetchAllOutputDir, *maxSimultaneousDownloads)

	service, err := appContext.SafeGet("command_handler.fetch_all")
	if err != nil {
		return err
	}

	handler := service.(cliContract.CommandHandlerInterface)

	go func() {
		goterm.Clear()
		goterm.MoveCursor(1, 1)

		pm := map[string] dto.EpisodeDownloadProgress{}

		initTime := time.Now()

		flushProgressMap := func() {
			var keys []string
			for k := range pm {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			if time.Since(initTime) >= time.Second {
				goterm.Clear()
				goterm.MoveCursor(1, 1)

				for _, k := range keys {
					goterm.Println(
						fmt.Sprintf("%s%% of \"%s - %s\" done",
							pm[k].GetPercentage(),
							pm[k].GetSeriesTitle(),
							pm[k].GetEpisodeTitle(),
						),
					)
				}

				goterm.Flush()
				initTime = time.Now()
			}
		}

		for progressDtoInterface := range handler.GetDtoOutputChannel() {
			switch progressDtoInterface.(type) {
			case dto.EpisodeDownloadProgress:
				progressDto := progressDtoInterface.(dto.EpisodeDownloadProgress)

				key := progressDto.GetSeriesTitle() + progressDto.GetEpisodeTitle()

				pm[key] = progressDto

				flushProgressMap()

				break
			}
		}

		flushProgressMap()
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

	appContext, err := config.InitCliContext(*sockProxyUrl)
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

	handler := service.(cliContract.CommandHandlerInterface)

	if _, err := handler.Handle(command); err != nil {
		return err
	}

	return nil
}
