package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hallo")
}
*/

func handleConnection(conn net.Conn) {
	buf := bufio.NewWriter(conn)
	buf.WriteString("hallo")
	buf.Flush()
	conn.Close()
}

func main() {
	port := "80"
	//fmt.Println(port)
	//http.HandleFunc("/", hello)
	//http.ListenAndServe(":"+port, nil)
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
