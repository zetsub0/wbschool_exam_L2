Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
[3, 2, 3]
У слайса cap = 3. Измение первого элемента сработало, так как внутри функции был указатель на тот же слайс. Как только сработал append, капасити слайса увеличился вдвое и слайс внутри функции перестал указывать область слайса s.
```