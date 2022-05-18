package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type command struct {
	name string
	args []string
}

// мапа соответствий названия действия функции-обработчику.
var cmdMap = map[string]func(cmd command) (string, error){
	"cd":   cd,
	"pwd":  pwd,
	"echo": echo,
	"kill": kill,
	"ps":   ps,
	"exit": exit,
}

func cd(cmd command) (string, error) {
	err := os.Chdir(cmd.args[0])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Changed dir to %s", cmd), nil
}

func pwd(cmd command) (string, error) {
	cmdCompl := exec.Command(cmd.name)
	stdout, err := cmdCompl.Output()
	if err != nil {
		return "", err
	}

	return string(stdout), nil
}

func echo(cmd command) (string, error) {
	cmdCompl := exec.Command(cmd.name, cmd.args...)
	stdout, err := cmdCompl.Output()
	if err != nil {
		return "", err
	}

	return string(stdout), nil
}

func kill(cmd command) (string, error) {
	var (
		pid  int
		proc *os.Process
		err  error
	)
	if len(cmd.args) < 1 {
		log.Println("not enough arguments")
	}
	for _, value := range cmd.args {

		pid, err = strconv.Atoi(value)
		if err != nil {
			return "", fmt.Errorf("incorrect process pid")
		}
		proc, err = os.FindProcess(pid)
		if err != nil {
			return "", fmt.Errorf("process id not found")
		}
		err = proc.Kill()
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("process id %d killed", pid), nil
}

func ps(cmd command) (string, error) {
	cmdCompl := exec.Command(cmd.name, cmd.args...)
	stdout, err := cmdCompl.Output()
	if err != nil {
		return "", err
	}
	return string(stdout), nil
}

func exit(cmd command) (string, error) {
	fmt.Printf("Enter command %s. Shell completion!\n", cmd.name)
	os.Exit(0)
	return "", nil
}

func processCmd(cmd command) (string, error) {
	if cmdFunc, ok := cmdMap[cmd.name]; ok {
		return cmdFunc(cmd)
	}
	return "", fmt.Errorf("command not implemented")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$ ")
		scanner.Scan()
		input := scanner.Text()
		tokens := []string{}
		// strings.Split разбивает строку по заданному символу
		for _, token := range strings.Split(input, " ") {
			token = strings.TrimSpace(token) // отрезаем пробелы
			if token != "" {
				tokens = append(tokens, token)
			}
		}

		cmd := &command{}
		cmd.name = tokens[0]
		cmd.args = tokens[1:]

		res, err := processCmd(*cmd)
		if err != nil {
			fmt.Printf("%s: %v\n", input, err)
			continue
		}
		fmt.Println(res)
	}

}
