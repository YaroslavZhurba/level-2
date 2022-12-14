package sortlines

import (
	"bufio"
	"sort"
	"strconv"
	"unicode"
	"strings"
	"reflect"
	"fmt"
	f "dev03/flags"
	"log"
	"os"
)

func readLines(path string) ([]string, error) {
	var result []string
	file, err := os.Open(path)
    if err != nil {
        return result, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        result = append(result, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return result, nil
}

func removeLeadSpaces(s string) string {
	runes := []rune(s)
	limit := len(runes) - 1
	i := 0
	for {
		if i > limit {
			return ""
		}
		if unicode.IsSpace(runes[i]) {
			i++
			continue
		}
		break
	}
	return string(runes[i:])
}

func lessStringColumnComparator(lines [][]string, isLeadSpaces bool, column int) (func (i, j int) bool) {
	column_index := column - 1
	return func(i, j int) bool {
		var s1, s2 string
		if isLeadSpaces {
			s1 = removeLeadSpaces(lines[i][column_index])
			s2 = removeLeadSpaces(lines[j][column_index])
		} else {
			s1 = lines[i][column_index]
			s2 = lines[j][column_index]
		}
		return s1 < s2
	}
}

func lessStringComparator(lines []string, isLeadSpaces bool) (func (i, j int) bool) {
	return func(i, j int) bool {
		var s1, s2 string
		if isLeadSpaces {
			s1 = removeLeadSpaces(lines[i])
			s2 = removeLeadSpaces(lines[j])
		} else {
			s1 = lines[i]
			s2 = lines[j]
		}
		return s1 < s2
	}
}

func lessNumber(sa, sb string) bool {
	a := getNumber(sa)
	b := getNumber(sb)

	return a < b
}

func lessNumberComparator(lines []string) (func (i, j int) bool) {
	return func(i, j int) bool {
		s1 := lines[i]
		s2 := lines[j]
		return lessNumber(s1, s2)
	}
}

func lessNumberColumnsComparator(lines [][]string, column int) (func (i, j int) bool) {
	column_index := column - 1
	return func(i, j int) bool {
		s1 := lines[i][column_index]
		s2 := lines[j][column_index]
		return lessNumber(s1, s2)
	}
}

func getNumber(s string) int {
	s_no_spaces := removeLeadSpaces(s)
	runes := []rune(s_no_spaces)
	
	i := 0
	for i < len(runes) && unicode.IsDigit(runes[i]) {
		i++
	}
	number, _ := strconv.Atoi(string(runes[:i]))
	return number
}

func isNumber(s string) bool {
	s_no_spaces := removeLeadSpaces(s)
	runes := []rune(s_no_spaces)
	if s_no_spaces == "" || !unicode.IsDigit(runes[0]){
		return false
	}
	return true
}

func sortSplit(s string) ([]string) {
	result := make([]string, 0)
	flag := false
	part := ""
	for _, c := range s {
		if unicode.IsSpace(c) {
			if (flag) {
				result = append(result, part)
				flag = false
				part = string(c)
			} else {
				part += string(c)
			}
		} else {
			flag = true
			part += string(c)
		}
	}
	if part != "" {
		result = append(result, part)
	}
	return result
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func makeUnique(lines []string) []string {
	map_s := make(map[string]bool, 0)
	result := make([]string, 0)

	for _, line := range lines {
		_, contains := map_s[line]
		if contains {
			continue
		}
		result = append(result, line)
		map_s[line] = true
	}
	
	return result
}


func Sort(flags f.Flags, path string) ([]string, error) {
	original_lines, err := readLines(path)
	if err != nil {
		return make([]string, 0), err
	}

	lines := make([]string, 0)
	columns_splited := make([][]string, 0)
	numbers := make ([]string, 0)
	if flags.Column > 0 {
		if flags.Number {
			for _, line := range original_lines {
				line_splited := sortSplit(line)
				if len(line_splited) < flags.Column || 
				  !isNumber(line_splited[flags.Column - 1]) {
					lines = append(lines, line)
				} else {
					columns_splited = append(columns_splited, line_splited)
				}
			}
			// make sort here
			sort.SliceStable(lines, lessStringComparator(lines, false))
			sort.SliceStable(columns_splited, 
			  lessNumberColumnsComparator(columns_splited, flags.Column))
			for _, line_splited := range(columns_splited) {
				lines = append(lines, strings.Join(line_splited, ""))
			}
		} else {
			if flags.LeadSpace {
				for _, line := range original_lines {
					line_splited := sortSplit(line)
					if len(line_splited) < flags.Column ||
					  removeLeadSpaces(line_splited[flags.Column - 1]) == ""{
						lines = append(lines, line)
					} else {
						columns_splited = append(columns_splited, line_splited)
					}
				}
				sort.SliceStable(lines, lessStringComparator(lines, false))
				sort.SliceStable(columns_splited, 
				  lessStringColumnComparator(columns_splited, flags.LeadSpace, flags.Column))
				for _, line_splited := range(columns_splited) {
					lines = append(lines, strings.Join(line_splited, ""))
				}
			} else {
				for _, line := range original_lines {
					line_splited := sortSplit(line)
					if len(line_splited) < flags.Column {
						lines = append(lines, line)
					} else {
						columns_splited = append(columns_splited, line_splited)
					}
				}
				sort.SliceStable(lines, lessStringComparator(lines, false))
				sort.SliceStable(columns_splited, 
				  lessStringColumnComparator(columns_splited, flags.LeadSpace, flags.Column))
				for _, line_splited := range(columns_splited) {
					lines = append(lines, strings.Join(line_splited, ""))
				}
			}
		}
	} else {
		if flags.Number {
			for _, line := range original_lines {
				if  !isNumber(line) {
					lines = append(lines, line)
				} else {
					numbers = append(numbers, line)
				}
			}
			sort.SliceStable(lines, lessStringComparator(lines, flags.LeadSpace))
			sort.SliceStable(numbers, 
			  lessNumberComparator(numbers))
			lines = append(lines, numbers...)
		} else {
			lines = make([]string, len(original_lines))
			copy(lines, original_lines)
			sort.SliceStable(lines, lessStringComparator(lines, flags.LeadSpace))
		}
	}

	if flags.IsSorted {
		if reflect.DeepEqual(original_lines, lines) {
			fmt.Println("Sorted")
		} else {
			fmt.Println("Not sorted")
		}
	}

	if flags.Reverse {
		reverse(lines)
	}

	if flags.Unique {
		lines = makeUnique(lines)
	}
	
	return lines, nil
}