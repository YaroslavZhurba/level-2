package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	fmt.Println("Server started")

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	notify := make(chan error)
	go func() {
		for {
			data, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				notify <- err
				return
			}

			if strings.TrimSpace(data) == "STOP" {
				fmt.Println("Server finished")
				notify <- nil
				return
			}

			fmt.Fprintf(conn, "Server: %s", data)
		}
	}()

	err := <-notify
	if err != nil {
		log.Fatal("connection dropped", err)
	}
	os.Exit(0)
}