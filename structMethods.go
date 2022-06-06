package main

import "fmt"

type question struct{
	question string
	answer string
}

func (q *question) getQuestion() string{
	return q.question
}

func (q question) getAnswer() string{
	return q.answer
}

func tryChangingQuestions(q question){
	q.answer = "changed?"
}

func changeQuestions(q *question){
	q.answer = "changed"
}

func makeQuiz(){
	question1 := question{question: "Who is that", answer: "I dont know"}
	fmt.Println(question1.getQuestion())
	fmt.Println(question1.getAnswer())

	fmt.Println("Trying to change the questions in a function")
	tryChangingQuestions(question1)
	fmt.Println(question1.getQuestion())
	fmt.Println(question1.getAnswer())


	question1Pointer := &question1
	fmt.Println(question1Pointer.getQuestion())
	fmt.Println(question1Pointer.getAnswer())

	fmt.Println("Changing the questions by passing the struct reference in a function")
	changeQuestions(question1Pointer)
	fmt.Println(question1.getQuestion())
	fmt.Println(question1.getAnswer())
}