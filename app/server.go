package main

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/resp"
	"net"
	"os"
	"strings"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	resp := resp.NewResp(conn)
	for {
	value, err := resp.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(value)
	if value.Typ == "array" && len(value.Array) > 0 {
		command := strings.ToUpper(value.Array[0].Bulk)
		switch command {
		case "PING":
			conn.Write([]byte("+PONG\r\n"))
		case "ECHO":
			if len(value.Array) < 2 {
				conn.Write([]byte("-ERR wrong number of arguments for 'echo' command\r\n"))
			} else {
				response := fmt.Sprintf("$%d\r\n%s\r\n", len(value.Array[1].Bulk), value.Array[1].Bulk)
				conn.Write([]byte(response))
			}
		default:
			conn.Write([]byte("-ERR unknown command '" + command + "'\r\n"))
		}
	} else {
		conn.Write([]byte("-ERR invalid request\r\n"))
	}
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
