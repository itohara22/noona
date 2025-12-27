package main

import (
	"log"
	"os"
)

func ReadFile(fileLocation *string) []byte {

	if *fileLocation == "" {
		log.Fatal("Provide torrent file location")
	}

	fileData, err := os.ReadFile(*fileLocation)
	if err != nil {
		log.Fatal("cannot read file")
		log.Fatal(err.Error())
	}
	return fileData
}
