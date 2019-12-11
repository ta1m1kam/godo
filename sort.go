package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func sort_tasks(filename string) error {
	var bottom bytes.Buffer
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
			_, err = fmt.Fprintf(&bottom, "%s\n", line)
			if err != nil {
				return err
			}
		} else {
			_, err = fmt.Fprintf(w, "%s\n", line)
			if err != nil {
				return err
			}
		}
	}
	_, err = bottom.WriteTo(w)
	if err != nil {
		return err
	}

	err = os.Remove(filename)
	if err != nil {
		return err
	}

	os.Rename(filename+"_", filename)
	return nil
}
