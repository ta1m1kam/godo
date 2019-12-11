package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func clean(filename string) error {
	w, err := os.Create(filename + "_")
	if err != nil {
		return err
	}
	defer w.Close()

	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	br := bufio.NewReader(f)
	for n := 1; ; n++ {
		b, _, err := br.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		line := string(b)
		if !strings.HasPrefix(line, "-") {
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
