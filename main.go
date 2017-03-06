package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MarquisIO/Organisator/configuration"
	"github.com/MarquisIO/Organisator/github"
)

func main() {

	githubToken := flag.String("t", "", "Set token using -t option.")
	filePath := flag.String("f", "", "Add a file using -f option.")

	flag.Parse()

	// Check if token and config file exist
	if len(*githubToken) == 0 || len(*filePath) == 0 {
		fmt.Println("Set github token and configuration file path.")
		os.Exit(2)
	}

	config, err := configuration.FromFileAt(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := github.Application{
		Token:        *githubToken,
		Organisation: config.Organisation,
		Config:       config,
	}

	if err := app.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
