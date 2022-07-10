package main

import (
	"log"

	"github.com/hsmtkk/parallel-download/command"
)

func main() {
	cmd := command.Command
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}