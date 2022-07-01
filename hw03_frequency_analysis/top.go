package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`[^\wа-яА-Я-]`)

type wordsCount struct {
	Word  string
	Count int64
}

func Top10(str string) []string {
	words := splitString(str)
	groupedWords := groupWords(words)
	sortedWords := sortWords(groupedWords)
	return getTopWords(sortedWords)
}

func splitString(str string) []string {
	strWithoutSymbols := re.ReplaceAllString(str, " ")
	words := strings.Fields(strWithoutSymbols)
	return words
}

func groupWords(words []string) map[string]wordsCount {
	wordsMap := make(map[string]wordsCount)

	for _, word := range words {
		if word == "-" {
			continue
		}

		key := strings.ToLower(word)
		if v, ok := wordsMap[key]; ok {
			v.Count++
			wordsMap[key] = v
		} else {
			wordsMap[key] = wordsCount{
				Word:  key,
				Count: 1,
			}
		}
	}

	return wordsMap
}

func sortWords(wordsMap map[string]wordsCount) []wordsCount {
	wordsSlice := make([]wordsCount, 0, len(wordsMap))

	for _, word := range wordsMap {
		wordsSlice = append(wordsSlice, word)
	}

	sort.SliceStable(wordsSlice, func(i, j int) bool {
		if wordsSlice[i].Count == wordsSlice[j].Count {
			return wordsSlice[i].Word < wordsSlice[j].Word
		}
		return wordsSlice[i].Count > wordsSlice[j].Count
	})

	return wordsSlice
}

func getTopWords(source []wordsCount) []string {
	words := make([]string, 0)

	for i, word := range source {
		words = append(words, word.Word)
		if i == 9 {
			break
		}
	}

	return words
}
