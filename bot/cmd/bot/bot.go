package main

import (
	"bot/internals/socket"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conn, err := socket.ConnectToServer("192.168.1.101:8080")
	if err != nil {
		fmt.Println("Socket connection failed:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Socket connection established.")

	go socket.ReadFromServer(conn)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Press Ctrl+C to exit...")
	<-sig

	fmt.Println("Exiting now.")
}
