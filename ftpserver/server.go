package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Start TCP server and listen for requests on port 9181

	listener, err := net.Listen("tcp", ":9181")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server listening on TCP port: 9181")

	// Accept Connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	}
	fmt.Println(line)
}
