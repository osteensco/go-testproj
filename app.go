package main

import "fmt"

func intro() {
	fmt.Println("Hey man, here's a quiz and stuff")
}

var name string

func inputName() {
	fmt.Print("Input your name:")
	fmt.Scan(&name)
}

type QuizProblem struct {
	question string
	input    string
	answer   string
}

func runQuiz() {

}

func main() {
	intro()
	inputName()
	fmt.Printf("Good day %v, lets take a quiz", name)
}
