package anagram

import (
	"sort"
	"strings"
)

type mapAnagrams_t map[string][]string
type setStrings_t map[string]struct{}

func makeAnagramsSortedLettersKey(input []string) mapAnagrams_t {
	result := make(mapAnagrams_t)
	usedWords := make(setStrings_t)

	for _, word := range input {
		word = strings.ToLower(word)
		letters := strings.Split(word, "")
		sort.Strings(letters)
		sortedWord := strings.Join(letters, "")

		if _, ok := usedWords[word]; !ok {
			result[sortedWord] = append(result[sortedWord], word)
			usedWords[word] = struct{}{}
		}
	}

	return result
}

func deleteSingleWords(anagrams_map mapAnagrams_t) {
	for key, anagrams := range anagrams_map {
		if len(anagrams) < 2 {
			delete(anagrams_map, key)
		}
	}
}

func makeGoodKeys(anagrams_map mapAnagrams_t) {
	for key, anagrams := range anagrams_map {
		delete(anagrams_map, key)
		new_key := anagrams[0]
		anagrams_map[new_key] = anagrams
	}
}

func sortAnagrams(anagrams_map mapAnagrams_t) {
	for k := range anagrams_map {
		sort.Strings(anagrams_map[k])
	}
}


func MakeAnagrams(input []string) mapAnagrams_t {
	result := makeAnagramsSortedLettersKey(input)
	deleteSingleWords(result)
	makeGoodKeys(result)
	sortAnagrams(result)
	return result
}



