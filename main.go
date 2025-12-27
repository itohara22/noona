package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"log"
)

func main() {
	fileLocation := flag.String("f", "", "torrent file location")
	flag.Parse()

	torrentFileBytesData := ReadFile(fileLocation)

	bencodeDecoder := newBencodeDecoder()
	res := bencodeDecoder.Decode(torrentFileBytesData)

	resMap, ok := res.(map[string]any)
	if !ok {
		log.Fatal("tch tch tch")
	}

	en := NewBencodeEncoder()
	enCodedData := en.Encode(resMap["info"])

	hash := sha1.Sum(enCodedData)
	fmt.Print(hash)

}

// info, ok := resMap["info"].(map[string]any)
// if !ok {
// 	log.Fatal("tch tch info")
// }
// pieces := info["pieces"]
// fmt.Println(pieces)

// trackerList := resMap["announce-list"].([]any)
// for _, t := range trackerList {
// 	tStrArr, ok := t.([]any)
// 	if !ok {
// 		log.Fatal("something wrong with annouce-list")
// 	}
// 	tBytes, ok := tStrArr[0].([]byte)
// 	if !ok {
// 		log.Fatal("something wrong with annouce-list[inside stuff]")
// 	}
// 	fmt.Println(string(tBytes))
// }
