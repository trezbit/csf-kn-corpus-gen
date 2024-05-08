package utils

import (
	"bufio"
	"fmt"
	"os"
)

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	CheckError(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
