package main

import (
	"fmt"
	"net"
	"strings"
)

const (
	connType = "tcp"
	connHost = "localhost"
	connPort = "3333"
)

func main() {

	l, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		panic("Listen() failed: " + err.Error())
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Accept() failed: " + err.Error())
			continue
		}

		go handleReading(c)
	}
}

func handleReading(conn net.Conn) {
	defer conn.Close()

	fmt.Println("handleReading: " + conn.RemoteAddr().String())

	for {
		b := make([]byte, 1024)

		n, err := conn.Read(b)
		if err != nil {
			fmt.Println("Read() failed: " + err.Error())
			break
		}

		if n > 0 {
			fmt.Printf("Read(): %v", string(b))

			var bldr strings.Builder
			bldr.Write([]byte("Resent: "))
			bldr.Write([]byte(bldr.String()))
		}
	}

	fmt.Println("closing gorouting for handleReading: " + conn.RemoteAddr().String())
}
