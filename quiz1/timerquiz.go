package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

type Quiz struct {
	problems []Problem
	score    int
}

func main() {
	quiz := Quiz{}

	// Open the file containing quiz questions
	file, err := os.Open("quiz1.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Scan and read the quiz file
	fin := bufio.NewScanner(file)
	for fin.Scan() {
		row := strings.Split(fin.Text(), ",")
		problem := Problem{question: row[0], answer: row[1]}
		quiz.problems = append(quiz.problems, problem)
	}

	fmt.Printf("Your timer for Quiz starts now\n")
	fmt.Printf("-------------------------------\n")
	fmt.Printf("-------------------------------\n")

	c := 0

	// Start Timer
	timer := time.NewTimer(time.Second * time.Duration(30))
	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Printf("TIMES UP!!!!!!!!!!\nYou answered %d questions correctly out of %d questions.\n", quiz.score, c)
	}()

	for i, problem := range quiz.problems {
		fmt.Printf("Question No.%d:: %s\n", i+1, problem.question)
		c++
		var ans string
		fmt.Scan(&ans)

		if strings.EqualFold(problem.answer, ans) {
			quiz.score++
		}
	}
	fmt.Printf("You answered %d questions correctly out of attempted %d questions.\n", quiz.score, c)
}
