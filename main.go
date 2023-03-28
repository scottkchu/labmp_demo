package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	jsonBytes, err := ioutil.ReadFile("questions.json")
	if err != nil {
		panic(err)
	}

	qas, err := UnmarshalQuestions(jsonBytes)
	if err != nil {
		panic(err)
	}

	questions := qas.GetQuestionsList()
	answers := qas.GetAnswersList()
	allQuestionsEmbeddings, err := EmbedQuestions(questions)
	if err != nil {
		panic(err)
	}

	question := []string{"i want to return these pants"}
	questionEmbeddings, err := EmbedQuestions(question)
	if err != nil {
		panic(err)
	}

	bestMatchIndex, bestMatchScore := GetBestMatch(questionEmbeddings[0], allQuestionsEmbeddings)
	fmt.Println("Best match:", questions[bestMatchIndex])
	fmt.Println("Best match score:", bestMatchScore)
	fmt.Println("Best match answer:", answers[bestMatchIndex])
}