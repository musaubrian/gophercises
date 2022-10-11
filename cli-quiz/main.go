package main

import "flag"

func cliQuiz() {
	csvFileName := flag.String("filename", "problems.csv", "the csv file name")
	flag.Parse()

	println(*csvFileName)
}

func main() {
	cliQuiz()
}
