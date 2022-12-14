package main

import (
	c "dev06/cut"
	f "dev06/flags"
	"flag"
)

var flags f.Flags  

func init() {
	flags = f.Flags{}
	flag.StringVar(&flags.Columns, "f", "1", "Field to select")
	flag.StringVar(&flags.Delimiter, "d", "\t", "Delimeter. Tab is by default.")
	flag.BoolVar(&flags.Separated, "s", false, "Print only string contains delimeter")
}

func main() {
	flag.Parse()
	c.Cut(flags)
}