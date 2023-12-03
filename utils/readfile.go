package utils

import (
	"bufio"
	"os"
	"path"
)

func FileToLines(filePath string) ([]string, error) {
	wdir, err := os.Getwd()
	Check(err)

	file, err := os.Open(path.Join(wdir, filePath))
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	Check(scanner.Err())

	return lines, nil
}
