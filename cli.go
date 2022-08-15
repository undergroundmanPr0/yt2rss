package main

import (
	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"encoding/json"
)

type config struct {
	Api_key string
	Location string
}


func main() {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	getChannelID := getCmd.String("c", "", "Get Youtube channel ID")
	getFilename := getCmd.String("f", "", "Get filename to store the rss feed xml")
	
	if len(os.Args) < 2 {

		fmt.Println("expected \"get\" command")
		os.Exit(1)
	}

	if os.Args[1] != "get" {
		fmt.Println("Add appropriate commands and flags.")
		os.Exit(1)
	}


	switch os.Args[1] {
		case "get":
			handleGet(getCmd, getChannelID, getFilename)
		default:
			fmt.Println("Include the get command to run the api")
	}

}

func handleGet(getCmd *flag.FlagSet, channelID *string, filename *string) {
	getCmd.Parse(os.Args[2:])

	if *channelID == "" || *filename == "" {
		fmt.Println("All required fields are missing")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	var api_config []config
	fileBytes, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println("config.json not found, include one")
		os.Exit(1)
	}

	err = json.Unmarshal(fileBytes, &api_config)

	if err != nil {
		panic(err);
		os.Exit(1)
	}

	apiID := api_config[0].Api_key
	location := api_config[0].Location


	fileNameFinal := ""
	if location != "" {
		fileNameFinal = location + "/" + *filename
	} else {
		fileNameFinal = *filename
	}
	createRSSFeed(*channelID, apiID, fileNameFinal)
	return
}
