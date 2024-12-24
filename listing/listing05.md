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
Программа выведет: `error`.

Произойдет это потому, что переменная `err` является интерфейсом типа `error`. Когда функция `test()` возвращает `nil`, это значение 
типа `*customError`, который не является `nil` как тип. Интерфейс считается ненулевым, если он хранит тип, отличный от нуля, даже если 
значение внутри интерфейса равно nil.