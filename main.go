package main

import (
	"crypto/sha1"
	"flag"
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
	torrentSize := GetSize(infoDictionary)
	trackers :=  GetTrackers(decodedDictionaryMap)
	// httpTrackers :=  GetHttpTracker(trackers)

	urlCompoents := UrlCompoents{
		tracker: trackers[0],
		infoHash: hash,
		port: 6888,
		left: torrentSize ,

	}

	GenerateAnnounceUrl(urlCompoents)

	// TODO: need to add functionality for udp urls

	UdpRequest(urlCompoents)
}
