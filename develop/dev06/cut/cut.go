package cut

import (
	"bufio"
	f "dev06/flags"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type nums_t map[int]struct{}

func parseFields(s string) nums_t {
	if s == "" {
		log.Fatal("Flag -f can't be empty")
	}

	result := make(nums_t)
	nums := strings.Split(s, ",")
	
	for i, v := range nums {
		number, err := strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Error fields with args number %d\n", i + 1)
			log.Fatal(err)
		}
		result[number] = struct{}{}
	}

	return result
}

func cutPrint(splited_line []string, delimiter string) {
	for _, s := range splited_line {
		fmt.Printf("%s%s",s,delimiter)
	}
	fmt.Println()
}


func splitLinesByDelimiter(line, delimiter string) []string {
	return strings.Split(line, delimiter)
}

func getColumnsValue(splited_line []string, numbers_set nums_t) []string {
	result := make([]string, 0)
		for num := range numbers_set {
			if num > 0 && num <= len(splited_line) {
				result = append(result, splited_line[num-1])
			}
		}
	return result
}

func Cut(flags f.Flags) {
	fmt.Println()
	numbers_set := parseFields(flags.Columns)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if line == "\\q\n" {
			break
		}

		if flags.Separated {
			if !strings.Contains(line, flags.Delimiter) {
				continue
			}
		}

		splited_line := splitLinesByDelimiter(line, flags.Delimiter)

		cutPrint(getColumnsValue(splited_line, numbers_set), flags.Delimiter)
	}
}

