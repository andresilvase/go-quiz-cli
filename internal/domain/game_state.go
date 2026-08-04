package domain

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/andresilvase/go-quiz-cli/internal/utils"
)

var KnowledgeBase map[int]string = map[int]string{
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

func (g *GameState) ReadPlayerName() error {
	fmt.Print("Digite o seu nome para continuar: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		return fmt.Errorf("erro ao ler o nome do jogador! %w", err)
	}

	g.PlayerName = name[:len(name)-1]

	return nil
}

func (g *GameState) FileReader(topic int) error {
	var topicPath string

	if path, ok := KnowledgeBase[topic]; ok {
		topicPath = path
	} else {
		topicPath = KnowledgeBase[3]
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
			convertedNumber, err := utils.ToInt(record[len(record)-1])

			if err != nil {
				return fmt.Errorf("ocorreu um erro ao preparar o csv para o quiz %w", err)
			}

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

func (g *GameState) Init() error {
	for index, question := range g.Questions {
		fmt.Printf("\n\033[33m%d. %s\n\033[0m", index+1, question.Text)
		for i, v := range question.Options {
			fmt.Printf("[%d] %s\n", i+1, v)
		}

		fmt.Printf("Digite a opção correta: ")
		readyToGo := false
		for !readyToGo {
			var input string
			_, err := fmt.Scanln(&input)

			if err != nil {
				return fmt.Errorf("\nInit Error: %w", err)
			}

			userAnswer, err := utils.ToInt(input)

			if err != nil {
				fmt.Printf("\nError: %v", err)
			}

			readyToGo = userAnswer > 0 && userAnswer <= len(question.Options)

			if !readyToGo {
				fmt.Printf("Digite um valor númerico entre 1 e %d. ", len(question.Options))
			}

			if userAnswer == question.Answer {
				g.Score++
			}
		}
	}

	return nil
}

func (g *GameState) End() {
	fmt.Println("\033[32m\nFim de jogo!\033[0m")
	fmt.Printf("Você acertou %d de %d questões.\n", g.Score, len(g.Questions))
}
