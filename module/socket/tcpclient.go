package socket

// tcp client

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"time"
)

func init() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:8003")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Error("Client connect error !", err.Error())
		return
	}

	defer conn.Close()

	log.Info(conn.LocalAddr().String(), ": Client connected!")

	onMessageRecived(conn)

}

func onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	b := []byte(conn.LocalAddr().String() + "Say hello to Server ...\n")
	conn.Write(b)

	for {
		msg, err := reader.ReadString('\n')
		log.Info("ReadString")
		log.Info(msg)

		if err != nil || err == io.EOF {
			log.Info(err)
			break
		}

		time.Sleep(time.Second * 2)

		log.Info("writing...")

		b := []byte(conn.LocalAddr().String() + " write data to Server... \n")
		_, err = conn.Write(b)

		if err != nil {
			log.Error(err)
			break
		}
	}
}
