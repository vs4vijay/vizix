package tcp

import (
	"bufio"
	"fmt"
	"io"
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
		go handleServerConnection(connection)
		//c := make(chan os.Signal)
		//signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		//<-c
	}
}

func handleServerConnection(connection net.Conn) {
	client := connection.RemoteAddr().String()
	fmt.Printf("[%s] +Connected\n", client)

	connection.Write([]byte(welcomeMessage()))

	for {
		clientData, err := bufio.NewReader(connection).ReadString('\n')

		if err == io.EOF {
			fmt.Printf("[%s] -Disconnected\n", client)
			break
		} else if err != nil {
			fmt.Println("Error in reading from client", err)
		}

		cData := strings.TrimSpace(clientData)

		serverData, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error in reading input", err)
		}

		sData := strings.TrimSpace(serverData)

		if len(sData) != 0 {
			fmt.Printf("[%s] %s\n", client, string(cData))
			connection.Write([]byte(sData))
		}
	}
	connection.Close()
}

func welcomeMessage() string {
	return "Welcome to VIZIX\n"
}
