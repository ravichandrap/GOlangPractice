package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Println(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\n Hello from tcp server. \n")
		fmt.Fprintln(conn, "How are you ")
		fmt.Fprintln(conn, "Wll, i hope!")

		conn.Close()
	}
}
