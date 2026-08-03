# Go Quiz CLI 🧠

[Português](README.pt.md) | **English**

![Go](https://img.shields.io/badge/Go-1.23.2-00ADD8?logo=go&logoColor=white)
![Interface](https://img.shields.io/badge/interface-CLI-4EAA25?logo=gnubash&logoColor=white)
![Status](https://img.shields.io/badge/status-learning%20project-f59e0b)

A small, interactive quiz for the terminal—and a hands-on playground for staying up to date with the **Go programming language**.

Rather than learning Go only through isolated examples, this project applies the language to a complete command-line program: it reads questions from CSV files, validates user input, manages game state, and calculates a final score. The interface and quiz content are in Portuguese.

## What you can practice

- Structs, methods, slices, and maps
- Pointers and state management
- Reading and parsing CSV files
- Terminal input and formatted output
- Type conversion and input validation
- Error creation, wrapping, and propagation
- Organizing a Go project into commands and internal data

## Features

- Four quiz categories: **Bible**, **Science**, **History**, and **Mathematics**
- Multiple-choice questions loaded from CSV files
- Validation for topic and answer selections
- Colored terminal feedback
- A final score after all questions from the selected category
- A simple data format that makes new topics and questions easy to add

## Run locally

### Prerequisite

- [Go 1.23.2 or newer](https://go.dev/doc/install)

Clone the repository and start the game from the command directory:

```bash
git clone https://github.com/andresilvase/go-quiz-cli.git
cd go-quiz-cli/cmd
go run .
```

You will be prompted to enter your name, select a topic, and answer every question available in that category:

```text
Seja Bem vindo!
Digite o seu nome para continuar: Gopher

Olá Gopher! Vamos jogar?

Escolha um dos temas abaixo.
[1] Bíblia
[2] Ciência
[3] História
[4] Matemática
```

> The program currently uses paths relative to `cmd`, so run it from that directory.

## Project structure

```text
go-quiz-cli/
├── cmd/
│   └── main.go             # CLI flow and game logic
├── internal/
│   └── topics/             # CSV question banks
│       ├── biblia.csv
│       ├── ciencia.csv
│       ├── historia.csv
│       └── matematica.csv
└── go.mod
```

## Add your own questions

Each topic is a CSV file with one question per row:

```csv
Pergunta,Opção 1,Opção 2,Opção 3,Opção 4,Resposta
Quanto é 7 x 8?,48,54,56,64,3
```

The last column contains the number of the correct option, starting at `1`. To extend an existing topic, add another row using the same format.

To create a completely new category:

1. Add a CSV file under `internal/topics/`.
2. Register its path in the `knowledgeBase` map in `cmd/main.go`.
3. Run the application and select the new category.

## Ideas for continued practice

This repository is intentionally small, making it a useful base for experimenting with new Go concepts and releases. Good next challenges include:

- Use `embed` so the quiz can run from any directory
- Shuffle questions and answer choices with `math/rand`
- Add a configurable question limit and countdown timer
- Split the game, input, and storage logic into packages
- Add table-driven unit tests
- Persist high scores as JSON or in a database
- Improve accessibility by making terminal colors optional
- Build release binaries with a CI workflow

## Why this project exists

Go is easiest to retain by building, revisiting, and improving real programs. **Go Quiz CLI** is a practical study project: small enough to understand in one sitting, but rich enough to exercise the language fundamentals that appear in production applications.

Fork it, break it, improve it, and use it to try the next thing you learn in Go. 🐹

## License

No license has been added yet. If you plan to reuse or distribute this project, add a license that matches your goals.
