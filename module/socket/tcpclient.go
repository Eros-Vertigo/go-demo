package socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func init() {
	// 打开连接
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		// 由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n")
	// 给服务器发送信息知道程序退出
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedClient = strings.Trim(input, "\n")
		if trimmedClient == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedClient))
	}
}
