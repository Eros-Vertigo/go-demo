package socket

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"time"
)

func init() {
	setupServer()
}

func setupServer() {
	ln, err := net.Listen("tcp", ":8004")
	if err != nil {
		log.Error(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Error(err)
		}
		go handleConnection(conn)
		//go handleClient(conn)
	}
}

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	b := []byte(conn.LocalAddr().String() + "Say hello to Server...\n")
	conn.Write(b)
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println("ReadString")
		fmt.Println(msg)

		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second * 2)
		fmt.Println("writing...")

		b := []byte(conn.LocalAddr().String() + "write data to Server...\n")
		_, err = conn.Write(b)

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
