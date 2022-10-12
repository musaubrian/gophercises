package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func cliQuizTimer() {
	csvFileName := flag.String(
		"file",
		"problems.csv",
		"the csv file name")

	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatalf("%s", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("%s", err)
	}

	correctAnswers := 0

	for _, value := range data {
		question := value[0]
		answer := value[1]

		fmt.Printf("%s = ", question)

		var userAnswer string
		fmt.Scanln(&userAnswer)

		if userAnswer == answer {
			correctAnswers++
		}
	}
	fmt.Printf("You answered %d/%d correctly\n", correctAnswers, len(data))
}
