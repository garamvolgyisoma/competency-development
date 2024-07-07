package main

import (
	"fmt"
	"net"
)

func createServer() {
	listener, err := net.Listen("tcp", "localhost:1420")
	if err != nil { // nil is go's null
		fmt.Println("Error creating listener:", err)
		return
	}
	defer listener.Close() // defer: defers the execution until the surrounding function (createServer in this case) returns, ideally it won't happen here (only if error occours)
	fmt.Println("Server is running on localhost:1420")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		go handleConnection(conn) // concurrency stuff: starts a goroutine -> a function that's capable of running concurrently with other functions
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024) // make: allocates memory and initializes the underlying structure -> returned value is ready to use right away
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:", err)
		return
	}

	response := "HTTP/1.1 200 OK\n" +
		"Content-Type: text/html\n" +
		"Content-Length: 208\n" +
		"\n" +
		"<html><head><style>body { background-image: url('https://steamuserimages-a.akamaihd.net/ugc/34105835836189690/243E08990753A1958AE0057BE02DDA842EFDC5DF/?imw=5000&imh=5000&ima=fit&impolicy=Letterbox&imcolor=%23000000&letterbox=false'); background-size: cover; }</style></head></html>"

	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}

	err = conn.(*net.TCPConn).CloseWrite()
	if err != nil {
		fmt.Println("Error shutting down connection:", err)
		return
	}
}

func main() {
	createServer()
}
