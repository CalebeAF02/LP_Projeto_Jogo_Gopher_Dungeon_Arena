# Roteiro para Vídeo de Apresentação

## Objetivo

Guiar a apresentação de 10 minutos do projeto **Gopher Dungeon Arena** de forma clara, técnica e organizada. O foco deve ser explicar a solução, mostrar o funcionamento e destacar a linguagem Go quando pertinente.

## Estrutura do vídeo (10 minutos)

1. Introdução — 1 minuto
   - Apresentar nome do projeto: **Gopher Dungeon Arena**
   - Explicar rapidamente o objetivo: jogo 2D de arena construído em Go
   - Dizer o que será mostrado: arquitetura, mecânicas, paralelismo e resultados

2. Visão geral do projeto — 1 minuto
   - Descrever brevemente o tema do jogo: dungeon, bots, jogador e mini mapa
   - Mostrar as principais cenas: menu, jogo, pausa e progresso
   - Indicar que o jogo foi feito com **Ebiten** em Go

3. Arquitetura e organização do código — 2 minutos
   - Explicar a estrutura de pastas:
     - `main.go`
     - `src/game.go`
     - `src/cenas/`
     - `src/sistema/`
     - `src/entidades/`
   - Explicar o padrão: **ECS** (Entity Component System)
   - Mostrar onde fica o loop de atualização (`Update`) e desenho (`Draw`)

4. Mecânicas principais do jogo — 2 minutos
   - Falar sobre os sistemas de atualização:
     - `SistemaInput`
     - `SistemaIA`
     - `SistemaSpawn`
     - `SistemaMovimento`
     - `SistemaEntidades`
     - `SistemaDebug`
   - Explicar como funciona o controle do jogador e a movimentação dos bots
   - Mencionar o sistema de colisão e como ele detecta interações

5. Destaque de Go e concorrência — 2 minutos
   - Explicar o diferencial de Go para este projeto:
     - `goroutines`
     - `channels`
     - `sync.WaitGroup`
     - `sync.RWMutex`
   - Falar sobre a evolução planejada do motor paralelo
   - Mostrar o documento em `docs/paralelismo_go.md` como roteiro técnico

6. Demonstração prática — 1 minuto
   - Mostrar o jogo rodando (ou capturas/trechos gravados)
   - Destacar fluxo de jogo:
     - iniciar partida
     - movimentar personagem
     - bots se movimentando
     - mini mapa
     - quando vence/perde

7. Conclusão e próximos passos — 1 minuto
   - Resumir o que foi entregue
   - Destacar o uso de Go como escolha técnica
   - Apontar melhorias futuras: paralelismo avançado, IA, multiplayer
   - Terminar com agradecimento e convite para análise do código

## Recomendações para gravação

- Fale devagar e com voz clara
- Use tela compartilhada com o editor e o jogo executando
- Mostre trechos do código enquanto explica pontos técnicos
- Use o documento `docs/paralelismo_go.md` como guia de roteiro
- Se possível, grave o jogo rodando em um ambiente estável

## Pontos a destacar no vídeo

- O projeto foi implementado em **Go**
- O motor usa **Ebiten** para renderizar e rodar o loop de jogo
- A arquitetura se baseia em um sistema de cenas e ECS
- Há planejamento para evoluir o motor para paralelismo seguro
- O professor deve ver o diferencial da linguagem Go no projeto

---

## Sugestão de tempo por bloco

- 0:00–1:00: introdução
- 1:00–2:00: visão geral do projeto
- 2:00–4:00: arquitetura e organização do código
- 4:00–6:00: mecânicas e sistemas do jogo
- 6:00–8:00: diferencial Go e concorrência
- 8:00–9:00: demonstração
- 9:00–10:00: conclusão
