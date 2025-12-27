package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fileLocation := flag.String("f", "", "torrent file location")
	flag.Parse()

	if *fileLocation == "" {
		log.Fatal("Provide torrent file location")
	}
	fmt.Println(*fileLocation)

	fileData, err := os.ReadFile(*fileLocation)
	if err != nil {
		log.Fatal("cannot read file")
		log.Fatal(err.Error())
	}
	dic := fileData

	bencodeDecoder := newBencodeDecoder()
	res := bencodeDecoder.Decode(dic)
	resMap, ok := res.(map[string]any)
	if !ok {
		log.Fatal("tch tch tch")
	}

	trackerList := resMap["announce-list"].([]any)
	for _, t := range trackerList {
		tStrArr, ok := t.([]any)
		if !ok {
			log.Fatal("something wrong with annouce-list")
		}
		tBytes, ok := tStrArr[0].([]byte)
		if !ok {
			log.Fatal("something wrong with annouce-list[inside stuff]")
		}
		fmt.Println(string(tBytes))
	}

	en := NewBencodeEncoder()
	enCodedData := en.Encode(resMap["info"])

	hash := sha1.Sum(enCodedData)
	fmt.Print(hash)

	info, ok := resMap["info"].(map[string]any)
	if !ok {
		log.Fatal("tch tch info")
	}
	pieces := info["pieces"]
	fmt.Println(pieces)

}
