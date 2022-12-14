package main

import (
	w "dev09/wget"
	"flag"
	"fmt"
	"os"
)

func main() {
	var destination string
	flag.StringVar(&destination, "d", "./data", "destination for output")
	flag.Parse()
	fmt.Println("Downloading...")
	// args := flag.Args()
	// if len(args) != 1 {
	// 	fmt.Println("There is should be exactly one argument")
	// 	os.Exit(0)
	// }

	// err := w.RunWget(destination, args[0])
	w.RunWget(destination, "https://lichess.org")
	fmt.Println("Finished")
	os.Exit(0)
}
