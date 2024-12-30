// Given a string S and a text file.
// Replace all empty lines in the file with the string S.
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
	replace := "String for replace"
	if err := replaceEmptyLines(FILENAME, replace); err != nil {
		fmt.Println(err)
		return
	}
}

func replaceLinesInFile(path, src, dst string) ([]string, error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return replaceLinesInStream(f, src, dst)
}

func addNewLine(s string) string {
	if strings.Contains(s, "\n") {
		return s
	} else {
		return s + "\n"
	}
}

func replaceLinesInStream(r io.Reader, src, dst string) ([]string, error) {
	scanner := bufio.NewScanner(r)
	strs := []string{}
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == strings.TrimSpace(src) {
			strs = append(strs, addNewLine(dst))
			continue
		}
		strs = append(strs, addNewLine(scanner.Text()))
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return strs, nil
}

func replaceEmptyLines(path, dst string) error {
	strs, err := replaceLinesInFile(path, "\n", dst)
	if err != nil {
		return err
	}
	fileContent := ""
	for _, s := range strs {
		fileContent += s
	}

	if err = os.WriteFile(path, []byte(fileContent), 0644); err != nil {
		return err
	}
	return nil
}
