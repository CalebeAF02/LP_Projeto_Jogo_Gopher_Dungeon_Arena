# Gopher Dungeon Arena

Gopher Dungeon Arena é um projeto acadêmico desenvolvido em Go para a disciplina de Linguagens de Programação da UnB. O objetivo é criar um jogo 2D de arena com mecânicas de movimentação, colisão, câmera, minimapa, inimigos e progressão de fases.

## Visão geral

O projeto foi pensado como uma aplicação prática para explorar conceitos de programação, arquitetura de software e desenvolvimento de jogos em Go. A implementação atual já possui uma estrutura funcional com menu inicial, cena de jogo, entidades, colisões básicas, bots com diferentes tipos de movimentação e persistência de progresso.

## Funcionalidades implementadas

- Menu inicial e telas de navegação
- Cena principal de jogo
- Jogador controlável
- Bots com diferentes comportamentos de movimentação
- Colisões básicas com paredes e objetos do mapa
- Comida, saída, portais e minimapa
- Câmera dinâmica seguindo o jogador
- Persistência de progresso em arquivo JSON

## Tecnologias

- Go
- Ebiten para renderização gráfica
- Estruturas modulares inspiradas em ECS
- Arquitetura baseada em cenas, interfaces e sistemas

## Como executar

No diretório do projeto, execute:

```bash
go run .
```

Para compilar o projeto:

```bash
go build ./...
```

## Controles principais

- Enter: iniciar o jogo
- Esc: voltar ao menu ou sair, conforme a tela
- Setas ou WASD: movimentar o jogador
- P: pausar o jogo
- Ctrl + M: alternar a posição do minimapa
- Ctrl + O: mostrar ou ocultar o minimapa

## Estrutura do projeto

- [main.go](main.go): ponto de entrada do jogo
- [src/game.go](src/game.go): controle principal do fluxo do jogo
- [src/cenas](src/cenas): telas e gerenciamento de cenas
- [src/entidades](src/entidades): personagens e objetos do mundo
- [src/sistema](src/sistema): lógica de atualização e comportamento
- [src/nivel](src/nivel): carregamento de fases e progresso

## Documentação

- [Documentação do jogo](docs/tecnico/documentacao_jogo.md)
- [Índice da documentação](docs/README.md)
- [Sumário](SUMARIO.md)

## Status

O projeto já está compilando corretamente e possui uma base funcional para evolução em gameplay, IA e refinamento técnico.
