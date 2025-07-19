package socketserver

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	clients      = make(map[net.Conn]struct{})
	clientsMutex sync.Mutex
)

func SocketServer(port int) {
	SPORT := ":" + strconv.Itoa(port)

	listener, err := net.Listen("tcp", SPORT)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	go handleConsoleInput()

	fmt.Println("Server is listening on port", SPORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("New client connected:", conn.RemoteAddr())

		clientsMutex.Lock()
		clients[conn] = struct{}{}
		clientsMutex.Unlock()

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer func() {
		clientsMutex.Lock()
		delete(clients, conn)
		clientsMutex.Unlock()
		conn.Close()
		fmt.Println("Client disconnected:", conn.RemoteAddr())
	}()

	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			return
		}
	}
}

func handleConsoleInput() {
	for {
		text, err := readConsoleInput()
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		if strings.TrimSpace(text) == "c" {
			printClientCount()
		} else {
			broadcastMessage(text)
		}
	}
}

func readConsoleInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	return text, err
}

func printClientCount() {
	clientsMutex.Lock()
	count := len(clients)
	clientsMutex.Unlock()
	fmt.Printf("Connected clients count: %d\n", count)
}

func broadcastMessage(message string) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	for client := range clients {
		_, err := client.Write([]byte(message))
		if err != nil {
			fmt.Println("Write error to", client.RemoteAddr(), ":", err)
		}
	}
}
