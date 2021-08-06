package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	buf := bufio.NewWriter(conn)
	buf.WriteString("hallo")
	buf.Flush()
	conn.Close()
}

func main() {
	port := os.Getenv("PORT")
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("rip")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("rip")
		}
		go handleConnection(conn)
	}
}
