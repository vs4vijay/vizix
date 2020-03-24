package tcp_server

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
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
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	client := connection.RemoteAddr().String()
	fmt.Printf("Client Connected %s\n", client)

	connection.Write([]byte(welcomeMessage()))

	for {
		clientData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println("Error in reading from client", err)
		}

		cData := strings.TrimSpace(clientData)

		serverData, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error in reading input", err)
		}

		sData := strings.TrimSpace(serverData)

		fmt.Printf("[%s] %s\n", client, string(cData))
		connection.Write([]byte(sData))
	}
	connection.Close()
}

func welcomeMessage() string {
	return "Welcome to VIZIX\n"
}
