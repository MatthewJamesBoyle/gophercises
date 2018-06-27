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
	csvFileName := flag.String("csv", "problems.csv", "csv file with question,answer")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz, in seconds.")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Failed to open the CSV file %s \n", *csvFileName)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, error := r.ReadAll()

	if error != nil {
		fmt.Printf("Couldn't parse the file: %s \n", *csvFileName)
		os.Exit(1)
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems {
		fmt.Printf("problem #%d: %s =\n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			trimmed := strings.ToLower(strings.TrimSpace(answer))
			answerCh <- trimmed
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <- answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("Congratulations, you got through every question in the quiz! You scored %d out of %d.\n", correct, len(problems))

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
