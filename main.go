package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileData, err := os.ReadFile("torrent_files/a.torrent")
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

	// en := NewBencodeEncoder()
	// enCodedData := en.Encode(resMap["info"])
	// // for _, v := range enCodedData {
	// // 	fmt.Print(v)
	// // }

	// hash := sha1.Sum(enCodedData)
	// fmt.Print(hash)

	// info, ok := resMap["info"].(map[string]any)
	// if !ok {
	// 	log.Fatal("tch tch info")
	// }
	// pieces := info["pieces"]
	// fmt.Println(pieces)

	//
	// list := resMap["announce-list"].([]string) // we cannt assert to []string
	// list := resMap["announce-list"].([]any)
	// for _, val := range list {
	// 	fmt.Println(val)
	// }

	// val := map[string]any{
	// 	"cow":  "moo",
	// 	"spam": "eggs",
	// }
	// val := []any{
	// 	"spam",
	// 	"eggs",
	// 	42,
	// }
}

// dic := "d3:cow3:moo4:spam4:eggs3:numi42e4:listl4:test3:one3:twoe3:fooi99ee"
