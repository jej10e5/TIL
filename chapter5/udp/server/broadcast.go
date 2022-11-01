package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type connections struct {
	addrs map[string]*net.UDPAddr
	//맵의 수정을 위해 락 시킨다.
	mu sync.Mutex
}

func broadcast(conn *net.UDPConn, conns *connections) {
	count := 0
	for {
		count++
		conns.mu.Lock()
		//알려진(확인한) 주소에 대해 루프로 반복 처리한다.
		for _, retAddr := range conns.addrs {
			//모두에서 메시지를 전송한다.
			msg := fmt.Sprintf("Sent %d", count)
			//msg를 모든 주소로 보내기
			if _, err := conn.WriteToUDP([]byte(msg), retAddr); err != nil {
				fmt.Printf("error encounted: %s", err.Error())
				continue
			}
		}
		conns.mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}
