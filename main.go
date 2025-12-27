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

	decodedDictionaryMap, ok := res.(map[string]any)
	if !ok {
		log.Fatal("Corrupt torrent file")
	}

	en := NewBencodeEncoder()

	if decodedDictionaryMap["info"] == nil {
		log.Fatal("Corrupt torrent file")
	}

	encodedInfoData := en.Encode(decodedDictionaryMap["info"])

	hash := sha1.Sum(encodedInfoData)
	fmt.Print(hash)

}

// info, ok := resMap["info"].(map[string]any)
// if !ok {
// 	log.Fatal("tch tch info")
// }
// pieces := info["pieces"]
// fmt.Println(pieces)
