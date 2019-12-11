package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"github.com/urfave/cli/v2"
	"sort"
)

const (
	todoFilename = ".todo"
)

func getStorageFile() string {
	filename := ""
	existCurTodo := false
	curDir, err := os.Getwd()
	if err ==  nil {
		filename =  filepath.Join(curDir, todoFilename)
		_, err =os.Stat(filename)
		if err == nil {
			existCurTodo = true
		}
	}

	if !existCurTodo {
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		filename = filepath.Join(home, todoFilename)
	}

	return filename
}

func main() {
	filename := getStorageFile()

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "task on the list",
				Action:  func(c *cli.Context) error {
					fmt.Println(filename)
					list(filename)
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action:  func(c *cli.Context) error {
					add(filename, c.Args().Get(0))
					list(filename)
					return nil
				},
			},
			{
				Name: "delete",
				Aliases: []string{"r"},
				Usage: "delete a task",
				Action: func(c *cli.Context) error {
					delete(filename, c.Args().Slice())
					list(filename)
					return nil
				},
			},
			{
				Name: "done",
				Aliases: []string{"d"},
				Usage: "done a task",
				Action: func(c *cli.Context) error {
					done(filename, c.Args().Slice())
					list(filename)
					return nil
				},
			},
			{
				Name: "undone",
				Aliases: []string{"u"},
				Usage: "undone a task",
				Action: func(c *cli.Context) error {
					undone(filename, c.Args().Slice())
					list(filename)
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
