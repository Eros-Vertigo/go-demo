package socket

// tcp server

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

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	log.Info("Server ready to read ...")

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Error("accept error :", err)
			continue
		}
		log.Info("A client connected : ", tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()

	defer func() {
		log.Info(" Disconnected : ", ipStr)
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	i := 0

	for {
		message, err := reader.ReadString('\n') // 将数据按照换行符进行读取
		if err != nil || err == io.EOF {
			break
		}

		log.Info(string(message))

		time.Sleep(time.Second * 3)

		msg := time.Now().String() + conn.RemoteAddr().String() + "Server Say hello! \n"
		b := []byte(msg)

		conn.Write(b)
		i++

		if i > 10 {
			break
		}
	}
}
