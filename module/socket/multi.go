package socket

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

const maxRead = 25

func initMulti() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(hostAndPort)
	for {
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		go connectionHandler(conn)
	}
}

func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address:port failed: '"+hostAndPort+"'")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP: ")
	println("Listening to: ", listener.Addr().String())
	return listener
}

func connectionHandler(conn net.Conn) {
	connForm := conn.RemoteAddr().String()
	println("Connection from: ", connForm)
	sayHello(conn)
	for {
		var buff = make([]byte, maxRead+1)
		length, err := conn.Read(buff[0:maxRead])
		buff[maxRead] = 0 // to prevent overflow
		switch err {
		case nil:
			handleMsg(length, err, buff)
		case syscall.EWOULDBLOCK:
			continue
		default:
			goto DISCONNECT
		}
	}
DISCONNECT:
	err := conn.Close()
	println("Closed connection: ", connForm)
	checkError(err, "Close: ")
}

func sayHello(to net.Conn) {
	buff := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := to.Write(buff)
	checkError(err, "Write: wrote "+string(wrote)+" bytes.")
}

func handleMsg(length int, err error, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0; ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
		}
		print(">")
	}
}

func checkError(err error, info string) {
	if err != nil {
		panic("ERROR: " + info + " " + err.Error()) // terminate program
	}
}
