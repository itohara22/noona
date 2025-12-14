package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileData, err := os.ReadFile("a.torrent")
	if err != nil {
		log.Fatal("cannot read file")
		log.Fatal(err.Error())
	}
	dic := string(fileData)

	bencodeDecoder := newBencodeDecoder()
	res := bencodeDecoder.Decode(dic)
	resMap, ok := res.(map[string]any)
	if !ok {
		log.Fatal("tch tch tch")
	}
	fmt.Println(resMap["info"])
}

// dic := "d3:cow3:moo4:spam4:eggs3:numi42e4:listl4:test3:one3:twoe3:fooi99ee"
