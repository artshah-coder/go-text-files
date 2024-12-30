// Given an integer K and a text file. Insert an empty line after the K-th file line.
// If there is no line with this number, then keep the file unchanged
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const FILENAME = "textfile"
const K = 3

func main() {
	if status, err := insertLine2File(FILENAME, "\n", 3); !status {
		fmt.Printf("Error while string inserting to file %s: %s.\n", FILENAME, err)
	}
}

func file2Lines(path string) ([]string, error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return linesFromReader(f)
}

func linesFromReader(r io.Reader) ([]string, error) {
	strs := []string{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		strs = append(strs, scanner.Text()+"\n")
	}
	if err := scanner.Err(); err != nil {
		return nil, scanner.Err()
	}
	return strs, nil
}

func insertLine2File(path, str string, index int) (bool, error) {
	lines, err := file2Lines(path)
	if err != nil {
		return false, err
	}
	fileContent := ""
	for i, line := range lines {
		if i == index {
			fileContent += str
		}
		fileContent += line
	}
	if err = os.WriteFile(path, []byte(fileContent), 0644); err != nil {
		return false, err
	}
	return true, nil
}
