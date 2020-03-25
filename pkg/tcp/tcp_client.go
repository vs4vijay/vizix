package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Connect(address string) {
	connection, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error in connecting to server", err)
		return
	}

	handleClientConnection(connection)
}

func handleClientConnection(connection net.Conn) {
	server := connection.RemoteAddr().String()
	fmt.Printf("Connected to Server %s\n", server)

	for {
		serverData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println("Error in reading from server", err)
		}

		sData := strings.TrimSpace(serverData)

		clientData, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error in reading input", err)
		}

		cData := strings.TrimSpace(clientData)

		fmt.Printf("[%s] %s\n", server, string(sData))
		connection.Write([]byte(cData))
	}
	connection.Close()
}
