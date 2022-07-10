package file

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file; %s; %w", filePath, err)
	}
	defer f.Close()

	urls := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error while reading file; %w", err)
	}

	return urls, nil
}
