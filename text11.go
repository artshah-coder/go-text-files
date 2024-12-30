// Given a text file. Duplicate all empty lines in it.
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
	if status, err := doubleNewLines(FILENAME); !status {
		fmt.Printf("Error writing file: %s.\n", err)
		return
	}
}

func DoubleLinesInFile(path, s string) ([]string, error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return DoubleLinesInReader(f, s)
}

func DoubleString(s string) string {
	return s + s
}

func DoubleLinesInReader(r io.Reader, s string) ([]string, error) {
	reader := bufio.NewReader(r)
	strs := []string{}
	var str string
	var err error
	for str, err = reader.ReadString('\n'); err == nil; str, err = reader.ReadString('\n') {
		if strings.TrimSpace(str) == strings.TrimSpace(s) {
			strs = append(strs, DoubleString(s))
			continue
		}
		strs = append(strs, str)
	}
	if err != nil && err != io.EOF {
		return nil, err
	}
	return strs, nil
}

func doubleNewLines(path string) (bool, error) {
	strs, err := DoubleLinesInFile(path, "\n")
	if err != nil {
		return false, err
	}

	fileContent := ""
	for _, s := range strs {
		fileContent += s
	}
	if err = os.WriteFile(path, []byte(fileContent), 0644); err != nil {
		return false, err
	}
	return true, nil
}
