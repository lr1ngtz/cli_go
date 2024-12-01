package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"todo"
)

var todoFileName = ".todo.json"

func getTask(r io.Reader, args ...string) ([]string, error) {
	if len(args) > 0 {
		return []string{strings.Join(args, " ")}, nil
	}

	s := bufio.NewScanner(r)
	var tasks []string

	for s.Scan() {
		line := s.Text()
		if len(line) > 0 {
			tasks = append(tasks, line)
		}
	}

	if err := s.Err(); err != nil && err != io.EOF {
		return nil, err
	}

	if len(tasks) == 0 {
		return nil, fmt.Errorf("task cannot be blank")
	}

	return tasks, nil
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s ToDo CLI app\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "flag.Usage prints a usage message")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}
	add := flag.Bool("add", false, "Add task to the ToDo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	delete := flag.Int("del", 0, "Item to be deleted")
	verboseOutput := flag.Bool("verb", false, "List all tasks verbose output")
	listNotCompleted := flag.Bool("listnc", false, "List not completed tasks")
	flag.Parse()

	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}
	l := &todo.List{}
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Print(l)
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, task := range t {
			l.Add(task)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *delete > 0:
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *verboseOutput:
		for _, task := range *l {
			fmt.Fprintf(
				os.Stderr,
				"Task: %s\nCreated at: %s\nCompleted at: %s\n\n\n",
				task.Task, task.CreatedAt.String(), task.CompletedAt.String(),
			)
		}
	case *listNotCompleted:
		for _, task := range *l {
			if !task.Done {
				fmt.Println(task)
			}
		}
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}
