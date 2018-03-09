package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/golangtokyo/gomopost"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 3 {
		return errors.New("not enough arguments: addr name message")
	}
	addr, name, message := args[0], args[1], args[2]

	client := gomopost.NewClient(addr)
	return client.Post(name, message)
}
