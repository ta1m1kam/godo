package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func undone(filename string, args []string) error {
	if len(args) == 0 {
		return nil
	}

	var ids []int
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			return err
		}
		ids = append(ids, id)
	}

	w, err := os.Create(filename + "_")
	if err != nil {
		return err
	}
	defer w.Close()

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	br := bufio.NewReader(f)
	for n := 1; ; n++ {
		b, _, err := br.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		match := false
		for _, id := range ids {
			if id == n {
				match = true
			}
		}
		line := strings.TrimSpace(string(b))
		if match && strings.HasPrefix(line, "-") {
			_, err = fmt.Fprintf(w, "%s\n", line[1:])
			if err != nil {
				return err
			}
			fmt.Printf("Task undone: %s\n", line[1:])
		} else {
			_, err = fmt.Fprintf(w, "%s\n", line)
			if err != nil {
				return err
			}
		}
	}
	err = os.Remove(filename)
	if err != nil {
		return err
	}

	os.Rename(filename+"_", filename)
	return nil
}
