package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Fail to connect to tcp server!")
	}
	defer conn.Close()

	io.WriteString(conn, "Hello, how are you!")
}