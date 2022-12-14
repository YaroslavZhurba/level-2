package grep

import (
	"fmt"
	"strings"
	f "dev05/flags"
)

func findMax(a, b uint) uint {
	if a > b {
		return a
	}

	return b
}

func copyKeys(m map[int]string) []int {
	c := make([]int, 0)
	for k := range m {
		c = append(c, k)
	}

	return c
}

func Grep(pattern string, input []string, params f.Flags) map[int]string {
	if params.IgnoreCase {
		pattern = strings.ToLower(pattern)
	}

	res := make(map[int]string)
	for i, v := range input {
		str := v
		if params.IgnoreCase {
			str = strings.ToLower(str)
		}

		var found bool
		if params.Fixed {
			found = str == pattern
		} else {
			found = strings.Contains(str, pattern)
		}

		if !found == params.Invert {
			res[i] = v
		}
	}

	if params.Count {
		fmt.Println(len(res))
		return nil
	}

	before := findMax(params.Before, params.Context)
	after := findMax(params.After, params.Context)

	getStringsAroundFound(res, input, before, after)

	return res
}

func getStringsAroundFound(m map[int]string, input []string, before uint, after uint) {
	c := copyKeys(m)

	for _, v := range c {
		for j := 1; j <= int(before); j++ {
			if v-j < 0 {
				continue
			}

			m[v-j] = input[v-j]
		}

		for j := 1; j <= int(after); j++ {
			if v+j > len(input)-1 {
				continue
			}

			m[v+j] = input[v+j]
		}
	}
}