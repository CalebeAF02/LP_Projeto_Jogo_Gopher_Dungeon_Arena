# Gopher Dungeon Arena

## Wiki definitiva do projeto

Este documento consolida a análise completa do projeto Gopher Dungeon Arena e serve como base para uma wiki oficial no GitHub. Ele reúne a visão geral do jogo, a arquitetura implementada, as mecânicas atuais, a estrutura do código e os próximos passos de evolução.

> Última atualização: 2026-07-08

---

## 1. Visão geral do projeto

Gopher Dungeon Arena é um jogo 2D desenvolvido em Go como projeto acadêmico para a disciplina de Linguagens de Programação, na UnB. A proposta central é demonstrar como uma linguagem de programação compilada, com foco em desempenho e simplicidade, pode ser usada para construir um pequeno motor de jogo com lógica de cenas, entidades, colisões, IA básica de inimigos, câmera, minimapa e persistência de progresso.

O projeto já está em uma fase funcional, com um fluxo completo de menu, cena de jogo, movimentação do jogador, presença de bots com diferentes comportamentos, interação com objetos do mapa e condição de vitória/derrota.

---

## 2. Objetivo acadêmico e técnico

O projeto foi pensado para explorar conceitos importantes de engenharia de software e programação:

- modularização e separação de responsabilidades;
- uso de interfaces para organizar cenas e sistemas;
- estrutura inspirada em ECS leve;
- desenvolvimento de lógica de jogo em um loop contínuo;
- manipulação de estado e persistência em arquivos JSON;
- uso de bibliotecas gráficas modernas, neste caso Ebiten.

Em termos práticos, o projeto busca mostrar que é possível criar uma base de jogo com organização razoável, mesmo sem depender de motores complexos tradicionais.

---

## 3. Tecnologias utilizadas

- Linguagem principal: Go
- Biblioteca gráfica: Ebiten
- Arquitetura: estrutura híbrida inspirada em ECS
- Persistência: arquivos JSON
- Organização: pacotes separados por responsabilidade

A implementação atual demonstra boa adoção dos recursos da linguagem, principalmente em relação à modularização, à tipagem explícita e ao uso de interfaces.

---

## 4. Arquitetura do projeto

### 4.1 Fluxo principal do jogo

O ponto de entrada do projeto está em [main.go](../../main.go). A execução começa com a criação do objeto principal de jogo, que inicializa a janela, define o título e roda o loop de atualização/desenho do Ebiten.

A lógica central do jogo está concentrada em [src/game.go](../../src/game.go), onde são controladas as transições entre cenas, como menu inicial, jogo, pause e retorno ao menu.

### 4.2 Máquina de estados de cenas

O projeto organiza a experiência do usuário em cenas. As principais telas são:

- menu inicial;
- cena de jogo;
- menu de pause;
- tela de progresso;
- telas informativas de vitória/derrota.

Essa abordagem facilita a evolução do jogo, pois cada tela pode ter sua própria responsabilidade e lógica.

### 4.3 Estrutura inspirada em ECS

Embora o projeto não implemente um ECS clássico completo, ele adota uma ideia semelhante de organização:

- entidades representam objetos do mundo;
- componentes guardam dados e estado;
- sistemas processam regras e comportamentos.

A interface básica de entidade está definida em [src/ecs/ecs.go](../../src/ecs/ecs.go), e serve como contrato para entidades como jogador, bots, comida, paredes e portais.

---

## 5. Estrutura de pastas

A organização atual do repositório é a seguinte:

- [main.go](../../main.go): ponto de entrada do programa
- [src/game.go](../../src/game.go): controlador principal do fluxo do jogo
- [src/cenas](../../src/cenas): telas e gerenciamento de cena
- [src/componentes](../../src/componentes): estruturas de dados e atributos do jogo
- [src/ecs](../../src/ecs): contrato básico de entidades
- [src/entidades](../../src/entidades): personagens, objetos e funcionalidades
- [src/interfaces](../../src/interfaces): abstrações para cenas e sistemas
- [src/sistema](../../src/sistema): sistemas de entrada, movimento, colisão, IA e desenho
- [src/nivel](../../src/nivel): carregamento de fases e persistência de progresso
- [src/config](../../src/config): constantes de janela, mapa e configuração geral
- [src/utils](../../src/utils): valores compartilhados e constantes de gameplay
- [docs](../../docs): documentação técnica e material de apresentação

---

## 6. Entidades e componentes principais

### 6.1 Jogador

O jogador é representado por [src/entidades/personagens/jogador.go](../../src/entidades/personagens/jogador.go). Ele é controlado pelo teclado e possui:

- corpo físico para colisão;
- vida e sangue;
- nível;
- pontuação;
- comportamento de movimento com colisão pixel a pixel.

### 6.2 Bots

Os inimigos são implementados em [src/entidades/personagens/bot.go](../../src/entidades/personagens/bot.go). Eles possuem diferentes tipos de movimentação, definidos por movimentadores específicos, e podem interagir com o ambiente e com o jogador.

### 6.3 Objetos do mapa

Os principais objetos do mundo são:

- paredes, que atuam como barreiras físicas;
- comida, que incrementa a pontuação;
- saída, que permite concluir a fase;
- portais de entrada e saída, que implementam o sistema de teletransporte.

Essas entidades estão localizadas em [src/entidades/objeto](../../src/entidades/objeto).

### 6.4 Componentes centrais

Os principais componentes usados pelo jogo são:

- CORPO: define posição e dimensão da entidade;
- VIDA: representa saúde, estado vivo/morto e sangue;
- NIVEL: armazena o nível da entidade;
- PONTUACAO: controla o progresso do jogador;
- MOVIMENTO: define o comportamento de deslocamento;
- ATIVIDADE: controla estados temporários, como movimento ou teletransporte;
- ENVIANDO_TELETRANSPORTE e RECEBENDO_TELETRANSPORTE: controlam a mecânica de portais.

---

## 7. Mecânicas atuais do jogo

### 7.1 Movimento

O jogador pode se mover com as setas ou com WASD. O movimento é processado de forma incremental e respeita a colisão com objetos e bordas do mundo.

### 7.2 Colisão

O sistema de colisão está implementado em [src/sistema/sistema_de_colisao.go](../../src/sistema/sistema_de_colisao.go). Ele verifica interações entre entidades e aplica regras específicas quando há choque entre tipos diferentes.

As colisões mais relevantes são:

- jogador com parede;
- bot com parede;
- jogador com bot;
- jogador com comida;
- jogador com saída;
- bot com portal de entrada.

### 7.3 Combate

A lógica de combate está organizada em [src/entidades/funcionalidades/combate.go](../../src/entidades/funcionalidades/combate.go). Quando jogador e bot colidem, ambos perdem sangue conforme as regras definidas no projeto.

### 7.4 Teletransporte

A mecânica de portais está implementada em [src/entidades/funcionalidades/teletransporte.go](../../src/entidades/funcionalidades/teletransporte.go). Quando um bot entra em um portal de entrada, o estado é marcado e a entidade passa a ser encaminhada para o portal de saída correspondente.

### 7.5 Progressão e vitória

A partida é concluída quando o jogador coleta os itens necessários e chega até a saída. O estado de progresso é atualizado e salvo para uso em execuções futuras.

### 7.6 Câmera e minimapa

O projeto conta com:

- câmera dinâmica seguindo o jogador;
- minimapa exibido na tela;
- alternância da posição do minimapa com atalhos de teclado.

Esses recursos são controlados em [src/ecs/camera.go](../../src/ecs/camera.go) e [src/ecs/miniMapa.go](../../src/ecs/miniMapa.go).

---

## 8. Sistemas implementados

### 8.1 Sistema de input

Responsável por capturar comandos do teclado e direcionar o movimento do jogador.

### 8.2 Sistema de IA

Os bots usam diferentes algoritmos de movimentação, organizados em movimentadores. O projeto já possui implementações para padrões horizontais, verticais, diagonais e variações mais lógicas.

### 8.3 Sistema de spawn

O sistema de spawn é usado para criar entidades do mapa e controlar a geração de elementos durante a partida.

### 8.4 Sistema de movimento

Responsável por atualizar a posição das entidades com base em suas regras internas.

### 8.5 Sistema de colisão

Responsável por verificar obstáculos e interações com o ambiente.

### 8.6 Sistema de desenho

Garante que o mundo, os objetos, a câmera e o HUD sejam renderizados corretamente na tela.

### 8.7 Sistema de debug

Inclui mecanismos para inspecionar o estado das entidades e facilitar a depuração do desenvolvimento.

### 8.8 Concorrência e paralelismo em Go

Um dos pontos mais interessantes do projeto é o uso de mecanismos de concorrência da linguagem Go para tornar a atualização das entidades mais dinâmica. O projeto explora esse tema principalmente em [src/sistema/sistema_de_movimento.go](../../src/sistema/sistema_de_movimento.go) e [src/sistema/sistema_de_ia.go](../../src/sistema/sistema_de_ia.go).

#### Goroutines

As goroutines são utilizadas para processar diferentes entidades do jogo em paralelo. No sistema de movimento, cada entidade recebe sua própria rotina concorrente para executar o método de atualização. Isso permite que várias entidades sejam processadas ao mesmo tempo, reduzindo o tempo de atualização do frame e deixando o sistema mais preparado para cenários com muitos inimigos e objetos interagindo.

O fluxo típico é:

- o sistema coleta as entidades ativas;
- cria uma goroutine para cada uma;
- cada goroutine executa a lógica de atualização;
- o fluxo principal aguarda o término com WaitGroup.

#### WaitGroup

O projeto utiliza sync.WaitGroup para garantir que todas as goroutines concluam antes de seguir para a próxima etapa do frame. Esse mecanismo é essencial para evitar que o jogo avance com entidades ainda em processamento.

#### Channels

O sistema de IA usa channels para organizar a comunicação entre as decisões paralelas dos bots. A ideia é simples: cada bot produz uma decisão de ação e envia essa decisão para um canal, que depois é consumido pelo fluxo principal. Esse padrão deixa a lógica mais organizada e prepara o projeto para evoluir para uma IA mais complexa e estruturada.

#### RWMutex e sincronização de estado

A cena de jogo também utiliza um lock de leitura e escrita em [src/cenas/cenaJogo.go](../../src/cenas/cenaJogo.go) para proteger o mapa de entidades. Esse recurso é importante porque, em um ambiente concorrente, é necessário evitar que múltiplas rotinas leiam e escrevam o mesmo estado ao mesmo tempo.

Em termos práticos, a sincronização ajuda a:

- proteger o mapa de entidades contra acessos simultâneos;
- evitar inconsistências durante remoções e inserções;
- manter o estado do jogo estável durante a atualização de vários sistemas.

#### Por que isso importa no projeto

Esse uso de concorrência mostra que o projeto não apenas implementa um jogo funcional, mas também tenta explorar recursos reais de Go para construir uma base mais moderna, escalável e alinhada com boas práticas de programação.

---

## 9. Persistência de dados

O progresso do jogador é armazenado em [progresso.json](../../progresso.json). O carregamento e salvamento do estado são feitos em [src/nivel/progresso.go](../../src/nivel/progresso.go) e [src/nivel/carregar.go](../../src/nivel/carregar.go).

Isso permite que a partida continue em um estado já conhecido entre diferentes execuções.

---

## 10. Como executar o projeto

### Requisitos

- Go instalado
- ambiente compatível com a biblioteca Ebiten

### Execução local

No diretório do projeto, execute:

```bash
go run .
```

### Compilação

```bash
go build ./...
```

---

## 11. Controles principais

- Enter: iniciar a partida
- Esc: retornar ao menu ou sair da tela atual
- Setas ou WASD: mover o jogador
- P: pausar o jogo
- Ctrl + M: alternar a posição do minimapa
- Ctrl + O: mostrar ou ocultar o minimapa

---

## 12. Pontos fortes do projeto

- estrutura modular e com boa separação de responsabilidades;
- implementação funcional de um loop de jogo em Go;
- uso de interfaces e abstrações para facilitar evolução;
- presença de mecânicas de gameplay já integradas;
- potencial para expansão com mais fases, IA e melhorias gráficas.

---

## 13. Limitações observadas

Apesar do estado funcional, ainda existem pontos que podem ser melhorados:

- parte dos sistemas ainda pode ser refinada para reduzir acoplamento;
- a IA dos bots é relativamente simples;
- a colisão é funcional, mas ainda pode ser expandida;
- a experiência visual pode ser aprimorada com mais animações e feedback;
- ainda não há suíte de testes automatizados no repositório.

---

## 14. Próximos passos recomendados

Para evoluir o projeto de forma consistente, os próximos passos mais úteis seriam:

1. consolidar a IA dos bots com decisões mais inteligentes;
2. ampliar o combate e as interações entre entidades;
3. melhorar a resposta de colisão e a robustez do mapa;
4. adicionar animações e sons;
5. implementar testes automatizados;
6. expandir o conteúdo com novos níveis e novos desafios.

---

## 15. Referências internas

- [Documentação técnica](../tecnico/documentacao_jogo.md)
- [Paralelismo em Go](../tecnico/paralelismo_go.md)
- [Resumo de avaliação](../avaliacao/Resposta_aos_criterios_de_avaliacao.md)
- [README principal](../../README.md)

Esta wiki representa uma versão consolidada do projeto e pode ser usada como base para exportação para uma wiki oficial do GitHub.
