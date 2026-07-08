# 🗺️ Wiki do Projeto: Gopher Dungeon Arena

Bem-vindo à documentação técnica oficial do **Gopher Dungeon Arena**. Este espaço centraliza os detalhes arquiteturais, decisões de design, modelagem de dados e critérios de engenharia de software utilizados no desenvolvimento do jogo.

---

## 1. Visão Geral e Contexto Acadêmico
O **Gopher Dungeon Arena** é um jogo arcade de arena estilo *dungeon crawler* desenvolvido de forma individual por **Calebe** para a disciplina de **Linguagens de Programação** (UnB - 2026). 
O projeto foi concebido para demonstrar a viabilidade prática, a performance e a elegância da linguagem **Go (Golang)** no domínio de simulações interativas e desenvolvimento de jogos 2D em tempo real.

---

## 2. Ciclo de Vida e Fluxo do Game Loop

O motor do jogo utiliza o ecossistema da biblioteca **Ebitengine (Ebiten)**, operando a uma taxa estável de 60 atualizações por segundo (FPS/UPS). O controle central do ciclo de vida do software é governado por três métodos fundamentais na estrutura principal:

*   **`Update() error`**: Executa o processamento assíncrono da lógica de estados e inputs, delegando a execução para a cena atualmente ativa (`CenaCorrente.Update()`).
*   **`Draw(tela *ebiten.Image)`**: Realiza a varredura gráfica, limpando o buffer de vídeo com a cor base branca e renderizando as matrizes e formas na tela (`CenaCorrente.Draw()`).
*   **`Layout(l, a int) (int, int)`**: Modula dinamicamente a proporção interna da janela de exibição em relação ao tamanho da tela.

### Máquina de Estados de Cenas (FSM)
O jogo implementa uma interface contratual `ICena` para alternar o fluxo de execução entre telas de forma isolada:
Use o código com cuidado.[Menu Iniciar] ──(ENTER)──> [Cena Jogo] ──(P)──> [Menu Pause]│                         │                    │(ESC)                     (morte)              (ESC)▼                         ▼                    ▼Encerra Jogo             Tela Derrota          Menu Iniciar
---

## 3. Padrão Arquitetural ECS (Entity Component System)

Para contornar a ausência de herança tradicional de classes em Go, o projeto adota uma arquitetura **ECS Leve**. Isso desacopla completamente os dados brutos de suas respectivas lógicas operacionais.

*   **Entidade**: Um contêiner lógico abstrato identificado por um `EntidadeID` único, que carrega um dicionário dinâmico do tipo `map[string]interface{}` para armazenar seus componentes.
*   **Componente**: Estruturas de dados puras (sem comportamento intrínseco), como `Vida`, `Nivel`, `Pontuacao` e `CORPO` (geometria de colisão).
*   **Sistema**: Algoritmos puros que realizam varreduras na coleção de entidades e aplicam transformações de estado caso os componentes necessários existam.

### Contrato Base da Entidade (`src/ecs/ecs.go`)
```go
type Entidade interface {
    GetID() EntidadeID
    GetTipo() string
    GetComponente(id string) interface{}
    ExisteComponente(id string) bool
    Atualizar()
    Desenhar(screen *ebiten.Image)
    DesenharMapa(screen *ebiten.Image, mapaX, mapaY float64)
}
```

---

## 4. Ecossistema de Sistemas de Atualização

Os subsistemas implementam o contrato `ISistemaAtualizar` e processam a lógica de simulação sequencialmente a cada frame:

1.  **`SistemaInput`**: Captura barramentos de teclado físicos, mapeando as teclas `WASD/Setas` para atualizar os vetores de velocidade do jogador.
2.  **`SistemaSpawn`**: Responsável pelo controle de densidade do mapa. Instancia o labirinto e as bordas de colisão na inicialização. Executa rotinas automáticas de *respawn* a cada 30 segundos ou sob demanda (tecla `B` para debug), garantindo posições válidas livres de sobreposição espacial.
3.  **`SistemaMovimento`**: Varre as entidades e aciona os comportamentos de deslocamento no espaço do mundo de forma atômica.
4.  **`SistemaColisao`**: Verifica a interseção de caixas delimitadoras envolventes (**AABB - Axis-Aligned Bounding Boxes**). Previne o tunelamento separando as verificações independentemente no eixo X e no eixo Y.
5.  **`SistemaDesenho`**: Projeta o mundo escalado (2x o tamanho da janela), controla a **Câmera Dinâmica** focada na posição central do jogador e renderiza o HUD junto ao **Mini Mapa** no canto da tela.
6.  **`SistemaDebug`**: Utiliza recursos de reflexão em tempo de execução (`reflect`) para inspecionar e despejar a árvore de memória interna das entidades no console através das teclas `F1` e `F2`.

---

## 5. Inteligência Artificial e Padrões de Movimentação

Os Bots inimigos recebem polimorfismo comportamental através da interface `Movimentador`. O motor encapsula **9 algoritmos matemáticos distintos de movimentação**, mapeados e identificados por cores exclusivas na arena:

| ID | Movimentador | Comportamento Lógico |
| :--- | :--- | :--- |
| 1 | **Simples** | Vetores aleatórios que mudam ao impactar obstáculos ou de forma probabilística. |
| 2 | **Vertical** | Movimento alternado estrito no eixo Y (cima/baixo). |
| 3 | **Vertical Constante** | Deslocamento unidirecional contínuo no eixo Y. |
| 4 | **Horizontal** | Movimento alternado estrito no eixo X (esquerda/direita). |
| 5 | **Horizontal Constante** | Deslocamento unidirecional contínuo no eixo X. |
| 6 | **Diagonal** | Reflexões angulares perfeitas em eixos de 45°. |
| 7 | **Lógico Linha** | Tenta manter linhas retas contínuas mudando direções sob janelas lógicas. |
| 8 | **Lógico Diagonal** | Segue equações matemáticas discretas para simular ziguezagues diagonais. |
| 9 | **Lógico Duplo** | Executa duas direções paralelas concorrentes, tornando o Bot mais rápido e evasivo. |

---

## 6. Mecânicas Avançadas do Núcleo (Core)

### Sistema de Combate e Dano
Quando o `SistemaColisao` detecta interseção direta entre a entidade `JOGADOR` e um `BOT`, a rotina de combate calcula o decaimento explícito de HP com base no multiplicador de nível:
*   O jogador perde `COMBATE_BOT_RIT` de sangue.
*   O bot perde `COMBATE_JOGADOR_RIT` de sangue.
*   Se o componente `Vida.Sangue` atingir zero, a propriedade `Status` é alterada para `false` e o `SistemaEntidades` remove o elemento da memória.

### Barramento de Teletransporte por Portais
A física de teletransporte é gerida por comunicação de dados em componentes de acoplamento:
1.  O Bot intercepta um volume de `PortalEntrada`.
2.  A entidade do Bot tem seus movimentos bloqueados e é transferida para o componente `ENVIANDO_TELETRANSPORTE`, iniciando uma contagem regressiva gráfica com quinas rotativas.
3.  Após atingir o limite de frames, o Bot é extraído pelo componente `RECEBENDO_TELETRANSPORTE` vinculado ao `PortalSaida`, sendo restabelecido fisicamente no novo ponto geométrico da dungeon.

### Condições de Vitória e Derrota
*   **Vitória**: O jogador deve coletar o requisito de alimentos espalhados pela arena (incrementando o componente `Pontuacao`) e alcançar com sucesso a área do Portal de Saída.
*   **Derrota**: Ocorre quando o componente `Vida.Quantidade` ou `Sangue` do jogador chega a zero, disparando a transição imediata para a cena informativa de fim de jogo.

---

## 7. Mapeamento de Arquivos do Projeto

├── main.go                       # Ponto de entrada, inicialização do Ebiten e Game Loop└── src/├── game.go                   # Controlador principal da FSM de cenas do jogo├── assets/                   # Carregamento de fontes embutidas em binário via //go:embed├── cenas/                    # Implementações isoladas (Menu Iniciar, Jogo, Pausa)├── config/                   # Configurações de dimensão de tela, proporção do mapa e geradores├── ecs/                      # Contratos estruturais de Entidades, Câmera e Mini Mapa├── enum/                     # Centralização de chaves de Componentes, Cores e Tipos├── sistema/                  # Motores algorítmicos (Input, Movimento, Colisão, Spawn, Desenho, Debug)├── entidades/│   ├── personagens/          # Regras comportamentais estruturadas do Jogador e dos Bots│   ├── objeto/               # Entidades estáticas e dinâmicas (Paredes, PortalEntrada, PortalSaida)│   ├── funcionalidades/       # Escopos de execução de métodos matemáticos de Combate e Teletransporte│   └── geometria/            # Primitivas geométricas bidimensionais (Retângulos, Pontos e Vetores)└── utils/                    # Banco de constantes globais de velocidade, ticks e tamanhos físicos
---

## 8. Guia de Compilação e Execução

### Pré-requisitos
*   **Go SDK** instalado (versão 1.16 ou superior).
*   Drivers gráficos atualizados compatíveis com OpenGL/DirectX (Requisito do Ebitengine).

### Comandos de Terminal

```bash
# 1. Instalar as dependências do ecossistema gráfico
go get ://github.com

# 2. Executar o projeto diretamente em tempo de desenvolvimento
go run main.go

# 3. Compilar um executável binário nativo estático
# No Linux/macOS:
go build -o bin/Gopher_Dungeon_Arena
# No Windows (PowerShell):
go build -o bin/Gopher_Dungeon_Arena.exe
```

---
*Última atualização de especificação da engenharia: 2026-07-06*