package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cohere-ai/cohere-go"
)

type QA struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type QAs struct {
	Questions []QA `json:"questions"`
}

func (qas *QAs) GetQuestionsList() ([]string) {
	var questions []string

	for _, qa := range qas.Questions {
		questions = append(questions, qa.Question)
	}

	return questions
}

func (qas *QAs) GetAnswersList() ([]string) {
	var answers []string

	for _, qa := range qas.Questions {
		answers = append(answers, qa.Answer)
	}

	return answers
}

func UnmarshalQuestions (data []byte) (QAs, error) {
	var qas QAs
	err := json.Unmarshal(data, &qas)
	return qas, err
}

func EmbedQuestions(questions []string) ([][]float64, error) {
	CO_API_KEY := os.Getenv("CO_API_KEY")
  
	co, err := cohere.CreateClient(CO_API_KEY)
	if err != nil {
	  fmt.Println(err)
	  return nil, err
	}
  
	response, err := co.Embed(cohere.EmbedOptions{
	  Model: "large",
	  Texts: questions,
	})
	if err != nil {
	  fmt.Println(err)
	  return nil, err
	}
  
	// fmt.Println("Embeddings:", response.Embeddings[0])
	return response.Embeddings, err
}