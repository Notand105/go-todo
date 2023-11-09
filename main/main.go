package main

import (
	"flag"
	"fmt"
)

const (
	todoFile = ".todos.json"
)

func main() {
	fmt.Println("hello world")
	add := flag.Bool("add", false, "add a new todo")
}
