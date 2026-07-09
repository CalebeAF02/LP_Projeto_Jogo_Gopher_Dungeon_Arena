# Guia de slides para a apresentação do vídeo

## Objetivo
Este documento serve como base para criar os slides do vídeo de apresentação do projeto Gopher Dungeon Arena. A ideia é manter uma apresentação clara, visual e alinhada com o roteiro de 10 minutos.

## Estrutura sugerida dos slides

### Slide 1 — Título
- Título: Gopher Dungeon Arena
- Subtítulo: Jogo 2D de arena desenvolvido em Go
- Imagem: print do jogo ou tela inicial

Texto de apoio:
- Apresentar o projeto e a proposta principal de forma direta.

Roteiro falado:
- "Olá, este projeto é o Gopher Dungeon Arena, um jogo 2D de arena desenvolvido em Go, com mecânicas de combate, movimentação e interação entre jogador e bots."

---

### Slide 2 — O que é o projeto
- Tema do jogo: dungeon, arena, combate e sobrevivência
- Elementos principais: jogador, bots, mini mapa, portais e objetivos
- Destaque: jogo com foco em gameplay e estrutura organizada

Texto de apoio:
- Mostrar o contexto geral do jogo e o que o usuário pode experimentar.

Roteiro falado:
- "A ideia central do projeto é criar uma experiência de arena em que o jogador precisa sobreviver, se movimentar estrategicamente e enfrentar bots em um ambiente dinâmico."

---

### Slide 3 — Funcionalidades principais
- Menu inicial
- Tela de jogo
- Menu de pausa
- Tela de progresso e resultado
- Mecânicas de movimentação, colisão e combate

Texto de apoio:
- Mostrar que o jogo possui uma estrutura completa, com fluxo de partida e estados bem definidos.

Roteiro falado:
- "O jogo conta com diferentes cenas, como menu, partida, pausa e tela de progresso, além de mecânicas de movimentação, combate e interação com o ambiente."

---

### Slide 4 — Arquitetura do projeto
- Estrutura principal do código
- Pastas principais: main.go, src/game.go, src/cenas, src/sistema, src/entidades
- Conceito de ECS: entidades, componentes e sistemas
- Loop principal de atualização e desenho

Texto de apoio:
- Explicar de forma simples como o projeto foi organizado.

Roteiro falado:
- "A organização do projeto foi pensada para separar responsabilidades. O código está dividido em cenas, sistemas e entidades, seguindo um modelo baseado em ECS, que ajuda a manter a lógica do jogo mais modular."

---

### Slide 5 — Sistemas importantes
- Sistema de input
- Sistema de IA
- Sistema de spawn
- Sistema de movimento
- Sistema de colisão
- Sistema de entidades

Texto de apoio:
- Mostrar que o gameplay é organizado em módulos.

Roteiro falado:
- "Os principais sistemas do jogo cuidam da entrada do jogador, da inteligência dos bots, da criação de elementos na arena, da movimentação e das colisões."

---

### Slide 6 — Diferencial do projeto em Go
- Uso de Go como linguagem principal
- Bibliotecas e ferramentas: Ebiten
- Concurrência e paralelismo como parte do projeto
- Goroutines, channels e sincronização

Texto de apoio:
- Destacar por que Go foi uma escolha relevante para este projeto.

Roteiro falado:
- "Uma das grandes escolhas técnicas deste projeto foi usar Go. Além de permitir a construção do jogo com uma estrutura limpa, a linguagem também abriu espaço para explorar conceitos de concorrência e paralelismo de forma interessante."

---

### Slide 7 — Demonstração do jogo
- Captura do jogo rodando
- Fluxo: iniciar partida, movimentar, bots reagindo, vencer ou perder
- Destaque do mini mapa e das interações

Texto de apoio:
- Mostrar o jogo funcionando para deixar a apresentação mais concreta.

Roteiro falado:
- "Agora, para fechar, vamos ver o jogo em funcionamento, mostrando o fluxo principal da partida e os elementos que foram implementados."

---

### Slide 8 — Conclusão e próximos passos
- Projeto entregue com estrutura organizada
- Go como diferencial técnico
- Melhorias futuras: IA mais avançada, paralelismo, multiplayer e refinamento do gameplay

Texto de apoio:
- Finalizar com uma mensagem de conclusão e visão de evolução.

Roteiro falado:
- "No geral, o projeto entrega uma base forte para um jogo de arena em Go, com uma arquitetura organizada e boas possibilidades de evolução futura."

---

## Sugestões visuais para os slides
- Use screenshots do jogo em vez de muitos textos.
- Em slides técnicos, mostre trechos do código ou diagramas simples.
- Use cores consistentes e fontes limpas.
- Evite preencher os slides com excesso de informação.

## Sugestão de ordem de apresentação
1. Título
2. O que é o projeto
3. Funcionalidades principais
4. Arquitetura do projeto
5. Sistemas importantes
6. Diferencial do projeto em Go
7. Demonstração
8. Conclusão

## Observação final
Essa estrutura funciona bem para um vídeo de apresentação porque alterna entre explicação do contexto, parte técnica e demonstração prática.
