package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	readFile()
}

func readFile() {

	csvFilename := flag.String("csv", "problems.csv", "a csv file on the of 'question,answer'")
	timeLimit := flag.Int("limit", 30.,
		"the time limit fir the quiz in secounds")

	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {

		fmt.Println("failed to openm the CSV file: %s", *csvFilename)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {

		exit("failed to parse csv file")

	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("you scored %d . out of %d \n ", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}

		}
	}
	fmt.Printf("you scored %d . out of %d \n ", correct, len(problems))
}
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret

}

type problem struct {
	q string
	a string
}

func exit(msg string) {

	fmt.Println(msg)
	os.Exit(1)

}
