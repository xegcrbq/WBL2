package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с
поддержкой ряда простейших команд:
- cd <args> - смена директории (в качестве аргумента могут
быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте
аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в
формате *такой-то формат*
Так же требуется поддерживать функционал fork/exec-команд
Дополнительно необходимо поддерживать конвейер на пайпах
(linux pipes, пример cmd1 | cmd2 | .... | cmdN).
*Шелл — это обычная консольная программа, которая будучи
запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись
ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный
сеанс поддерживается до тех пор, пока не будет введена
команда выхода (например \quit).
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		currentDir, _ := os.Getwd()
		fmt.Printf("%v> ", currentDir)
		if scanner.Scan() {
			data := scanner.Text()
			commands := strings.Split(data, "|")

			RunCommands(commands)
		}

	}
}

func RunCommands(commands []string) {
	for _, command := range commands {
		args := strings.Split(command, " ")
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Fprintf(os.Stderr, "path required")
				continue
			}
			os.Chdir(args[1])
		case "echo":
			if len(args) < 2 {
				fmt.Fprintf(os.Stdout, "")
				continue
			}
			for i := 1; i < len(args)-1; i++ {
				fmt.Fprintf(os.Stdout, args[i]+" ")
			}
		case "kill":
			if len(args) < 2 {
				fmt.Fprintf(os.Stdout, "need pid\n")
				continue
			}
			cmd := exec.Command(args[0], args[1])

			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout

			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case `\quit`:
			fmt.Fprintf(os.Stdout, "Bye!")
			os.Exit(1)
		default:
			cmd := exec.Command(command)

			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout

			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}
