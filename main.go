package main

import (
	"log"

	"github.com/nino-k/gitlist/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
