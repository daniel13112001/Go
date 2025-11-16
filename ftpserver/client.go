package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

func main() {

	// Spawn 50 clients concurrently to send requests to the FTP Server

	//messages := []string{"Hi\n", "Hello\n", "Good morning\n"}

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			conn, err := net.Dial("tcp", ":9181")


			if err != nil {
				panic(err)
			}

			fmt.Println("Client-Server Connection opened successfully")
			defer conn.Close()

			writer := bufio.NewWriter(conn)

			writer.WriteString(("Hello from client\n"))
			writer.Flush()
		}()
	}
	wg.Wait()
}
