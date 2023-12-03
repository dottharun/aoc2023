package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFirstNumChar(line string) string {
	for _, ch := range line {
		if ch >= '0' && ch <= '9' {
			return string(ch)
		}
	}
	return ""
}

func reverseString(s string) string {
	runes := []rune(s)
	for lo, hi := 0, len(runes)-1; lo < hi; lo, hi = lo+1, hi-1 {
		runes[lo], runes[hi] = runes[hi], runes[lo]
	}
	return string(runes)
}

func main() {
	// open the doc
	// go through each line
	//// in each line find the end num chars and join them
	// convert those num strings to numbers
	//add them

	// ALWAYS RUN THIS PROGRAM FROM PROJECT ROOT
	wdir, err := os.Getwd()
	check(err)

	file, err := os.Open(path.Join(wdir, "01trebut", "data.txt"))
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	res := 0

	for scanner.Scan() {
		//getting the current line
		line := scanner.Text()

		//separating the numChars
		numStr := getFirstNumChar(line)
		numStr += getFirstNumChar(reverseString(line))

		//string to num
		num, err := strconv.Atoi(numStr)
		check(err)
		res += num
	}
	check(scanner.Err())

	fmt.Println("ans is:", res)
}
