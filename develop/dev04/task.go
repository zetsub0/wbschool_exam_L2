package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortRunes(str string) string {
	r := []rune(str)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	seen := make(map[string]string)

	words = removeDuplicates(words)

	for _, word := range words {
		wordLower := strings.ToLower(word)
		sortedWord := sortRunes(wordLower)

		if key, found := seen[sortedWord]; found {
			anagrams[key] = append(anagrams[key], wordLower)
		} else {

			seen[sortedWord] = wordLower
			anagrams[wordLower] = []string{wordLower}
		}
	}

	// Удаление множеств анаграмм, содержащих только одно слово
	for key, group := range anagrams {
		if len(group) < 2 {
			delete(anagrams, key)
		} else {
			sort.Strings(anagrams[key]) // Сортировка множеств анаграмм
		}
	}

	return anagrams
}

func removeDuplicates(slice []string) []string {
	encountered := map[string]struct{}{}
	var result []string

	for _, v := range slice {
		lowerCaseV := strings.ToLower(v)
		if _, ok := encountered[lowerCaseV]; !ok {
			encountered[lowerCaseV] = struct{}{}
			result = append(result, v)
		}
	}

	return result
}

func main() {
	words := []string{"eat", "ate", "tea", "bike", "kibe", "cab", "bca", "abc", "abc", "sdfsfd"}
	anagramMap := findAnagrams(words)
	for _, value := range anagramMap {
		fmt.Printf("\"%s\": %v\n", value[0], value)
	}
}
