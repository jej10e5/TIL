package main

import (
	"fmt"
	"net"
)

const addr = "localhost:8888"

func main() {
	fmt.Printf("client for server url: %s\n", addr)
	//udp addr 생성
	addr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		panic(err)
	}
	//8888로 연결
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	msg := make([]byte, 512)
	n, err := conn.Write([]byte("connected"))
	if err != nil {
		panic(err)
	}
	for {
		//conn로부터 msg 읽기
		n, err = conn.Read(msg)
		if err != nil {
			continue
		}
		fmt.Printf("%s\n", string(msg[:n]))
	}
}
