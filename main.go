package main

import (
	"flag"
	"fmt"
	"os"

	configuration "github.com/Shakarang/Orgmanager/Configuration"
)

var githubToken string

func main() {

	flag.StringVar(&githubToken, "t", "", "Set token using -t option.")
	filePath := flag.String("f", "", "Add a file using -f option.")

	flag.Parse()

	// Check if token and config file exist
	if len(githubToken) == 0 || len(*filePath) == 0 {
		fmt.Println("Set github token and configuration file path.")
		os.Exit(2)
	}

	config, err := configuration.FromFileAt(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(*config)

	// configuration.FromFileAt(path)
	// configuration.NewConfiguration()
}
