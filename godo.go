package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

const (
	todoFilename = ".todo"
)

func getStorageFile() string {
	filename := ""
	existCurTodo := false
	curDir, err := os.Getwd()
	if err == nil {
		filename = filepath.Join(curDir, todoFilename)
		_, err = os.Stat(filename)
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
				Action: func(c *cli.Context) error {
					listTasks(filename)
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					addTask(filename, c.Args().Get(0))
					listTasks(filename)
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"r"},
				Usage:   "delete a task from the list",
				Action: func(c *cli.Context) error {
					deleteTask(filename, c.Args().Slice())
					listTasks(filename)
					return nil
				},
			},
			{
				Name:    "done",
				Aliases: []string{"d"},
				Usage:   "done a task",
				Action: func(c *cli.Context) error {
					doneTask(filename, c.Args().Slice())
					listTasks(filename)
					return nil
				},
			},
			{
				Name:    "undone",
				Aliases: []string{"u"},
				Usage:   "undone a task",
				Action: func(c *cli.Context) error {
					undoneTask(filename, c.Args().Slice())
					listTasks(filename)
					return nil
				},
			},
			{
				Name:    "clean",
				Aliases: []string{"c"},
				Usage:   "clean done tasks",
				Action: func(c *cli.Context) error {
					cleanDoneTask(filename)
					listTasks(filename)
					return nil
				},
			},
			{
				Name:    "sort",
				Aliases: []string{"s"},
				Usage:   "sort tasks",
				Action: func(c *cli.Context) error {
					sortTasks(filename)
					listTasks(filename)
					return nil
				},
			},
			{
				Name:    "rename",
				Aliases: []string{"rn"},
				Usage:   "rename task on the list",
				Action: func(c *cli.Context) error {
					id, err := strconv.Atoi(c.Args().Get(0))
					if err != nil {
						return err
					}
					renameTask(filename, c.Args().Get(1), id)
					listTasks(filename)
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
