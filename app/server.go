package main

import (
	"fmt"
	"os"

	"net"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	// read datasource, init reader

	for i := 0; i < 2; i++ {
		// write
		_, err = conn.Write([]byte("+PONG\r\n"))

		if err != nil {
			fmt.Println("Error write msg: ", err.Error())
			os.Exit(1)
		}
	}
}
