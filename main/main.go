package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Notand105/go-todo/todo"
)

const (
	todoFile = "./.todos.json"
)

func main() {
	fmt.Println("hello world")
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "Mark a todo to complete")
	del := flag.Int("delete", 0, "delete a todo ")
	list := flag.Bool("list", false, "show the list of todos")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		handleError(err)
		todos.Add(task)
		err = todos.Store(todoFile)
		handleError(err)
	case *complete > 0:
		err := todos.Complete(*complete)
		handleError(err)
		err = todos.Store(todoFile)
		handleError(err)
	case *del > 0:
		err := todos.Delete(*del)
		handleError(err)
		err = todos.Store(todoFile)
		handleError(err)
	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}

}

func handleError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(1)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {

	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed")
	}

	return text, nil

}
