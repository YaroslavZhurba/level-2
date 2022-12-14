package main

import (
	unpack "dev02/unpack"
	"fmt"
	"os"
)

func main() {
	s := `\15qwerty\1\2\3`
	s_unpacked, err := unpack.Unpack(s)
	if err != nil {
		fmt.Println(s)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(s)
	fmt.Println(s_unpacked)

}
