// Given a text file. Replace all uppercase Russian letters in it with lowercase ones,
// and all lowercase ones with uppercase ones.
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
	if err := changeCase(FILENAME); err != nil {
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

	return reader2Lines(f)
}

func addNewLine(s string) string {
	if !strings.Contains(s, "\n") {
		s += "\n"
	}
	return s
}

func reader2Lines(r io.Reader) ([]string, error) {
	strs := []string{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		strs = append(strs, addNewLine(scanner.Text()))
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return strs, nil
}

func changeCase(path string) error {
	strs, err := file2Lines(path)
	if err != nil {
		return err
	}

	contentFile := ""
	for _, str := range strs {
		for _, ch := range str {
			switch {
			case ch >= 'А' && ch <= 'Я':
				contentFile += string(ch + 32)
			case ch >= 'а' && ch <= 'я':
				contentFile += string(ch - 32)
			case ch == 'Ё':
				contentFile += string('ё')
			case ch == 'ё':
				contentFile += string('Ё')
			default:
				contentFile += string(ch)
			}
		}
		fmt.Println()
	}

	return os.WriteFile(path, []byte(contentFile), 0644)
}
