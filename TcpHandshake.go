package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net"
)

func TcpHandshake(tcpIdWithPort string, urlComp UrlCompoents) (*net.TCPConn,error){

	fmt.Println(tcpIdWithPort)

	tcpAddr,err := net.ResolveTCPAddr("tcp",tcpIdWithPort)
	if err !=nil{
		log.Fatal("tcp addr: "+ err.Error())
	}

	conn,err := net.DialTCP("tcp",nil,tcpAddr)
	if err !=nil{
		fmt.Println("tcp conn failed")
		return nil,err
	}

	buf := make([]byte, 68)
	buf[0] = 19
	copy(buf[1:20], []byte("BitTorrent protocol"))
	// 8 reserved bytes already zero
	copy(buf[28:48], urlComp.infoHash[:])
	copy(buf[48:68], urlComp.peerId[:])

	_, err = conn.Write(buf)
	if err!=nil{
		conn.Close()
		fmt.Println("writing to connec failed")
		return nil,err
	}

	resp := make([]byte, 68)

	_, err = io.ReadFull(conn, resp)
	if err != nil {
		fmt.Println("reading from connec failed")
		conn.Close()
		return nil,err
	}

	if !bytes.Equal(resp[28:48], urlComp.infoHash[:]){
		err := fmt.Errorf("info hash not same")
		conn.Close()
		return nil,err
	}

	return conn,nil
}




func GetTcpConn(trackers []string, hash [20]byte, torrentSize int) (*net.TCPConn, error){
	var peerId [20]byte
	_,err := rand.Read(peerId[:])
	if err != nil {
		log.Fatal(err.Error())
		return nil,nil
	}

	var tcpConn *net.TCPConn
	connected := false

	for _,v := range trackers {

		urlCompoents := UrlCompoents{
			tracker: v,
			infoHash: hash,
			port: 6888,
			left: torrentSize ,
			peerId: peerId,
		}

		ips := UdpRequest(urlCompoents)
		for _,ip := range ips{
			conn,err := TcpHandshake(ip,urlCompoents)

			if err!=nil{
				fmt.Println(err.Error())
				continue
			}
			tcpConn = conn
			connected = true
			break
		}

		if connected {
			break
		}
	}

	if !connected{
		return nil,fmt.Errorf("Tcp connections failed")
	}

	fmt.Println("TCP Connection successful")
	return tcpConn,nil

}
