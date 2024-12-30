// Given an integer K and a text file. Remove the first K characters
// from each line of the file (if the length of the line is less than K,
// then remove all characters from it).
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

const FILENAME = "textfile"
const K = 5

func main() {
	if err := delCharFromLines(FILENAME, K); err != nil {
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

func reader2Lines(r io.Reader) ([]string, error) {
	strs := []string{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return strs, nil
}

func addNewLine(s string) string {
	if !strings.Contains(s, "\n") {
		s += "\n"
	}
	return s
}

func delCharFromLines(path string, num int) error {
	strs, err := file2Lines(path)
	if err != nil {
		return err
	}

	contentFile := ""
	for _, str := range strs {
		if utf8.RuneCountInString(str) <= num {
			str = "\n"
			contentFile += str
			continue
		}
		i := 0
		for j, _ := range str {
			i++
			if i == num+1 {
				str = addNewLine(str[j:])
				break
			}
		}
		contentFile += str
	}

	return os.WriteFile(path, []byte(contentFile), 0644)
}
