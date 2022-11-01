package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const addr = "localhost:8888" //8888포트로 주소 정해놓음

func echoBackCapitalized(conn net.Conn) {
	//conn에 리더(reader)를 설정한다(io.Reader)
	reader := bufio.NewReader(conn)

	//읽어온 데이터의 첫줄 가져오기
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("error reading data: %s\n", err.Error())
		return
	}
	//출력한 다음 데이터를 다시 보낸다.
	fmt.Printf("Received: %s", data)
	conn.Write([]byte(strings.ToUpper(data))) //읽은 데이터 대문자로 만들기
	//완료된 연결을 종료한다.
	conn.Close()
}

func main() {
	//addr 주소로 tcp 통신 서버를 연다
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	//끝나기 전에 서버 닫기
	defer ln.Close()
	fmt.Printf("listening on: %s\n", addr)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("encounted an error accepting connection: %s\n", err.Error())
			//err가 있으면 다시 한다.
			continue
		}
		//이 작업을 비동기로 처리하면
		//잠재적으로 워커 풀을 위해
		//좋은 사용 사례가 될 것 이다.
		go echoBackCapitalized(conn)
	}
}
