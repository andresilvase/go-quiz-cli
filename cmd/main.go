package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var knowledgeBase map[int]string = map[int]string{
	1: "../internal/topics/biblia.csv",
	2: "../internal/topics/ciencia.csv",
	3: "../internal/topics/historia.csv",
	4: "../internal/topics/matematica.csv",
}

type GameState struct {
	PlayerName string
	Score      int
	Questions  []Question
}

type Question struct {
	Text    string
	Options []string
	Answer  int
}

func (g *GameState) ReadPlayerName() {
	fmt.Print("Digite o seu nome para continuar: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("\033[31merro ao ler o nome do jogador! %w\033[0m", err)
	}

	g.PlayerName = name[:len(name)-1]
}

func (g *GameState) FileReader(topic int) error {
	var topicPath string

	if path, ok := knowledgeBase[topic]; ok {
		topicPath = path
	} else {
		topicPath = knowledgeBase[3]
	}

	csvFile, err := os.Open(topicPath)

	if err != nil {
		return fmt.Errorf("ocorreu um erro ao abrir arquivo %w", err)
	}

	defer csvFile.Close()

	fileReader := csv.NewReader(csvFile)

	csvRows, err := fileReader.ReadAll()

	if err != nil {
		return fmt.Errorf("ocorreu um erro ao ler o arquivo %w", err)
	}

	for index, record := range csvRows {
		if index > 0 {
			convertedNumber, _ := toInt(record[len(record)-1])
			question := &Question{
				Text:    record[0],
				Options: record[1 : len(record)-1],
				Answer:  convertedNumber,
			}

			g.Questions = append(g.Questions, *question)
		}
	}

	return nil
}

func (g *GameState) Init() {
	reader := bufio.NewReader(os.Stdin)

	for index, question := range g.Questions {
		fmt.Printf("\n\033[33m%d. %s\n\033[0m", index+1, question.Text)
		for i, v := range question.Options {
			fmt.Printf("[%d] %s\n", i+1, v)
		}

		fmt.Printf("Digite a opção correta: ")
		readyToGo := false
		for !readyToGo {
			value, _ := reader.ReadString('\n')
			userAnswer, _ := toInt(value)

			readyToGo = userAnswer > 0 && userAnswer <= len(question.Options)

			if !readyToGo {
				fmt.Printf("Digite um valor númerico entre 1 e %d. ", len(question.Options))
			}

			if userAnswer == question.Answer {
				g.Score++
			}
		}
	}
}

func toInt(value string) (int, error) {
	if strings.Contains(value, "\n") {
		value = value[:len(value)-1]
	}

	intValue, err := strconv.Atoi(value)

	if err != nil {
		return 0, errors.New("você precisa digitar um valor numérico.\n")
	}

	return intValue, nil
}

func main() {
	fmt.Println("Seja Bem vindo!")

	game := &GameState{}
	game.ReadPlayerName()

	fmt.Printf("\nOlá \033[32m%s\033[0m! Vamos jogar?\n\n", game.PlayerName)
	fmt.Println("Escolha um dos temas abaixo.")

	for index, value := range knowledgeBase {
		topicName := strings.Split(strings.Split(value, "/")[3], ".csv")[0]

		fmt.Printf("[%d] %s \n", index, strings.ToUpper(topicName[0:1])+topicName[1:])
	}

	fmt.Print("Digite: ")

	topicSelected := false
	topic := 0

	for !topicSelected {
		reader := bufio.NewReader(os.Stdin)
		userChoice, _ := reader.ReadString('\n')
		userChoiceInt, err := toInt(userChoice)

		if err != nil {
			fmt.Printf("\nError: %v", err)
		}

		topic = userChoiceInt
		topicSelected = topic > 0 && topic <= len(knowledgeBase)

		if !topicSelected {
			fmt.Printf("Escolha uma opção válida entre 1 e %d: ", len(knowledgeBase))
		}
	}

	if err := game.FileReader(topic); err != nil {
		fmt.Printf("\nError: %v", err)
	}

	game.Init()
	fmt.Println("Fim de jogo!")
	fmt.Printf("\nVocê acertou %d de %d questões", game.Score, len(game.Questions))
}
