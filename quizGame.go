package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type question struct{
	ques string
	ans string
}

func generateQuestions(fileName string) []question{

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error :- The CSV file could not be read.");
		panic(err)
	}
	defer file.Close();

	quesAns, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Error :- The questions from the CSV file could not be read.");
		panic(err)
	}

	var quiz []question


	for _, line := range quesAns {
		qa := question{
			ques: line[0],
			ans: line[1],
		}
		quiz = append(quiz, qa)
	}
	return quiz
}



func main(){

	fmt.Println("Welcome to the CSV Quiz Game designed in Go Lang");
	fileName := flag.String("questions","problems.csv","The CSV file from where questions would be read.")
	flag.Parse()

	var quiz []question
	var score int
	score = 0
	quiz = generateQuestions(*fileName)
	var answer string

	for i, questions := range quiz {
		fmt.Printf("Question %d :- %s? ",i+1,questions.ques)
		fmt.Scan(&answer)
		if answer == questions.ans{
			score++;
		}
	}

	fmt.Printf("You answered %d questions correctly out of a total of %d questions.",score, len(quiz))

}