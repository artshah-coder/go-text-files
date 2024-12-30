// Given a non-empty text file. Remove the first line from it.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const FILENAME = "textfile"

func main() {
	if err := removeTheFirstLine(FILENAME); err != nil {
		fmt.Println(err)
		return
	}
}

func file2Lines(path string) ([]string, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return stream2Lines(f)
}

func addNewLine(s string) string {
	if !strings.Contains(s, "\n") {
		s += "\n"
	}
	return s
}

func stream2Lines(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	strs := []string{}
	for scanner.Scan() {
		strs = append(strs, addNewLine(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return strs, nil
}

func removeTheFirstLine(path string) error {
	strs, err := file2Lines(path)
	if err != nil {
		return err
	}
	fileContent := ""
	for i, s := range strs {
		if i == 0 {
			continue
		}
		fileContent += s
	}

	if err = os.WriteFile(path, []byte(fileContent), 0644); err != nil {
		return err
	}
	return nil
}
