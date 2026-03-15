package main

import (
	"crypto/rand"
	"log"
	"net/url"
	"strconv"
)

type UrlCompoents struct {
	tracker string;
	infoHash [20]byte;
	port int;
	left int;
}

func GenerateAnnounceUrl(compoents UrlCompoents) string {

	var peerId [20]byte
	_,err := rand.Read(peerId[:])
	if err != nil {
		log.Fatal(err.Error())
	}

	baseUrl,err := url.Parse(compoents.tracker)

	if err != nil {
		log.Fatal(err.Error())
	}

	q := baseUrl.Query()

	q.Add("info_hash",string(compoents.infoHash[:]))
	q.Add("port", strconv.Itoa(compoents.port))
	q.Add("left", strconv.Itoa(compoents.left))
	q.Add("compact", "1")
	q.Add("peer_id",string(peerId[:]))
	q.Add("uploaded","0")
	q.Add("downloaded","0")

	resUrl := baseUrl.String()+"?"+q.Encode()
	return resUrl
}
