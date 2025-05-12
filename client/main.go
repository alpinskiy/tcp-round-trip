package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	tcp4LoopbackRTT()
}

func tcp4LoopbackRTT() {
	var procAttr os.ProcAttr
	srv, err := os.StartProcess("server", nil, &procAttr)
	if err != nil {
		log.Fatal(err)
	}
	defer srv.Kill()
	addr, err := net.ResolveTCPAddr("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	var conn *net.TCPConn
	for {
		conn, err = net.DialTCP("tcp", nil, addr)
		if err == nil {
			closeOnInterrupt(conn)
			break
		}
		time.Sleep(0)
	}
	log.Println("started")
	buf := make([]byte, 1)
	num := time.Duration(0)
	start := time.Now()
	for {
		if _, err := conn.Write(buf); err != nil {
			break
		}
		num++
	}
	elapsed := time.Since(start)
	log.Println("stopped")
	rtt := elapsed / num
	rps := int(float64(num) / elapsed.Seconds())
	fmt.Println()
	fmt.Printf("RTT %v\n", rtt)
	fmt.Printf("RPS %d\n", rps)
}

func closeOnInterrupt(conn *net.TCPConn) {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	go func() {
		<-sigc
		if conn != nil {
			conn.Close()
		}
	}()
}
