package command

import (
	"log"
	"sync"

	"github.com/hsmtkk/parallel-download/download"
	"github.com/hsmtkk/parallel-download/file"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:  "parallel-download url-list-file output-directory",
	Args: cobra.ExactArgs(2),
	Run:  run,
}

var parallel int

func init() {
	Command.Flags().IntVar(&parallel, "parallel", 4, "number of go routines")
}

func run(cmd *cobra.Command, args []string) {
	urlListFile := args[0]
	outputDirectory := args[1]

	urls, err := file.ReadLines(urlListFile)
	if err != nil {
		log.Fatal(err)
	}

	urlsChan := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < parallel; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			downloader := download.New(id, urlsChan, outputDirectory)
			downloader.Run()
		}(i)
	}

	for _, url := range urls {
		urlsChan <- url
	}
	close(urlsChan)

	wg.Wait()
}
