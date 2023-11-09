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
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "Mark a todo to complete")
	del := flag.Int("delete", 0, "delete a todo ")
	list := flag.Bool("la", false, "show the list of todos")
	due := flag.Bool("ld", false, "show the list of todos that are not completed")
  purge := flag.Bool("purge", false, "delete every todo")

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
		todos.Print(true)
	case *complete > 0:
		err := todos.Complete(*complete)
		handleError(err)
		err = todos.Store(todoFile)
		handleError(err)
		todos.Print(true)
	case *del > 0:
		err := todos.Delete(*del)
		handleError(err)
		err = todos.Store(todoFile)
		handleError(err)
		todos.Print(true)
	case *list:
		todos.Print(true)
	case *due:
		todos.Print(false)
  case *purge:
    err := todos.Purge()
		handleError(err)
		err = todos.Store(todoFile)
		handleError(err)
		todos.Print(true)
	default:
		//fmt.Fprintln(os.Stdout, "invalid command")
		//os.Exit(1)
    fmt.Println("-add <text>  -> to add a new task")
    fmt.Println("-complete <index>  -> to complete a certain task")
    fmt.Println("-delete <index>  -> to delete a certain task")
    fmt.Println("-la  -> show every todo")
    fmt.Println("-la  -> show no completed todos")
    fmt.Println("-purge  -> delete every todo")

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
