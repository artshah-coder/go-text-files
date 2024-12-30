// Two text files are given. Add to the end of each line of the first file
// the corresponding line of the second file. If the second file is shorter
// than the first, then do not change the remaining lines of the first file.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const FILENAME1 = "textfile1"
const FILENAME2 = "textfile2"

func main() {
	err := mixFiles(FILENAME1, FILENAME2)
	if err != nil {
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
	scanner := bufio.NewScanner(r)
	strs := []string{}
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return strs, nil
}

func mixFiles(path1, path2 string) error {
	f1Lines, err := file2Lines(path1)
	if err != nil {
		return err
	}
	f2Lines, err := file2Lines(path2)
	if err != nil {
		return err
	}

	contentFile := ""
	for i := 0; i < len(f1Lines); i++ {
		if i < len(f2Lines) {
			contentFile += addNewLine(f1Lines[i] + f2Lines[i])
			continue
		}
		contentFile += addNewLine(f1Lines[i])
	}

	return os.WriteFile(path1, []byte(contentFile), 0644)
}
