package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	// Парсинг аргументов командной строки
	timeoutFlag := flag.Duration("timeout", 10*time.Second, "таймаут подключения")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Необходимо указать хост и порт")
		os.Exit(1)
	}

	host := flag.Args()[0]
	port := flag.Args()[1]

	dialer := net.Dialer{Timeout: *timeoutFlag}

	conn, err := dialer.Dial("tcp", host+":"+port)
	if err != nil {
		fmt.Printf("Ошибка подключения к серверу: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Запуск горутины для чтения данных из соединения
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Соединение с сервером закрыто")
				os.Exit(0)
			}
			fmt.Print(string(buf[:n]))
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		fmt.Println("Программа завершена")
		os.Exit(0)
	}()

	fmt.Println("Соединение установлено. Введите данные:")
	for {
		var input string
		fmt.Scanln(&input)
		conn.Write([]byte(input + "\n"))
	}
}
