package main

import (
	"crypto/sha1"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
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
	udpTrackers := GetUdpTracker(trackers)


	conn,err := GetTcpConn(udpTrackers, hash, torrentSize)
	if err !=nil{
		log.Fatal("tcp connections failed")
	}
	defer conn.Close()

	// send interested after handshake
	interestedMsg := make([]byte,5)
	binary.BigEndian.PutUint32(interestedMsg[0:4],1)
	interestedMsg[4] = 2

	_, err = conn.Write(interestedMsg)
	if err != nil {
			log.Fatal(err.Error())
	}

	fmt.Println("sent interested")

	for {
			lenBuf := make([]byte, 4)

			_, err = io.ReadFull(conn, lenBuf)
			if err != nil {
				log.Fatal(err.Error())
			}

			length := binary.BigEndian.Uint32(lenBuf)

			if length == 0 {
				fmt.Println("0 length keep-alive")
				continue
			}

			msg := make([]byte, length)

			_, err = io.ReadFull(conn, msg)
			if err != nil {
				log.Fatal(err.Error())
			}

			msgID := msg[0]
			payload := msg[1:]

			if msgID != 1 {
				fmt.Println(string(payload))
				log.Fatal("connection not unchocked by peer")
			}
	}

}
