# Go Quiz CLI 🧠

**Português** | [English](README.md)

![Go](https://img.shields.io/badge/Go-1.23.2-00ADD8?logo=go&logoColor=white)
![Interface](https://img.shields.io/badge/interface-CLI-4EAA25?logo=gnubash&logoColor=white)
![Status](https://img.shields.io/badge/status-projeto%20de%20estudo-f59e0b)

Um quiz pequeno e interativo para o terminal — e um ambiente prático para estudar e se manter atualizado com a **linguagem de programação Go**.

Em vez de aprender Go apenas com exemplos isolados, este projeto aplica a linguagem na construção de um programa de linha de comando completo: ele carrega perguntas de arquivos CSV, valida as entradas do usuário, gerencia o estado do jogo e calcula a pontuação final.

## O que você pode praticar

- Structs, métodos, slices e maps
- Ponteiros e gerenciamento de estado
- Leitura e processamento de arquivos CSV
- Entrada de dados e saída formatada no terminal
- Conversão de tipos e validação de entradas
- Criação, encapsulamento e propagação de erros
- Organização de um projeto Go em comandos e dados internos

## Funcionalidades

- Quatro categorias: **Bíblia**, **Ciência**, **História** e **Matemática**
- Perguntas de múltipla escolha carregadas de arquivos CSV
- Validação da categoria e das respostas selecionadas
- Textos coloridos no terminal
- Exibição da pontuação após três perguntas
- Formato de dados simples, facilitando a inclusão de perguntas e categorias

## Como executar

### Pré-requisito

- [Go 1.23.2 ou mais recente](https://go.dev/doc/install)

Clone o repositório e inicie o jogo pelo diretório do comando:

```bash
git clone https://github.com/andresilvase/go-quiz-cli.git
cd go-quiz-cli/cmd
go run .
```

O programa solicitará seu nome, a categoria desejada e as respostas para três perguntas:

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

> Atualmente, o programa utiliza caminhos relativos ao diretório `cmd`. Por isso, execute-o a partir desse diretório.

## Estrutura do projeto

```text
go-quiz-cli/
├── cmd/
│   └── main.go             # Fluxo da CLI e lógica do jogo
├── internal/
│   └── topics/             # Bancos de perguntas em CSV
│       ├── biblia.csv
│       ├── ciencia.csv
│       ├── historia.csv
│       └── matematica.csv
└── go.mod
```

## Como adicionar perguntas

Cada categoria possui um arquivo CSV com uma pergunta por linha:

```csv
Pergunta,Opção 1,Opção 2,Opção 3,Opção 4,Resposta
Quanto é 7 x 8?,48,54,56,64,3
```

A última coluna guarda o número da alternativa correta, começando em `1`. Para ampliar uma categoria existente, adicione uma linha seguindo o mesmo formato.

Para criar uma categoria completamente nova:

1. Adicione um arquivo CSV em `internal/topics/`.
2. Registre o caminho do arquivo no map `knowledgeBase`, em `cmd/main.go`.
3. Execute a aplicação e selecione a nova categoria.

## Ideias para continuar praticando

Este repositório é intencionalmente pequeno, tornando-o uma boa base para experimentar novos conceitos e versões do Go. Algumas ideias para evoluí-lo:

- Usar `embed` para executar o quiz a partir de qualquer diretório
- Embaralhar perguntas e alternativas com `math/rand`
- Adicionar um limite configurável de perguntas e um cronômetro
- Separar as lógicas de jogo, entrada e armazenamento em pacotes
- Criar testes unitários orientados a tabelas
- Persistir recordes em JSON ou em um banco de dados
- Tornar as cores do terminal opcionais para melhorar a acessibilidade
- Gerar binários de release por meio de um fluxo de integração contínua

## Por que este projeto existe?

A melhor forma de fixar Go é construir, revisitar e aprimorar programas reais. O **Go Quiz CLI** é um projeto prático de estudo: pequeno o suficiente para ser compreendido de uma só vez, mas completo o bastante para exercitar fundamentos encontrados em aplicações de produção.

Faça um fork, experimente, quebre, melhore e use este projeto para testar tudo o que aprender em Go. 🐹

## Licença

Este projeto ainda não possui uma licença. Caso pretenda reutilizá-lo ou distribuí-lo, adicione uma licença compatível com seus objetivos.
