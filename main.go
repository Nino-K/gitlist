package main

import (
	"log"

	"github.com/Nino-K/gitlist/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
