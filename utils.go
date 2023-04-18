package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configs struct {
}

func getConfigs() Configs {
	// Read configs
	config := Configs{}
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func(configFile *os.File) {
		errConfigClose := configFile.Close()
		if errConfigClose != nil {
			fmt.Println("Error while closing config file: ", err)
		}
	}(configFile)
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&config); err != nil {
		log.Fatal(err)
	}
	return config
}
