package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	domain "github.com/andresilvase/go-quiz-cli/internal/domain"
	utils "github.com/andresilvase/go-quiz-cli/internal/utils"
)

func displayTopicChoice() {
	fmt.Println("Escolha um dos temas abaixo.")

	for index, value := range domain.KnowledgeBase {
		topicName := strings.Split(strings.Split(value, "/")[3], ".csv")[0]

		fmt.Printf("[%d] %s \n", index, strings.ToUpper(topicName[0:1])+topicName[1:])
	}

	fmt.Print("Digite: ")
}

func readTopicChoice() (int, error) {
	topicSelected := false
	topic := 0

	for !topicSelected {
		reader := bufio.NewReader(os.Stdin)
		userChoice, err := reader.ReadString('\n')

		if err != nil {
			return 0, fmt.Errorf("\nreadTopicChoice Error: %w", err)
		}

		userChoiceInt, err := utils.ToInt(userChoice)

		if err != nil {
			fmt.Printf("\nError: %v", err)
		}

		topic = userChoiceInt
		topicSelected = topic > 0 && topic <= len(domain.KnowledgeBase)

		if !topicSelected {
			fmt.Printf("Escolha uma opção válida entre 1 e %d: ", len(domain.KnowledgeBase))
		}
	}

	return topic, nil
}

func main() {
	game := &domain.GameState{}

	fmt.Println("Seja Bem vindo!")

	game.ReadPlayerName()
	fmt.Printf("\nOlá \033[32m%s\033[0m! Vamos jogar?\n\n", game.PlayerName)

	displayTopicChoice()

	topic, err := readTopicChoice()

	if err != nil {
		log.Fatal("\nError: %w", err)
	}

	if err := game.FileReader(topic); err != nil {
		log.Fatal("\nError: %w", err)
	}

	game.Init()
	game.End()
}
