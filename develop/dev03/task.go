package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	reverse     = flag.Bool("r", false, "Sort in reversed order?")
	numerically = flag.Bool("n", false, "Sort numerically")
	unique      = flag.Bool("u", false, "Dont print duplicates")
	column      = flag.Int("k", 1, "Order by specified column")
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(mySort(data, *reverse, *numerically, *unique, *column))
}

func mySort(data []byte, r, n, u bool, k int) (string, error) {
	rows := strings.Split(string(data), "\n")
	var res string

	if u { // if user asked for unique
		rows = removeDuplicateStr(rows)
	}

	if n {
		numbers := make([]int, 0)
		for _, row := range rows {
			if numRow, err := strconv.Atoi(row); err == nil {
				numbers = append(numbers, numRow)
			} else {
				return "", errors.New("not numerical data")
			}
		}
		sort.Ints(numbers)
		if r {
			sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
		}
		for _, row := range numbers {
			res += fmt.Sprintln(row)
		}
		return res, nil
	}

	matrix := make([][]string, 0)
	for _, row := range rows {
		rowSlice := strings.Split(row, " ")
		matrix = append(matrix, rowSlice)
	}

	if k < 0 || k >= len(matrix[0]) {
		return "", errors.New(fmt.Sprintf("Incorrect column number: %d\n", k))
	}

	sort.Slice(matrix, func(i, j int) bool {
		for x := k; x < len(matrix[i]); x++ {
			if matrix[i][k] == matrix[j][k] {
				continue
			}
			if r {
				return matrix[i][k] > matrix[j][k]
			} else {
				return matrix[i][k] < matrix[j][k]
			}
		}
		return true
	})

	for _, rowSlice := range matrix { // Print out
		for i := 0; i < len(rowSlice); i++ {
			res += rowSlice[i]
			if i != len(rowSlice)-1 { // if last word of row
				res += " "
			} else {
				res += "\n"
			}
		}
	}

	return res, nil
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := make([]string, 0)
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
