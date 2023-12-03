package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileToLines(filePath string) ([]string, error) {
	wdir, err := os.Getwd()
	check(err)

	file, err := os.Open(path.Join(wdir, filePath))
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())

	return lines, nil
}

type BallSet struct {
	red   int
	blue  int
	green int
}

func parseRound(str string) BallSet {
	shows := strings.Split(str, ", ")

	round := BallSet{0, 0, 0}

	for _, show := range shows {
		show = strings.TrimSpace(show)

		parts := strings.Split(show, " ")

		count, err := strconv.Atoi(parts[0])
		check(err)

		switch parts[1] {
		case "red":
			round.red = count
		case "green":
			round.green = count
		case "blue":
			round.blue = count
		default:
			panic("error in round parsing")
		}
	}

	return round
}

var targetRed int = 12
var targetGreen int = 13
var targetBlue int = 14

// gives game's id , its validity, minimum possible BallSet for the entire game
func parseGame(line string) (int, bool, BallSet) {
	parts := strings.Split(line, ": ")
	gameInfo, game := parts[0], parts[1]
	id, err := strconv.Atoi(strings.Split(gameInfo, " ")[1])
	check(err)

	//convert game to rounds
	rounds := strings.Split(game, "; ")

	isValid := true

	minBalls := BallSet{0, 0, 0}

	for _, roundStr := range rounds {
		round := parseRound(roundStr)

		if round.red > targetRed || round.blue > targetBlue || round.green > targetGreen {
			isValid = false
		}

		//finding min balls needed
		minBalls.red = max(minBalls.red, round.red)
		minBalls.green = max(minBalls.green, round.green)
		minBalls.blue = max(minBalls.blue, round.blue)
	}

	return id, isValid, minBalls
}

func Solve() {
	//get file to string arr
	//parse game to return the game id and valid are not
	//add the passed game numbers

	lines, err := fileToLines("02cubeCon/data.txt")
	check(err)

	sum := 0

	for _, line := range lines {
		id, isValid, _ := parseGame(line)

		fmt.Printf("id is %v, its validity %v \n", id, isValid)

		if isValid {
			sum += id
		}
	}

	fmt.Println("p1 answer is: ", sum)
}

func Solve2() {
	lines, err := fileToLines("02cubeCon/data.txt")
	check(err)

	mulSum := 0

	for _, line := range lines {
		_, _, minBallsNeeded := parseGame(line)

		mulSum += (minBallsNeeded.red * minBallsNeeded.green * minBallsNeeded.blue)

	}

	fmt.Println(mulSum)
}

func main() {
	// Solve()
	Solve2()
}
