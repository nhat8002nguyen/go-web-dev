package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	defer listener.Close()

	if err != nil {
		log.Panic(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scann := bufio.NewScanner(conn)
	for scann.Scan() {
		text := scann.Text()
		fmt.Println(text)
	}
	defer conn.Close()

	fmt.Println("Code right here!")
}