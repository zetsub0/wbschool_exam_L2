Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```go
Создается два канала при помощи функции asChan, которая с рандомной задержкой добавляет значения в канал.
Эти оба канала объединяются в канал c в функции merge. Происходит чтение из канала c. 
Однако как только один из каналов будет полностью прочитан начнется спам нулевых значений.
Для избежания этого можно немного доработать select в функции merge:

select {
	case v, ok := <-a:
		if !ok {
			return
		}
		c <- v
	case v, ok := <-b:
		if !ok {
			return
		}
		c <- v
	}

```
