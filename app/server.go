package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 128)

	for {
		n, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		log.Printf("commands: \n%s", buf[:n])
		conn.Write([]byte("+PONG\r\n"))
	}
}

func main() {
	fmt.Println("Logs from your program will appear here!")
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		if err != nil {
			fmt.Println("Error parsing command: ", err.Error())
			os.Exit(1)
		}
		go handleConn(conn)

	}

}
