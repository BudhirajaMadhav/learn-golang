package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	timeLimit := flag.Int64("time", 30, "time to solve the quiz")

	flag.Parse()

	// Open and read the file into lines
	file, err := os.Open(*csvFileName)
	errFound(err)
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	errFound(err)

	// parse the problem set into slice of problem struct 
	problems := parseLinesToProblems(lines)

	// shuffle the problem slice
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(problems), func(i, j int) {problems[i], problems[j] = problems[j], problems[i]})

	// input reader
	input := bufio.NewReader(os.Stdin)

	score := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) //time.second is 1000000000 nanoseconds, duration gives x(parameter) nanoseconds

	for idx, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", idx+1, problem.q)

		answerCh := make(chan string)

		go func() {
			answer, _ := input.ReadString('\n')
			answer = strings.TrimSpace(answer)
			answerCh <- answer
		}()

		// wait for anyone to finish.
		select {
		case <-timer.C:
			fmt.Println("\nYour total score is", score, "out of", len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.a {
				score++
			}
		}

	}

	fmt.Println("Your total score is", score, "out of", len(problems))

}

func parseLinesToProblems(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for idx, line := range lines {
		problems[idx].q = line[0]
		problems[idx].a = line[1]
	}

	return problems

}

func errFound(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
