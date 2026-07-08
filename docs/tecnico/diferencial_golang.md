# O que o professor quer ver em Go

## 1. Objetivo

Este documento explica o que deve ser apresentado ao professor sobre a linguagem **Go**, focando no diferencial da linguagem e não em recursos de outras linguagens como C.

O objetivo é mostrar porque Go foi uma escolha valiosa para este projeto e quais recursos específicos de Go foram explorados.

---

## 2. Foco do professor

O professor quer ver:
- características próprias de **Go**
- o que a linguagem oferece e facilita
- por que Go é diferente de C
- como a linguagem ajuda na construção do jogo

O professor não quer uma comparação direta com o que já existe em C. Ele quer entender o diferencial de Go: suas ferramentas, estilo e vantagens.

---

## 3. Principais diferenciais de Go para o projeto

### 3.1. Concorrência integrada

Go traz **concorrência nativa** como parte da linguagem, em vez de ser apenas uma biblioteca:
- `goroutine` para executar tarefas leves concorrentemente
- `channel` para comunicação segura entre rotinas
- `sync.WaitGroup` para sincronização de grupos de goroutines
- `sync.Mutex` / `sync.RWMutex` para proteger o estado compartilhado

No projeto, essa pilha é ideal para IA e para evoluir o motor de jogo para múltiplos bots e multiplayer.

### 3.2. Sintaxe simples e legível

Go tem uma sintaxe enxuta e sem muitos símbolos complexos, o que ajuda em:
- leitura rápida do código
- manutenção do projeto
- escrita de sistemas de jogo com menos boilerplate

Exemplos no projeto:
- declaração de structs simples
- funções curtas e diretas
- interfaces claras para sistemas e entidades

### 3.3. Interfaces explícitas e tipagem segura

Go permite definir interfaces simples e poderosas, como:
- `ISistemaAtualizar`
- `ICenaJogo`
- `Entidade`

Isso facilita a arquitetura ECS sem precisar de herança. A tipagem forte de Go evita muitos erros em tempo de compilação.

### 3.4. Gerenciamento de pacotes e compilação rápida

Go tem um sistema de módulos integrado (`go.mod`) que simplifica:
- dependências
- versionamento
- compilação

O projeto usa `go.mod` para gerenciar o módulo principal e as importações locais.

### 3.5. Recursos de desenvolvimento e depuração

Go oferece ferramentas de linha de comando úteis:
- `go test`
- `go test -race`
- `go run`
- `go fmt`

Essas ferramentas ajudam a manter qualidade e detectar condições de corrida antes de entregar o projeto.

---

## 4. O que mostrar na apresentação ao professor

### 4.1. Código que demonstra o diferencial de Go

- a estrutura de `main.go` e como Ebiten é usado
- `src/game.go` com `Update()` e `Draw()` simples
- `src/interfaces/` com contratos claros para cena e sistemas
- `src/entidades/personagens/bot.go` com métodos de movimento e componentes

### 4.2. Concorrência planejada

- explique o plano de paralelizar `SistemaIA`
- cite `goroutines` para bots
- explique `sync.WaitGroup` para sincronização por frame
- mencione o uso futuro de `sync.RWMutex` e `channels`

### 4.3. Arquitetura idiomática de Go

- destaque o uso de pacotes organizados
- explique como `interfaces` permitem desacoplamento
- cite o padrão ECS como fácil de implementar em Go

### 4.4. Vantagens sobre C em contexto de projeto

Não é necessário dizer “Go é melhor que C”, mas você pode argumentar:
- C não tem `goroutine`/`channel` nativos
- Go tem runtime leve que simplifica concorrência
- Go oferece coleta de lixo e tipagem segura, diminuindo erros de memória
- O pacote padrão de Go já traz suporte para sincronização e modularidade

---

## 5. Exemplo prático para apresentar

### 5.1. Demonstração do loop de jogo

- `main.go` chama `ebiten.RunGame(game)`
- `src/game.go` delega `Update()` e `Draw()` para a cena corrente
- `src/cenas/cenaJogo.go` executa os sistemas em sequência

### 5.2. Demonstração de um sistema em Go

- mostrar `src/sistema/sistema_de_movimento.go`
- explicar que é simples criar um `WaitGroup` e paralelizar execuções
- mostrar estrutura de `Bot.Atualizar()` e como ela segue o estilo Go

---

## 6. Como o professor deve avaliar o conteúdo

O professor deve reconhecer que você apresentou:
- conhecimento da linguagem Go
- uso dos conceitos nativos da linguagem
- entendimento de concorrência em Go
- aplicação de Go em um projeto de jogos
- arquitetura modular sustentada por interfaces

---

## 7. Recomendações para o documento

- mantenha o foco em **Go**, não em C
- cite recursos próprios da linguagem
- dê exemplos de sintaxe e padrões Go usados no código
- indique claramente que o projeto usa Go como base técnica principal
- use [docs/tecnico/paralelismo_go.md](tecnico/paralelismo_go.md) como suporte técnico para a parte de concorrência
