package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var fieldsFlag = *flag.String("f", "1", "which columns to print")
var delimiter = *flag.String("d", " ", "delimiter for splitting lines")
var separateFlag = *flag.Bool("s", false, "show only lines with delimiter")

func main() {
	data, _ := os.ReadFile("x.txt")
	fields := "1, 2,3"
	fmt.Println(cut(string(data), fields, delimiter, separateFlag))
}

func parseFields(fields string) ([]int, error) {
	stringIndexes := strings.Split(strings.ReplaceAll(fields, ", ", ","), ",")

	res := make([]int, 0, len(stringIndexes))

	for _, indexStr := range stringIndexes {
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return res, err
		}

		res = append(res, index-1)
	}

	return res, nil
}

func cut(data string, fields string, delimiter string, s bool) (string, error) {
	lines := strings.Split(data, "\n")
	var result []string

	for _, line := range lines {
		if s && !strings.Contains(line, delimiter) {
			continue
		}

		parts := strings.Split(line, delimiter)
		selectedFields, err := parseFields(fields)
		if err != nil {
			return "", err
		}

		var selectedParts []string
		for _, field := range selectedFields {
			if field < len(parts) {
				selectedParts = append(selectedParts, parts[field])
			}
		}

		result = append(result, strings.Join(selectedParts, delimiter))
	}

	return strings.Join(result, "\n"), nil
}
