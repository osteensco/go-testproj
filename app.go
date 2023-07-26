package main

import (
	"fmt"
	"reflect"
)

func intro() {
	fmt.Println("Hey man, here's a quiz and stuff")
}

var name string

func inputName() {
	fmt.Println("Input your name:")
	fmt.Scan(&name)
}

func inputAnswer(answer string) string {
	fmt.Scan(&answer)
	return answer
}

type QuizProblem struct {
	Question string
	Input    string
	Answer   string
}

func makeQuestion(q string, i string, a string) QuizProblem {
	return QuizProblem{
		Question: q,
		Input:    i,
		Answer:   a,
	}
}

type Quiz struct {
	Q1 QuizProblem
	Q2 QuizProblem
	Q3 QuizProblem
}

func runQuiz(q Quiz) {
	score := make(map[int]int)

	quiztype := reflect.TypeOf(q)

	quizstruct := reflect.ValueOf(q)

	for i := 0; i < quiztype.NumField(); i++ {

		quizProblem := quizstruct.Field(i).Interface().(QuizProblem)

		fmt.Println(quizProblem.Question)

		quizProblem.Input = inputAnswer(quizProblem.Input)

		var point int

		if quizProblem.Input == quizProblem.Answer {
			point = 1
		} else {
			point = 0
		}

		score[i] = point
	}

	var totalscore float64
	for _, v := range score {
		totalscore += float64(v)
	}

	totalscore = totalscore / float64(len(score))

	fmt.Printf("Total Score: %v\n", totalscore)

}

func main() {
	intro()
	inputName()
	fmt.Printf("Good day %v, lets take a quiz\n", name)
	newQuiz := Quiz{
		makeQuestion("is nash a good boy?\n", "", "yes"),
		makeQuestion("am i a good dev?\n", "", "it depends"),
		makeQuestion("who's the 2024 college football national champs?\n", "", "Tennessee"),
	}
	runQuiz(newQuiz)
}
