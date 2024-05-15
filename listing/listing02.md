Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

В функции test() используется именованный возвращаемый параметр, defer может изменить его.

Функция anotherTest() не использует именованных возвращаемых параметров. Возвращаемое значение уже предопределено при вызове return, и defer не может на него влиять.

```
