package telnet

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Client struct {
	Conn net.Conn
}

func NewTelnetClient(host string, port string, timeout int) (*Client, error) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Duration(timeout)*time.Second)

	for i := 0; i < timeout; i++ {
		time.Sleep(time.Second)
		conn, err = net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Duration(timeout)*time.Second)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, err
	}

	err = conn.(*net.TCPConn).SetKeepAlive(true)
	if err != nil {
		return nil, err
	}

	err = conn.(*net.TCPConn).SetKeepAlivePeriod(30 * time.Second)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected")

	return &Client{Conn: conn}, nil
}

func (c *Client) Start() {
	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	go handle(c.Conn, errChan)

	for {
		select {
		case <-sigChan:
			fmt.Println("\n Client has been terminated")
			return
		case err := <-errChan:
			log.Printf("got error: %s\n", err)
			return
		}
	}
}

func handle(conn net.Conn, errChan chan error) {
	std_reader := bufio.NewReader(os.Stdin)
	server_reader := bufio.NewReader(conn)
	for {
		fmt.Print("-> ")
		text, err := std_reader.ReadString('\n')
		if err != nil {
			errChan <- err
			return
		}

		fmt.Fprintf(conn, text+"\n")
		if text == "STOP\n" {
			fmt.Println("Client has been terminated")
			return
		} 

		message, err := server_reader.ReadString('\n')
		if err != nil {
			fmt.Println(text)
			errChan <- err
			return
		}

		fmt.Print("<-: " + message)
	}
}