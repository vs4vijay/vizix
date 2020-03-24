package tcp_server

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func Start(port string) {
	port = fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp4", port)
	if err != nil {
		fmt.Println("Error in starting server", err)
		return
	}
	defer listener.Close()

	rand.Seed(time.Now().Unix()) // CHECK

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error in accepting connection", err)
			return
		}
		handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	client := connection.RemoteAddr().String()
	fmt.Printf("Client Connected %s", client)

	for {
		netData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println("Error in reading data", err)
		}

		data := strings.TrimSpace(netData)

		fmt.Printf("[%s] %s\n", client, data)
		connection.Write([]byte("VIZIX\n"))
	}
	connection.Close()
}
