package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func renameTask(filename, taskName string, id int) error {
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
		if id == n {
			match = true
		}

		line := strings.TrimSpace(string(b))
		if match {
			if strings.HasPrefix(line, "-") {
				fmt.Fprintf(w, "-%s\n", taskName)
			} else {
				fmt.Fprintf(w, "%s\n", taskName)
			}
		} else {
			fmt.Fprintf(w, "%s\n", line)
		}
	}

	err = os.Remove(filename)
	if err != nil {
		return err
	}

	os.Rename(filename+"_", filename)
	return nil
}
