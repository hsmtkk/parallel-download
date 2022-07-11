package download

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type downloader struct {
	id              int
	urlsChan        <-chan string
	oupputDirectory string
}

func New(id int, urlsChan <-chan string, outputDirectory string) *downloader {
	return &downloader{id, urlsChan, outputDirectory}
}

func (d *downloader) Run() {
	log.Printf("downloader %d begin", d.id)
	for url := range d.urlsChan {
		if err := d.downloadAndSave(url); err != nil {
			log.Print(err)
		}
	}
	log.Printf("downloader %d end", d.id)
}

func (d *downloader) downloadAndSave(urlString string) error {
	log.Printf("download begin %s", urlString)
	resp, err := http.Get(urlString)
	if err != nil {
		return fmt.Errorf("failed to download; %s; %w", urlString, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("got non 2xx status code; %d; %s", resp.StatusCode, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response; %w", err)
	}
	u, err := url.Parse(urlString)
	if err != nil {
		return fmt.Errorf("failed to parse url; %s; %w", urlString, err)
	}
	path := filepath.Join(d.oupputDirectory, u.Path)
	if err := os.WriteFile(path, body, 0644); err != nil {
		return fmt.Errorf("failed to write file; %s; %w", path, err)
	}
	log.Printf("download end %s", urlString)
	return nil
}
