package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var x string
	fmt.Scanln(&x)
	fmt.Println(Unpack(x))
}

func Unpack(str string) (string, error) {

	if len(str) == 0 {
		return "", nil
	}

	if unicode.IsDigit(rune(str[0])) {
		return "", errors.New("incorrect string")
	}

	res := make([]rune, 0, len(str))
	last := ' ' // buffer

	rStr := []rune(str)

	for i := 0; i < len(rStr); i++ {
		if string(rStr[i]) == `\` { // escaping backslash
			if i+1 < len(rStr) {
				i++
				last = rStr[i]
				res = append(res, last)
				continue
			}
		}
		if unicode.IsDigit(rStr[i]) {
			num := 0
			for i < len(rStr) && unicode.IsDigit(rStr[i]) {
				digit, _ := strconv.Atoi(string(rStr[i]))
				num = num*10 + digit // if digit > 9
				i++
			}
			i--
			res = append(res, Repeat(last, num-1)...)
		} else {
			last = rStr[i]
			res = append(res, last)
		}
	}
	return string(res), nil
}

func Repeat(char rune, n int) []rune {
	res := make([]rune, 0)
	for i := 0; i < n; i++ {
		res = append(res, char)
	}
	return res
}
