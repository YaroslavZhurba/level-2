package main

import (
	"fmt"
	"strings"
	"dev04/anagram"
)

func main() {
	input := []string{"ПЯТКА", "ПЯТАК", "ТЯПКА", "ЛИСТОК", "СЛИТОК", "СТОЛИК", "кАРА", "арка", "кара","рак", "КаРа"}
	m := anagram.MakeAnagrams(input)
	fmt.Printf("input: %s\n\n", strings.Join(input, ", "))
	for k, v := range m {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ", "))
	}
}