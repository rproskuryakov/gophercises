package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func conductQuiz(filename string, timeLimit int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	nRightAnswers := 0
	i := 0
	fmt.Printf("Press enter to start quiz: ")
	timer := time.NewTimer(time.Second * time.Duration(timeLimit))
	go iterateThroughQuiz(*reader, timeLimit, &i, &nRightAnswers)
	<-timer.C
	fmt.Printf("\nYou scored %v out of %v.\n", nRightAnswers, i)
	defer file.Close()
}

func iterateThroughQuiz(r csv.Reader, timeLimit int, nQuestions *int, nRightAnswers *int) {
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		askQuestion(record, nRightAnswers, *nQuestions)
		*nQuestions++
	}
}

func askQuestion(record []string, nRightAnswers *int, nQuestions int) bool {
	answer := ""
	fmt.Printf("Problem #%v: %v = ", nQuestions, record[0])
	fmt.Scanf("%s", &answer)
	if record[1] != answer {
		fmt.Printf("Wrong answer.\n")
		return false
	} else {
		*nRightAnswers++
		return true
	}
}

func main() {
	filename := flag.String("csv", "problems.csv", "Path to csv file with quiz data.")
	timeLimit := flag.Int("limit", 30, "Limit in seconds for to answer each question.")
	flag.Parse()
	fmt.Printf("filename is: %v and time limit is: %v\n", *filename, *timeLimit)
	conductQuiz(*filename, *timeLimit)
}
