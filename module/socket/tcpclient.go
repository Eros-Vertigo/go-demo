package socket

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

func init() {
	//fmt.Println("tcpClient")
	//time.Sleep(5 * time.Second)
	//setupClient()
}

func setupClient() {
	fmt.Println("setupClient")
	conn, err := net.Dial("tcp", ":8004")
	if err != nil {
		log.Error(err)
		return
	}
	fmt.Fprintf(conn, "GET /HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Error(err)
	}
	fmt.Println(status)
}
