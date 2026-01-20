package main

import "log"

func GetSize(info map[string]any) int {

	//for single file torrent
	length, ok := info["length"].(int)
	if ok {
		return length
	}

	// multi file torrent
	files, ok := info["files"].([]any) // in files array each file have a length
	if !ok {
		log.Fatal("Invalid torrent length")
	}

	total := 0
	for _, f := range files {
		file, ok := f.(map[string]any)
		if !ok {
			log.Fatal("Broken torrent lenghts")
		}
		length, exist := file["length"]
		if !exist {
			log.Fatal("Broken torrent lenghts")
		}
		total += length.(int)
	}
	return total
}
