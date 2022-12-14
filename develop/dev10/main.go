package main

import (
	"flag"
	"log"
	"dev10/telnet"
)

type Flags struct {
	host string
	port string
	timeout int
}

var flags Flags

func init() {
	flags = Flags{}
	flag.StringVar(&flags.host, "host", "localhost", "host name")
	flag.StringVar(&flags.port, "port", "3000", "port name")
	flag.IntVar(&flags.timeout, "timeout", 10, "timeout for connection")
}

func main() {
	flag.Parse()
	client, err := telnet.NewTelnetClient(flags.host, flags.port, flags.timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Conn.Close()

	client.Start()
}