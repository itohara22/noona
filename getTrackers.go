package main

import (
	"log"
)

func GetTrackers(decodedDict map[string]any) []string {
	trackers := []string{}
	trackerList,ok := decodedDict["announce-list"].([]any)
	if ok {
		for _, t := range trackerList {
			tStrArr, ok := t.([]any)
			if !ok {
				log.Fatal("Invalid trackers")
			}
			tBytes, ok := tStrArr[0].([]byte)
			if !ok {
				log.Fatal("Invalid trackers")
			}
			a := string(tBytes)
			trackers = append(trackers, a)
		}
	}

	trackerAnnounce,ok := decodedDict["announce"].([]byte)
	if !ok{
		log.Fatal("announce broken")
	}
	trackers = append(trackers,string(trackerAnnounce))

	return trackers
}
