package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	
	file, err := os.Open(*csvFileName)
	if err != nil {
		panic("Failed to open the CSV file: " + *csvFileName)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		panic("Failed to parse the provided CSV file.")
	}

	quiz := make([]problem, len(lines))

	for i, line := range lines {
		quiz[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	var correct = 0

	for i, problem := range quiz {
		var answer string
		fmt.Printf("Question %d: %s = \n", i+1, problem.question)
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(quiz))
}