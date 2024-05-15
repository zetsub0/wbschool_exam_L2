package main

import (
	"bufio"
	"fmt"
	ps2 "github.com/mitchellh/go-ps"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

func main() {
	wd, _ := os.Getwd()
	fmt.Print(wd + "> ")
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); fmt.Print(wd + "> ") {
		if query := scanner.Text(); query != "\\quit" {
			execute(query)
		} else {
			break
		}
		wd, _ = os.Getwd()
	}
}

func cd(dir string) error {
	return os.Chdir(dir)
}

func pwd() (string, error) {
	if workDir, err := os.Getwd(); err != nil {
		return "", err
	} else {
		return workDir, nil
	}
}

func echo(args ...string) {
	fmt.Println(strings.Join(args, " "))
}

func kill(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return process.Kill()
}

func ps() {
	pSlice, err := ps2.Processes()
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range pSlice {
		fmt.Println(p.PPid(), p.Pid(), p.Executable())
	}
}

func execute(query string) {
	commands := strings.Split(query, " | ")
	for _, command := range commands {
		commandSlice := strings.Split(command, " ")
		switch commandSlice[0] {

		case "pwd":
			if wd, err := pwd(); err != nil {
				log.Println(err)
				return
			} else {
				fmt.Println(wd)
			}

		case "cd":
			err := cd(commandSlice[1])
			if err != nil {
				log.Println(err)
				return
			}

		case "echo":
			echo(commandSlice[1:]...)

		case "ps":
			ps()

		case "kill":
			if pid, err := strconv.Atoi(commandSlice[1]); err != nil {
				fmt.Println(err.Error())
				return
			} else {
				kill(pid)
			}
		}
	}

}
