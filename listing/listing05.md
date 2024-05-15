Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error

customError имплементирует интерфейс error, поэтому на этапе err = test() не возникло ошибок.
при проверке err != nil мы провалились в println, так как интерфейс err имеет лишь значение nil, но не тип.

```
