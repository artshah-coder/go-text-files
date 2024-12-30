// Given a text file. Remove all empty lines from it.
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
	if status, err := RmEmptyLines(FILENAME); !status {
		fmt.Printf("Error writing file: %s.\n", err)
		return
	}
}

func RmLinesInFile(path, s string) ([]string, error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return RmLinesInReader(f, s)
}

func RmLinesInReader(r io.Reader, s string) ([]string, error) {
	reader := bufio.NewReader(r)
	strs := []string{}
	var str string
	var err error
	for str, err = reader.ReadString('\n'); err == nil; str, err = reader.ReadString('\n') {
		if strings.TrimSpace(str) == strings.TrimSpace(s) {
			continue
		}
		strs = append(strs, str)
	}
	if err != nil && err != io.EOF {
		return nil, err
	}
	return strs, nil
}

func RmEmptyLines(path string) (bool, error) {
	strs, err := RmLinesInFile(path, "\n")
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
