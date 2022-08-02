package socket

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"time"
)

func init() {
	//setupClient()
}

func setupClient() {
	time.Sleep(5 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8004")
	if err != nil {
		log.Error(err)
	}
	fmt.Fprintf(conn, "GET /HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Error(err)
	}
	fmt.Println(status)
}
