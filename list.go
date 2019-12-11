package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	doneMark1 = "\u2610"
	doneMark2 = "\u2611"
)

func listTasks(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	br := bufio.NewReader(f)
	n := 1
	for {
		b, _, err := br.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		line := string(b)
		if strings.HasPrefix(line, "-") {
			fmt.Printf("\x1b[32m%s\x1b[0m %3d: %s\n", doneMark2, n, strings.TrimSpace(line[1:]))
		} else {
			fmt.Printf("\x1b[31m%s\x1b[0m %3d: %s\n", doneMark1, n, strings.TrimSpace(line))
		}
		n++
	}

	return nil
}
