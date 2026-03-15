package main

import (
	"encoding/binary"
	crand "crypto/rand"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
)


func UdpRequest(urlComponents UrlCompoents){

	a := strings.Split(urlComponents.tracker, "/")
	udpAddr, err := net.ResolveUDPAddr("udp",a[2])
	if err!=nil{
		log.Fatal(err.Error())
	}

	conn,err := net.DialUDP("udp", nil, udpAddr)
	if err!=nil{
		log.Fatal("dial udp: "+err.Error())
	}

	defer conn.Close()

	data := make([]byte,16)

	binary.BigEndian.PutUint64(data[0:8],0x41727101980)
	binary.BigEndian.PutUint32(data[8:12],0)

	transactionId := rand.Uint32()
	binary.BigEndian.PutUint32(data[12:16],transactionId)

	conn.Write(data)

	respData := make([]byte,16)
	_,err = conn.Read(respData)

	if err!=nil{
		log.Fatal("read udp conn: "+err.Error())
	}

	resTransactionId := binary.BigEndian.Uint32(respData[4:8])

	if transactionId != resTransactionId{
		log.Fatal("Transaction ids not same")
	}

	connectionId := binary.BigEndian.Uint64(respData[8:16])

	buf := generateUdpAnnounceRequestData(connectionId, urlComponents)

	conn.Write(buf)

	resp := make([]byte, 1500)
	n, _ := conn.Read(resp)
	peerBytes := resp[20:n]
	for i := 0; i < len(peerBytes); i += 6 {
	    ip := net.IP(peerBytes[i : i+4])
	    port := binary.BigEndian.Uint16(peerBytes[i+4 : i+6])
	    fmt.Println(ip, port)
	}

}

func generateUdpAnnounceRequestData(connectionId uint64, urlComp UrlCompoents)[]byte{

// 	Offset  Size    Name    Value
// 0       64-bit integer  connection_id
// 8       32-bit integer  action          1 // announce
// 12      32-bit integer  transaction_id
// 16      20-byte string  info_hash
// 36      20-byte string  peer_id
// 56      64-bit integer  downloaded
// 64      64-bit integer  left
// 72      64-bit integer  uploaded
// 80      32-bit integer  event           0 // 0: none; 1: completed; 2: started; 3: stopped
// 84      32-bit integer  IP address      0 // default
// 88      32-bit integer  key
// 92      32-bit integer  num_want        -1 // default
// 96      16-bit integer  port
// 98

	var peerId [20]byte
	_,err := crand.Read(peerId[:])
	if err != nil {
		log.Fatal(err.Error())
	}
	transactionId := rand.Uint32()
	buf := make([]byte,98)

	binary.BigEndian.PutUint64(buf[0:8],connectionId)
	binary.BigEndian.PutUint32(buf[8:12],1)
	binary.BigEndian.PutUint32(buf[12:16],transactionId)

	copy(buf[16:36], urlComp.infoHash[:])
	copy(buf[36:56], peerId[:])

	binary.BigEndian.PutUint64(buf[56:64],0)
	binary.BigEndian.PutUint64(buf[64:72],uint64(urlComp.left))
	binary.BigEndian.PutUint64(buf[72:80],0)
	binary.BigEndian.PutUint32(buf[80:84],0)
	binary.BigEndian.PutUint32(buf[84:88],0)
	binary.BigEndian.PutUint32(buf[88:92],rand.Uint32())
	binary.BigEndian.PutUint32(buf[92:96],0xFFFFFFFF)
	binary.BigEndian.PutUint16(buf[96:98],uint16(urlComp.port))

	return buf

}
