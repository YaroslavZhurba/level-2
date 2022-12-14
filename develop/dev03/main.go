package main

import (
	f "dev03/flags"
	s "dev03/sort_lines"
	"flag"
	"fmt"
	"log"
)

var flags f.Flags

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func init() {
	flags = f.Flags{}

	flag.IntVar(&flags.Column, "k", 0, "Sort by column number")
	flag.BoolVar(&flags.Number, "n", false, "Sort by numbers")
	flag.BoolVar(&flags.Reverse, "r", false, "Sort in a reverse order")
	flag.BoolVar(&flags.Unique, "u", false, "Don't print dublicate lines")
	flag.BoolVar(&flags.LeadSpace, "b", false, "Ignore leading spaces")
	flag.BoolVar(&flags.IsSorted, "c", false, "Check if file is already sorted")
}

func main() {
	flag.Parse()

	lines, err := s.Sort(flags, flag.Arg(0))
	// lines, err := s.Sort(flags, "sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	printLines(lines)
}
