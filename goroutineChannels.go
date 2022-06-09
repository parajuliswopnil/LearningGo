package main

import (
	"fmt"
	"time"
)

type Questions struct{
	question string
	answer []string
}

func sendSomethingInChannel()int{
	initialTime := time.Now().Second()
	nextTime := initialTime + 2
	for initialTime < nextTime{
		time.Sleep(time.Second)
		present := time.Now().Second()
		if present < nextTime{
			continue
		} else{
			time.Sleep(time.Second + time.Duration(10000))
			return 1
		}
	}
	
	return 0
}

func (q Questions)askQuestion()(string, []string){
	return q.question, q.answer
}

func demoChannel() {
	questions1 := Questions{question: "Highest mountain?", answer: []string{"nepal", "india", "china"}}
	questions2 := Questions{question: "Tallest animal?", answer: []string{"lion", "tiger", "giraffe"}}
	questions3 := Questions{question: "Fastest animal?", answer: []string{"tiger", "cheetah", "rhino"}}
	questions4 := Questions{question: "Largest animal?", answer: []string{"elephant", "snake", "whale"}}

	questionList := []Questions{questions1, questions2, questions3, questions4}
	channel := make(chan int)
	for i := 0; i < len(questionList); i++{
		questionToAsk := questionList[i]
		q, a := questionToAsk.askQuestion()
		fmt.Println(q)
		fmt.Println(a)
	
	for i := 0; i < 10; i--{
		go func()  {
			channel <- sendSomethingInChannel()
		}()

		channelMessage := <- channel

		if channelMessage == 1{
			break
		}


	fmt.Println(channelMessage)
		}
	}
}