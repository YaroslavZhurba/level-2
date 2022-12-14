package unpack

import (
	"fmt"
	"unicode"
	"strings"
	"strconv"
)

func getNumber(runes []rune, i int) (int, int, bool) {
	length := len(runes)
	if i >= length || !unicode.IsDigit(runes[i]) {
		return 0, i, false
	}
	var number_runes strings.Builder
	number_runes.WriteRune(runes[i])
	i++
	for i < length {
		if !unicode.IsDigit(runes[i]) {
			break
		}
		number_runes.WriteRune(runes[i])
		i++
	}
	number, _ := strconv.Atoi(number_runes.String())
	return number, i, true
}

func Unpack(str string) (string, error) {
	runes := []rune(str)
	length := len(runes)
	i := 0

	var result strings.Builder

	if length <= 0 {
		return "",  nil
	}
	if unicode.IsDigit(runes[0]) {
		return "", fmt.Errorf("unexpected digit in position %d", i+1)
	}

	for i < length {
		current_rune := runes[i]
		if string(current_rune) == "\\" {
			i++
			if i >= length {
				return "", fmt.Errorf("unexpected end of string in position %d", i+1)
			} else {
				current_rune = runes[i]
			}
		}
		i++
		number, new_i, isNumber := getNumber(runes, i)
		i = new_i

		if !isNumber {
			result.WriteRune(current_rune)
			continue
		}
		result.WriteString(strings.Repeat(string(current_rune), number))
	}
	return result.String(), nil
}