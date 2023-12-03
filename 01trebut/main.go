package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

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

func p1() {
	// open the doc
	// go through each line
	//// in each line find the end num chars and join them
	// convert those num strings to numbers
	//add them

	lines, err := utils.FileToLines("01trebut/data.txt")
	utils.Check(err)
	
	res := 0

	for _, line := range lines {
		//separating the numChars
		numStr := getFirstNumChar(line)
		numStr += getFirstNumChar(reverseString(line))
	
		//string to num
		num, err := strconv.Atoi(numStr)
		utils.Check(err)
		res += num
	}

	fmt.Println("ans is:", res)
}

func p2() {
	//get each line
	//do a front scan
	//do a back scan
	//join the digits
	// add the nums
}

func main() {
	p1()
	p2()
}
