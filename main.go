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

	infoDictionary, ok := decodedDictionaryMap["info"].(map[string]any)
	if !ok {
		log.Fatal("Invalid info dict. Broken torrent")
	}

	encodedInfoData := en.Encode(infoDictionary)
	hash := sha1.Sum(encodedInfoData)

	// fmt.Println(infoDictionary)
	fmt.Println(GetSize(infoDictionary))
	fmt.Println(hash)
	fmt.Println(GetTrackers(decodedDictionaryMap))

}

// info, ok := resMap["info"].(map[string]any)
// if !ok {
// 	log.Fatal("tch tch info")
// }
// pieces := info["pieces"]
// fmt.Println(pieces)
