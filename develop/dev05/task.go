package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	afterFlag   = flag.Int("A", 0, "Print NUM lines of trailing context after matching lines.")
	beforeFlag  = flag.Int("B", 0, "Print NUM lines of leading context before matching lines.")
	contextFlag = flag.Int("C", 0, "Print NUM lines of output context.")
	countFlag   = flag.Bool("c", false, "Suppress normal output; instead print a count of matching lines for each input file. With the -v option count non-matching lines.")
	ignoreFlag  = flag.Bool("i", false, "Ignore case distinctions in patterns and input data, so that characters that differ only in case match each other.")
	invertFlag  = flag.Bool("v", false, "Invert the sense of matching, to select non-matching lines.")
	fixedFlag   = flag.Bool("F", false, "Interpret PATTERNS as fixed strings, not regular expressions.")
	lineNumFlag = flag.Bool("n", false, "Prefix each line of output with the 1-based line number within its input file.")
)

func main() {
	flag.Parse()
	*countFlag = true
	*invertFlag = true
	*ignoreFlag = true
	*fixedFlag = false
	*lineNumFlag = true
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Ошибка при парсинге файла:", err)
		return
	}

	fmt.Println(grep(string(data), "Wor.", *afterFlag, *beforeFlag, *contextFlag, *countFlag, *ignoreFlag, *invertFlag, *fixedFlag, *lineNumFlag))

}

func grep(data string, pattern string, A, B, C int, c, i, v, F, n bool) (string, error) {

	matchedLines := make(map[int]bool)
	counter := 0

	lines := strings.Split(strings.ReplaceAll(data, "\r\n", "\n"), "\n")

	if i {
		pattern = strings.ToLower(pattern)
	}

	for j, line := range lines {

		var (
			matched bool
			err     error
		)

		if i {
			line = strings.ToLower(line)
		}

		if !F {
			matched, err = regexp.MatchString(pattern, line)
			if err != nil {
				return "", err
			}
		} else { // if fixed search
			matched = strings.Contains(line, pattern)
		}

		if matched != v {
			counter++
			if A > 0 {
				C = 0
				for k := j; k < min(len(lines), j+A+1); k++ {
					matchedLines[k] = true
				}
			}

			if B > 0 {
				C = 0
				for k := max(j-B, 0); k <= j; k++ {
					matchedLines[k] = true
				}
			}

			if C > 0 {
				for k := max(j-C, 0); k < min(len(lines), j+C+1); k++ {
					matchedLines[k] = true
				}
			}

			if A == 0 && B == 0 && C == 0 {
				matchedLines[j] = true
			}
		}
	}

	if counter == 0 {
		return "no match", nil
	}

	res := ""

	if c {
		res += fmt.Sprintf("%d", counter)
		return res, nil
	}

	for index, line := range lines {
		if matchedLines[index] {
			newLine := ""
			if n {
				newLine += strconv.Itoa(index) + " "
			}
			newLine += line
			res += newLine + "\n"
		}
	}

	return res, nil
}
