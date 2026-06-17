# Documentação Completa - Gopher Dungeon Arena

## 📋 Índice
1. [Visão Geral](#1-visão-geral)
2. [Objetivos do Projeto](#2-objetivos-do-projeto)
3. [Arquitetura do Sistema](#3-arquitetura-do-sistema)
4. [Estrutura de Pastas](#4-estrutura-de-pastas)
5. [Sistema de Jogo (Game Loop)](#5-sistema-de-jogo-game-loop)
6. [Sistema de Cenas](#6-sistema-de-cenas)
7. [ECS (Entity Component System)](#7-ecs-entity-component-system)
8. [Entidades do Jogo](#8-entidades-do-jogo)
9. [Componentes](#9-componentes)
10. [Sistemas de Atualização](#10-sistemas-de-atualização)
11. [Movimentadores e IA](#11-movimentadores-e-ia)
12. [Mecânicas de Jogo](#12-mecânicas-de-jogo)
13. [Fluxo de Jogo](#13-fluxo-de-jogo)
14. [Configuração e Constantes](#14-configuração-e-constantes)
15. [Como Compilar e Executar](#15-como-compilar-e-executar)

---

## 1. Visão Geral

**Gopher Dungeon Arena** é um jogo 2D de arena construído em **Go** utilizando a biblioteca **Ebiten** para renderização gráfica. O projeto é um trabalho acadêmico da UnB (Universidade de Brasília) para a disciplina de **Linguagens de Programação** (8º semestre).

### Características principais:
- **Motor de jogo baseado em ECS**: Entidades, componentes e sistemas para gerenciamento eficiente de objetos
- **Sistema de câmera dinâmica**: Câmera que segue o jogador
- **Mini mapa**: Visualização de mundo em tempo real
- **Múltiplas cenas**: Menu inicial, jogo, pausa, progresso, vitória e derrota
- **Inimigos com IA**: 9 tipos diferentes de movimentadores para bots
- **Sistema de combate**: Jogador pode atacar e ser atacado por bots
- **Colisão e física**: Sistema de detecção e resposta de colisões
- **Spawn dinâmico**: Bots aparecem aleatoriamente durante o jogo

---

## 2. Objetivos do Projeto

O projeto visa demonstrar:
- **Uso de Go para desenvolvimento de jogos**: Explorar potenciais de Go em projetos interativos
- **Concorrência**: Gerenciamento eficiente de múltiplas entidades
- **Tipagem segura**: Sistema de tipos forte e bem organizado
- **Arquitetura escalável**: Padrão ECS para facilitar expansão e manutenção
- **Performance**: Renderização suave mesmo com múltiplas entidades

---

## 3. Arquitetura do Sistema

### Estrutura em Camadas:

```
┌─────────────────────────────────┐
│     Ebiten (Rendering)          │
└─────────────────────────────────┘
         ↑          ↓
┌─────────────────────────────────┐
│   Game Loop (Update / Draw)      │
└─────────────────────────────────┘
         ↑          ↓
┌─────────────────────────────────┐
│    Cenas (Scene Management)      │
└─────────────────────────────────┘
         ↑          ↓
┌─────────────────────────────────┐
│  Sistemas (Update + Rendering)   │
├─────────────────────────────────┤
│ • SistemaInput                   │
│ • SistemaIA                      │
│ • SistemaSpawn                   │
│ • SistemaMovimento               │
│ • SistemaEntidades               │
│ • SistemaColisao                 │
│ • SistemaDesenho                 │
│ • SistemaDebug                   │
└─────────────────────────────────┘
         ↑          ↓
┌─────────────────────────────────┐
│    ECS (Entities + Components)   │
└─────────────────────────────────┘
```

---

## 4. Estrutura de Pastas

```
LP_Projeto_Jogo_Gopher_Dungeon_Arena/
├── main.go                          # Ponto de entrada
├── go.mod                           # Definição do módulo Go
├── build.ps1                        # Script de build para Windows
├── README.md                        # Instruções gerais
├── SUMARIO.md                       # Sumário do projeto
├── instrucoes.text                  # Instruções do jogo
├── ico.syso                         # Ícone do jogo
├── docs/
│   └── documentacao_jogo.md         # Esta documentação
├── src/
│   ├── game.go                      # Classe principal do jogo
│   ├── assets/                      # Recursos (fontes, imagens)
│   │   ├── assets.go
│   │   ├── fonte.go
│   │   ├── fontes/                  # Fontes TTF
│   │   └── imagens/
│   ├── cenas/                       # Gerenciamento de cenas
│   │   ├── cenaMenuIniciar.go       # Menu inicial
│   │   ├── cenaMenuPause.go         # Menu de pausa
│   │   ├── cenaJogo.go              # Cena principal
│   │   ├── cenaProgresso.go         # Tela de progresso
│   │   └── informativos/
│   │       ├── ganhou.go            # Tela de vitória
│   │       └── perdeu.go            # Tela de derrota
│   ├── componentes/                 # Definição de componentes
│   │   ├── componentes.go           # Enum de componentes
│   │   └── movimentacao/            # Tipos de movimentadores
│   │       ├── movimentadorSimples.go
│   │       ├── movimentadorVertical.go
│   │       ├── movimentadorHorizontal.go
│   │       ├── movimentadorDiagonal.go
│   │       ├── movimentadorLogicoLinha.go
│   │       ├── movimentadorLogicoDiagonal.go
│   │       ├── movimentadorLogicoDuplo.go
│   │       └── [outros movimentadores]
│   ├── config/                      # Configurações globais
│   │   └── configuracoes.go
│   ├── ecs/                         # Engine ECS
│   │   ├── ecs.go                   # Interface Entidade
│   │   ├── camera.go                # Sistema de câmera
│   │   ├── miniMapa.go              # Mini mapa
│   │   ├── barJogador.go            # Barra de vida do jogador
│   │   └── respostaColisao.go
│   ├── entidades/                   # Implementações de entidades
│   │   ├── personagens/
│   │   │   ├── jogador.go           # Classe Jogador
│   │   │   └── bot.go               # Classe Bot
│   │   ├── objeto/
│   │   │   ├── comida.go            # Item de comida
│   │   │   ├── parede.go            # Obstáculos
│   │   │   ├── portalEntrada.go     # Portal de entrada
│   │   │   ├── portalSaida.go       # Portal de saída
│   │   │   ├── saida.go             # Zona de saída
│   │   │   └── [sprites portais]
│   │   ├── funcionalidades/         # Lógica compartilhada
│   │   │   ├── combate.go           # Sistema de combate
│   │   │   ├── comer.go             # Comer alimento
│   │   │   ├── teletransporte.go    # Mecânica de teletransporte
│   │   │   └── concluirPartida.go   # Conclusão de partida
│   │   ├── geometria/               # Primitivas geométricas
│   │   │   ├── ponto.go
│   │   │   └── retangulo.go
│   │   └── outros/
│   │       └── time.go              # Classe Time (equipes)
│   ├── enum/                        # Enumerações
│   │   ├── cores/
│   │   │   └── cores.go
│   │   └── entidades/
│   │       └── entidade_tipo.go
│   ├── interfaces/                  # Contratos (interfaces)
│   │   ├── i_cena.go
│   │   ├── i_cenajogo.go
│   │   ├── i_game.go
│   │   ├── i_movimentador.go
│   │   ├── i_sistema_atualizar.go
│   │   ├── i_sistema_colisao.go
│   │   ├── i_sistema_desenhar.go
│   │   └── i_sistema_ia.go
│   ├── sistema/                     # Sistemas ECS
│   │   ├── sistema_de_colisao.go
│   │   ├── sistema_de_debug.go
│   │   ├── sistema_de_desenho.go
│   │   ├── sistema_de_entidades.go
│   │   ├── sistema_de_ia.go
│   │   ├── sistema_de_input.go
│   │   ├── sistema_de_movimento.go
│   │   └── sistema_de_spawn.go
│   └── utils/                       # Utilitários
│       ├── constantes.go
│       ├── desenho.go
│       └── sprits/
│           └── portais.go
└── tmp/                             # Arquivos temporários de build
```

---

## 5. Sistema de Jogo (Game Loop)

### `main.go` - Ponto de Entrada

```go
func main() {
    game := src.NovoGame()
    ebiten.SetWindowSize(config.JANELA_LARGURA, config.JANELA_ALTURA)
    ebiten.SetWindowTitle(config.NOME_JOGO)
    ebiten.MaximizeWindow()
    imagens.AplicarIconeJanela()
    
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}
```

**Responsabilidades:**
- Inicializa o jogo
- Define dimensões da janela (1280x720)
- Define o título ("Gopher_Dungeon_Arena")
- Maximiza a janela
- Aplica ícone customizado
- Inicia o loop de jogo do Ebiten

### `src/game.go` - Classe Principal

```go
type Game struct {
    CenaCorrente interfaces.ICena      // Cena atual renderizada
    CenaJogo     interfaces.ICenaJogo  // Referência à cena de jogo
}

// Métodos principais:
func (g *Game) Update() error       // Atualiza lógica de jogo
func (g *Game) Draw(tela *ebiten.Image)  // Renderiza na tela
func (g *Game) Layout(l, a int) (int, int)  // Define layout

// Métodos de transição:
func (g *Game) IniciarJogo()                // Menu → Jogo
func (g *Game) Pausar()                     // Jogo → Pausa
func (g *Game) Voltar()                     // Pausa → Jogo
func (g *Game) MudarTelaMenuIniciar()       // Qualquer → Menu
func (g *Game) MudarTelaProgresso()         // Qualquer → Progresso
func (g *Game) ReiniciarMudarTelaMenuIniciar()  // Reset completo
func (g *Game) Sair()                       // Encerra jogo
```

**Fluxo do Game Loop:**

1. **Inicialização**: `NovoGame()` cria as cenas iniciais
2. **Loop Ebiten** (60 FPS por padrão):
   - Chama `Update()` → `CenaCorrente.Update()`
   - Chama `Draw(screen)` → `CenaCorrente.Draw(screen)`
3. **Transições**: Métodos como `IniciarJogo()` trocam `CenaCorrente`

---

## 6. Sistema de Cenas

As cenas implementam a interface `ICena` e controlam o fluxo de tela.

### Cenas Disponíveis:

#### `CenaMenuIniciar`
- **Entrada**: Logo e texto do menu
- **Controles**:
  - `ENTER` → Inicia jogo
  - `ESC` → Sai do jogo
- **Transição**: Para `CenaJogo`
- **Renderização**: Usa `assets.FonteCache` para texto

#### `CenaJogo`
- **Cena principal** do jogo
- **Elementos**:
  - Mundo de 2560x1440 pixels (2x da janela)
  - Câmera dinâmica seguindo o jogador
  - Mini mapa
  - Entidades: Jogador, bots, paredes, portais, comida
- **Controles**:
  - Setas ou WASD → Movimento
  - `P` → Pausa
  - `B` → Spawn bots (debug)
- **Update**:
  - `SistemaInput`: Processa entrada
  - `SistemaIA`: IA dos bots
  - `SistemaSpawn`: Spawn de novos bots
  - `SistemaMovimento`: Atualiza posições
  - `SistemaEntidades`: Atualiza entidades
  - `SistemaDebug`: Debug
  - Sistema de colisão integrado
  - Remove entidades mortas
- **Transição**: Para `CenaMenuPause` (pausar) ou `InformativoGanhou`/`InformativoPerdeu`

#### `CenaMenuPause`
- **Entrada**: Quando `P` é pressionado
- **Controles**:
  - `SPACE` → Volta ao jogo
  - `ESC` → Sai do jogo
- **Renderização**: Instrções de controle na tela

#### `CenaProgresso`
- Tela de progresso/estatísticas

#### `InformativoGanhou` / `InformativoPerdeu`
- Telas de conclusão de partida
- Exibe resultado final

---

## 7. ECS (Entity Component System)

### Conceito

O ECS é um padrão arquitetural que separa dados (componentes) da lógica (sistemas):

- **Entidade**: Container de componentes, identificada por `EntidadeID`
- **Componente**: Dados puros (sem comportamento)
- **Sistema**: Lógica que opera sobre entidades com certas componentes

### Interface Entidade

```go
type Entidade interface {
    GetID() EntidadeID                              // ID único
    GetTipo() string                                // Tipo (BOT, JOGADOR, etc)
    GetComponente(id string) interface{}            // Obter componente
    ExisteComponente(id string) bool               // Verificar existência
    Atualizar()                                     // Update da entidade
    Desenhar(screen *ebiten.Image)                 // Renderizar
    DesenharMapa(screen *ebiten.Image, 
                 mapaX, mapaY float64)              // Renderizar no mini mapa
}
```

### Gerenciamento de Entidades

```go
// Em CenaJogo:
entidades map[ecs.EntidadeID]ecs.Entidade

// Métodos:
func (cj *CenaJogo) CriarEntidade() ecs.EntidadeID  // Cria novo ID
func (cj *CenaJogo) SetEntidade(id EntidadeID, e Entidade)  // Adiciona
func (cj *CenaJogo) RemoverEntidade(id EntidadeID)  // Remove
func (cj *CenaJogo) GetEntidades() map[EntidadeID]Entidade  // Lista todas
```

---

## 8. Entidades do Jogo

### 8.1 Jogador

```go
type Jogador struct {
    entidadeID ecs.EntidadeID
    nome       string
    cor        color.Color
    Componentes map[string]interface{}
}
```

**Componentes:**
- `CORPO`: Retângulo (posição e tamanho)
- `VIDA`: Vida (3 HP), sangue (100), status
- `NIVEL`: Nível (começa em 1)
- `PONTUACAO`: Coletas necessárias, itens coletados

**Métodos:**
- `GetID()`, `GetTipo()` → "JOGADOR"
- `ObterVida()`, `ObterCorpo()`, `ObterNivel()`
- `EstaVivo()`, `CorrigeSangue()`, `Renasce()`
- `SetPosicao(x, y)`, `SetNivel(n)`
- `Atualizar()`, `Desenhar()`

**Tamanho**: 10x10 pixels (mundo), redimensionado para tela

### 8.2 Bot (Inimigo)

```go
type Bot struct {
    entidadeID ecs.EntidadeID
    Id         int64
    Componentes map[string]interface{}
}
```

**Componentes:**
- `CORPO`: Retângulo
- `VIDA`: 1 HP, sangue (100), status
- `NIVEL`: Nível aleatório
- `MOVIMENTO`: Tipo de movimentação
- `ATIVIDADE`: Ação atual
- `SUB_TIPO`: Identificação do tipo

**Características:**
- Spawn aleatório a cada 30 segundos
- 9 tipos diferentes de movimento
- Podem ser destruídos em combate
- Dropam pontuação ao morrer

### 8.3 Comida

```go
type Comida struct {
    entidadeID ecs.EntidadeID
    estrutura  *geometria.Retangulo
    Componentes map[string]interface{}
    ciclos     int
}
```

**Componentes:**
- `CORPO`: Retângulo
- `ENERGIA`: Valor (100), status

**Funcionalidade:**
- Item coletável pelo jogador
- Aumenta pontuação do jogador
- Animação de ciclo visual

### 8.4 Parede

```go
type Parede struct {
    entidadeID ecs.EntidadeID
    corpo      *geometria.Retangulo
}
```

**Funcionalidade:**
- Obstáculo para colisão
- Forma labirinto no mapa
- Colidem com jogador e bots

### 8.5 Portais

#### Portal Entrada (`PortalEntrada`)
```go
type PortalEntrada struct {
    entidadeID ecs.EntidadeID
    corpo      *geometria.Retangulo
    anguloRotacao float64
}
```

**Componentes:**
- `CORPO`: Retângulo
- `ENVIANDO_TELETRANSPORTE`: Dados de teletransporte

**Funcionalidade:**
- Quando entidade entra, inicia teleporte
- Efeito visual: rotação animada
- Teletransporta para Portal Saída

#### Portal Saída (`PortalSaida`)
```go
type PortalSaida struct {
    entidadeID ecs.EntidadeID
    corpo      *geometria.Retangulo
    offsetBarras float64
}
```

**Componentes:**
- `CORPO`: Retângulo
- `RECEBENDO_TELETRANSPORTE`: Dados de recepção

**Funcionalidade:**
- Recebe entidades teleportadas
- Efeito visual: barras animadas
- Restaura entidade em nova posição

### 8.6 Time (Equipe)

```go
type Time struct {
    entidadeID ecs.EntidadeID
    nome       string
    cor        color.Color
    Membros    []ecs.Entidade
}
```

**Funcionalidade:**
- Agrupa entidades (jogador + possíveis aliados)
- Pode ser expandido para multiplayer

---

## 9. Componentes

### Enum de Componentes

```go
const (
    CORPO                    // Posição e tamanho
    SUB_TIPO                 // Subtipo identificador
    VIDA                     // Health/sangue
    NIVEL                    // Level e progressão
    ENVIANDO_TELETRANSPORTE  // Portal de entrada
    ATIVIDADE                // Ação atual
    RECEBENDO_TELETRANSPORTE // Portal de saída
    MOVIMENTO                // Tipo de movimento
    ENERGIA                  // Energia/comida
    PONTUACAO                // Pontos coletados
)
```

### Estruturas de Dados

#### `Vida`
```go
type Vida struct {
    TipoOrganismo string  // "JOGADOR", "BOT"
    Quantidade    int     // Vidas restantes
    Status        bool    // Vivo?
    Sangue        int     // HP (0-100)
}

func (v *Vida) EstaVivo() bool
func (v *Vida) PerdeSangue(quantidade, nivel int)
func (v *Vida) CorrigeSangue(nivel int)
```

#### `Nivel`
```go
type Nivel struct {
    Valor      int  // Level atual
    Progressao int  // Progresso para próximo nível
}
```

#### `Pontuacao`
```go
type Pontuacao struct {
    Coletado       int   // Itens coletados
    Requisito      int   // Necessário para vencer
    EntreiNaSaida  bool  // Entrou no portal de saída?
}
```

#### `Movimento`
```go
type Movimento struct {
    Tipo interfaces.Movimentador  // Tipo de movimento
    Cor  color.Color              // Cor para renderização
}
```

#### `Atividade`
```go
type Atividade struct {
    Acao int  // AIVIDADE_MOVIMENTO ou AIVIDADE_TELETRANSPORTE
}
```

---

## 10. Sistemas de Atualização

Os sistemas implementam `ISistemaAtualizar` e processam lógica a cada frame.

### 10.1 SistemaInput

```go
type SistemaInput struct{}
func (s *SistemaInput) Atualizar(cj interfaces.ICenaJogo)
```

**Funcionalidade:**
- Processa entrada do teclado
- Controla movimento do jogador
- Ativa debug (pressionando teclas específicas)

### 10.2 SistemaIA

```go
type SistemaIA struct{}
func (s *SistemaIA) Atualizar(cj interfaces.ICenaJogo)
```

**Funcionalidade:**
- [Atualmente vazio - pode ser expandido para IA avançada]

### 10.3 SistemaSpawn

```go
type SistemaSpawn struct {
    framesGereacao int
}
func (s *SistemaSpawn) Atualizar(cj interfaces.ICenaJogo)
```

**Funcionalidade:**
- Spawn inicial: 9 tipos de bots (um de cada)
- Spawn periódico: A cada 30 segundos (~1860 frames)
- Spawn aleatório quando `B` pressionado (debug)
- Garante posições válidas sem colisão

**Métodos:**
- `SpawnJogadores()`: Cria jogador
- `SpawnBotDeCadaTipo()`: Cria 9 bots (um movimentador cada)
- `SpawnarBotAleatorio()`: Spawn dinâmico
- `SpawnParedesAoRedor()`: Cria bordas do mapa
- `SpawnLabirinto()`: Cria obstáculos internos
- `SpawnarPortais()`: Cria 2 portais de teletransporte

### 10.4 SistemaMovimento

```go
type SistemaMovimento struct{}
func (s *SistemaMovimento) Atualizar(cj interfaces.ICenaJogo)
```

**Funcionalidade:**
- Chama `Atualizar()` de todas as entidades
- Permite que movimentadores façam seu trabalho

### 10.5 SistemaEntidades

```go
type SistemaEntidades struct{}
func (s *SistemaEntidades) Atualizar(cj interfaces.ICenaJogo)
```

**Funcionalidade:**
- Atualização genérica de entidades
- Gerenciamento de ciclo de vida

### 10.6 SistemaDebug

```go
type SistemaDebug struct{}
func (s *SistemaDebug) Atualizar(cj interfaces.ICenaJogo)
```

**Funcionalidade:**
- Imprime informações de debug no console
- Pode ser expandido para visualizações

### 10.7 SistemaDesenho

```go
type SistemaDesenhar struct{}
func (s *SistemaDesenhar) Desenhar(cj interfaces.ICenaJogo, tela *ebiten.Image)
```

**Funcionalidade:**
- Limpa tela com cor branca
- Desenha borda do mundo
- Renderiza todas as entidades
- Renderiza mini mapa
- Verifica condições de vitória/derrota

---

## 11. Movimentadores e IA

### Interface Movimentador

```go
type Movimentador interface {
    Mover(entidade ecs.Entidade,
          sistemaColisao ISistemaColisao,
          mundo *geometria.Retangulo,
          bot HabilidadeMovimentacao,
          r *rand.Rand)
    GetTipo() string
    GetCor() color.Color
}
```

### 9 Tipos de Movimentadores

#### 1. **MovimentadorSimples**
- Movimento aleatório em direções (cima, baixo, esquerda, direita)
- Muda de direção quando atinge obstáculo ou aleatoriamente

#### 2. **MovimentadorVertical**
- Movimento alternado entre cima e baixo
- Mantém X fixo, varia Y

#### 3. **MovimentadorVerticalConstante**
- Movimento contínuo apenas vertical
- Não muda de direção

#### 4. **MovimentadorHorizontal**
- Movimento alternado entre esquerda e direita
- Mantém Y fixo, varia X

#### 5. **MovimentadorHorizontalConstante**
- Movimento contínuo apenas horizontal
- Não muda de direção

#### 6. **MovimentadorDiagonal**
- Movimento em ângulos de 45°
- Combina movimento X e Y

#### 7. **MovimentadorLogicoLinha**
- Tenta mover em linhas retas
- Inclui mudanças direcionais

#### 8. **MovimentadorLogicoDiagonal**
- Movimento lógico diagonal
- Segue padrões matemáticos

#### 9. **MovimentadorLogicoDuplo**
- Movimento duplo (2 direções simultâneas)
- Mais rápido e complexo

### Características Comuns

Todos usam:
- **Colisão**: Verificam `SistemaColisao.VaiColidir()`
- **Aleatoriedade**: Usam `rand.Rand` para variabilidade
- **Limite de mundo**: Respeitam bordas do mapa
- **Cor identificadora**: Cada tipo tem cor diferente

---

## 12. Mecânicas de Jogo

### 12.1 Sistema de Combate

```go
func CombateJogadorBot(jogador ecs.Entidade, bot ecs.Entidade)
func ReduzSangue(entidade ecs.Entidade, rit int)
```

**Mecânica:**
- Quando jogador e bot colidem:
  - Jogador perde: `COMBATE_BOT_RIT` sangue
  - Bot perde: `COMBATE_JOGADOR_RIT` sangue
- Sangue reduz com base no nível da entidade
- Quando sangue ≤ 0, entidade morre (Status = false)

**Valores (em utils/constantes.go):**
- `COMBATE_BOT_RIT`: Dano por combate do bot
- `COMBATE_JOGADOR_RIT`: Dano por combate do jogador

### 12.2 Sistema de Colisão

```go
type SistemaColisao struct {
    cenaJogo interfaces.ICenaJogo
}

func (s *SistemaColisao) VaiColidir(origem string, 
                                      origemEntidade ecs.Entidade,
                                      meuCorpoAtual *geometria.Retangulo,
                                      proximoCorpo *geometria.Retangulo) 
    *ecs.RespostaColisao
```

**Tipos que colidem:**
- PAREDE
- JOGADOR
- BOT
- PORTAL_ENTRADA
- PORTAL_SAIDA
- COMIDA
- SAIDA

**Resposta de colisão:**
- Impede movimento se houver colisão
- Retorna informação sobre tipo de colisão
- Previne auto-colisão verificando posição exata

### 12.3 Sistema de Coleta

```go
func Comer(entidade ecs.Entidade)
```

**Mecânica:**
- Quando jogador colide com COMIDA:
  - Incrementa contador de `Pontuacao`
  - Remove item do mapa
  - Se atingiu requisito, pode entrar no portal de saída

### 12.4 Sistema de Teletransporte

```go
type EnviandoTeletransporte struct {
    TemBot         bool
    Bot            ecs.Entidade
    Contagem       int
    ConectadoSaida ecs.Entidade
}

type RecebendoTeletransporte struct {
    TemBot   bool
    Bot      ecs.Entidade
    Contagem int
}
```

**Mecânica:**
1. Entidade entra em Portal Entrada
2. Componente `ENVIANDO_TELETRANSPORTE` ativa
3. Efeito visual: rotação por alguns frames
4. Entidade aparece em Portal Saída
5. Componente `RECEBENDO_TELETRANSPORTE` desativa após animação

### 12.5 Vitória e Derrota

**Vitória:**
- Coletar todos os itens de comida (requisito = 3)
- Entrar no portal de saída
- Tela `InformativoGanhou` exibida

**Derrota:**
- Todos os jogadores morrem (Health = 0)
- Tela `InformativoPerdeu` exibida

---

## 13. Fluxo de Jogo

### Inicialização

```
main.go
  ↓
NovoGame()
  ├─ CenaMenuIniciar (Cena atual)
  └─ CenaJogo (Criada mas não ativa)
       ├─ NovoCenaJogo()
       │   ├─ Criar mundo (2560x1440)
       │   ├─ Criar câmera
       │   ├─ Criar mini mapa
       │   ├─ Instanciar 8 sistemas
       │   ├─ SpawnarPortais()
       │   ├─ SpawnParedesAoRedor()
       │   ├─ SpawnLabirinto()
       │   ├─ SpawnBotDeCadaTipo()
       │   └─ Spawn 2 comidas
```

### Game Loop (60 FPS)

```
Cada Frame (16.67ms):
  1. Game.Update()
     └─ CenaCorrente.Update()

  2. Game.Draw(screen)
     └─ CenaCorrente.Draw(screen)
```

### Enquanto em CenaJogo

```
CenaJogo.Update():
  1. SistemaInput.Atualizar()      → Entrada do usuário
  2. SistemaIA.Atualizar()          → [Vazio]
  3. SistemaSpawn.Atualizar()       → Novo bots a cada 30s
  4. SistemaMovimento.Atualizar()   → Move entidades
  5. SistemaEntidades.Atualizar()   → Update geral
  6. SistemaDebug.Atualizar()       → Debug
  7. SistemaColisao.Atualizar()     → Colisões
  8. RemoveEntidadesMortas()        → Limpeza

CenaJogo.Draw(screen):
  1. SistemaDesenho.Desenhar()
     ├─ Preencher tela (branco)
     ├─ Desenhar borda mundo
     ├─ Desenhar todas entidades
     ├─ Desenhar mini mapa
     └─ Verificar vitória/derrota
```

### Transições de Cena

```
Menu → ENTER → Jogo → P → Pausa → SPACE → Jogo
                  ↓                    ↓
                VITÓRIA            ESC → Menu
                DERROTA
```

---

## 14. Configuração e Constantes

### `src/config/configuracoes.go`

```go
// Dimensões da janela
const JANELA_LARGURA = 1280
const JANELA_ALTURA = 720

// Proporções do mundo (deve ser par)
const PROPORCAO_MUNDO = 2  // 2x = 2560x1440

// Dimensões do mini mapa
const PROPORCAO_MAPA = 8  // Proporcao_mundo * 4

// Posições do mini mapa (4 cantos opcionais)
const MM1_POS_X_MAPA = 50
const MM1_POS_Y_MAPA = MAPA_ALTURA / 3

const MM2_POS_X_MAPA = JANELA_LARGURA - (MAPA_LARGURA + MAPA_ALTURA/4)
const MM2_POS_Y_MAPA = MAPA_ALTURA / 3

// Nome do jogo
const NOME_JOGO = "Gopher_Dungeon_Arena"

// Funções auxiliares
func GeradorAleatorio() *rand.Rand
func XAleatorio(r *rand.Rand) float64
func YAleatorio(r *rand.Rand) float64
```

### `src/utils/constantes.go`

Contém tamanhos de entidades, valores de combate, etc:

```go
const JOGADOR_TAMANHO_MUNDO = 10
const BOT_TAMANHO_MUNDO = 10
const COMIDA_TAMANHO_MUNDO = 10
const PORTAL_ENTRADA_TAMANHO = 50
// ... etc
```

---

## 15. Como Compilar e Executar

### Pré-requisitos

- **Go 1.16+** instalado
- **Dependências do Ebiten**: SDL2 no Windows

### Instalação de Dependências

```bash
go get github.com/hajimehoshi/ebiten/v2
go get github.com/hajimehoshi/ebiten/v2/cmd/ebitenex@latest
```

### Compilação

#### Windows (PowerShell)

```powershell
# Via script fornecido
.\build.ps1

# Ou manualmente
go build -o bin/Gopher_Dungeon_Arena.exe
```

#### Linux/macOS

```bash
go build -o bin/Gopher_Dungeon_Arena
./bin/Gopher_Dungeon_Arena
```

### Execução

```bash
go run main.go
```

### Controles do Jogo

| Controle | Ação |
|----------|------|
| SETAS ou WASD | Mover |
| ENTER | Iniciar jogo |
| ESC | Sair/Menu |
| P | Pausar |
| SPACE | Despausar |
| B | Spawn bots (debug) |

---

## Resumo da Arquitetura

O projeto demonstra uma arquitetura robusta com:

✅ **ECS bem estruturado** para separação de dados e lógica  
✅ **Sistema de cenas** para gerenciamento de telas  
✅ **9 tipos de IA** diferentes para variedade  
✅ **Sistema de câmera dinâmica** com mini mapa  
✅ **Colisão precisa** com prevenção de tunelamento  
✅ **Componentes reutilizáveis** para fácil expansão  
✅ **Código bem organizado** em pacotes específicos  
✅ **Configuração centralizada** para fácil ajuste  

---

**Última atualização**: 2026-06-17

### `src/enum/componentes/componentes.go`
- Enumeração de componentes: `CORPO`, `SUB_TIPO`, `VIDA`, `NIVEL`, `ENVIANDO_TELETRANSPORTE`, `LIBERDADE`, `RECEBENDO_TELETRANSPORTE`.
- Define structs de componentes:
  - `SubTipo`
  - `Vida` (com `TipoOrganismo`, `Quantidade`, `Status`, `Sangue`)
  - `Nivel`
  - `EnviandoTeletransporte`
  - `Liberdade`
  - `RecebendoTeletransporte`
- Métodos de `Vida` controlam vida, renascimento e sangramento.

### Entidades principais

#### `src/entidades/personagens/jogador.go`
- Representa o jogador controlado pelo teclado.
- Usa `componentes.CORPO`, `componentes.VIDA`, `componentes.NIVEL`.
- Movimento pixel-a-pixel com verificação de colisão no eixo X e Y.
- Executa `Atira()` quando `SPACE` é pressionado (ainda vazio).
- Desenha jogador e barra de vida.
- Desenha jogador no mini mapa.

#### `src/entidades/personagens/bot.go`
- Representa bots inimigos.
- Usa `componentes.CORPO`, `componentes.SUB_TIPO`, `componentes.VIDA`, `componentes.NIVEL`, `componentes.LIBERDADE`.
- Recebe um `Movimentador` para definir o comportamento de movimento.
- Método `Mover()` delega em `movendo.Mover(...)`.
- Usa `TemBot` para impedir movimento durante teletransporte.
- Desenha bot e barra de vida.
- Desenha bot no mini mapa.

#### `src/entidades/objeto/parede.go`
- Representa paredes fixas no mapa.
- Só implementa corpo e desenho.
- Usada para colisão e geração de labirinto.

#### `src/entidades/objeto/portalEntrada.go` e `src/entidades/objeto/portalSaida.go`
- Portais de teletransporte com animação de quinas rotativas.
- Portal de entrada guarda um bot e começa contagem regressiva.
- Portal de saída recebe o bot após a contagem.
- Ambos usam componentes de teletransporte para compartilhar bots entre si.
- São usados no `SistemaColisao` para acionar teletransporte quando bots colidem com entradas.

#### `src/entidades/outros/time.go`
- Representa um time de jogadores.
- Agrupa jogadores e mantém referência de cor.
- Método `GetPosicaoTime()` retorna a posição de um jogador vivo.
- Essencialmente usado para organizar a câmera pelo jogador do time.

## 7. Sistemas do jogo

### `src/sistema/sistema_de_movimento.go`
- Chama `Atualizar()` em cada entidade.
- Isso dispara a lógica de movimento de jogadores/bots.

### `src/sistema/sistema_de_input.go`
- Placeholder para lógica de input global, atualmente vazio.

### `src/sistema/sistema_de_entidades.go`
- Placeholder para lógica de entidade geral, atualmente vazio.

### `src/sistema/sistema_de_ia.go`
- Placeholder para lógica de IA, atualmente vazio.

### `src/sistema/sistema_de_debug.go`
- Comandos de debug por tecla:
  - `F1`: lista principais entidades.
  - `F2`: lista entidades ordenadas por ID e campos públicos.
- Usa `reflect` para inspecionar entidades em tempo de execução.

### `src/sistema/sistema_de_spawn.go`
- Gera jogadores, bots, portais, paredes e labirinto.
- `SpawnJogadores()` cria jogador(s) e o time.
- `SpawnBotDeCadaTipo()` gera um bot para cada movimento disponível.
- `SpawnarBotAleatorio()` cria bots adicionais a cada 30 segundos e sob tecla `B`.
- `SpawnarPortais()` cria pares de portais de entrada/saída.
- `SpawnParedesAoRedor()` monta contorno e paredes externas.
- `SpawnLabirinto()` cria um labirinto interno com subdivisões.

### `src/sistema/sistema_de_desenho.go`
- Limpa a tela com `cores.BRANCO`.
- Desenha o mundo e cada entidade.
- Desenha mini mapa e entidades no mini mapa.
- Exibe mensagem de morte quando não há jogadores vivos.

### `src/sistema/sistema_de_colisao.go`
- Detecta colisões entre entidades.
- Controla interações:
  - jogador x bot => combate
  - bot x portal de entrada => teletransporte
- Ignora a auto-colisão da própria entidade.
- Método `ColideComTipo` permite verificar colisões com tipos específicos.
- Implementa helpers para barreiras, jogadores e bots.

## 8. Utilitários e constantes

### `src/utils/constantes.go`
- Define tamanhos e velocidades do mundo:
  - `JOGADOR_TAMANHO_MUNDO`, `BOT_TAMANHO_MUNDO`, `PAREDE_TAMANHO_MUNDO`, `PORTAL_ENTRADA_TAMANHO`.
  - Tamanhos reduzidos para mini mapa.
  - Velocidades e ritmos de combate.

### `src/utils/desenho.go`
- Não lido diretamente, mas o nome sugere funções de desenho auxiliares.
- Provavelmente usado para desenhar margens ou efeitos.

## 9. Recursos e assets

### `src/assets/` 
- `assets.go`: injeta a fonte embarcada com `//go:embed`.
- `fonte.go`: cria a fonte `GoTextFaceSource` usada em todas as cenas.
- `sprites/portal/teste.go`: arquivo experimental/inativo de portal com sprites em textura.
  - Esse arquivo parece ser uma variante de implementação de portais usando sprites em vez de formas geométricas.

## 10. Estruturas e organização técnica

### Arquitetura de cena
- `Game` mantém a cena atual e delega atualização/desenho.
- `CenaMenuIniciar`, `CenaMenuPause` e `CenaJogo` implementam `ICena`.
- `CenaJogo` expõe dados de entidade, câmera e sistemas para os subsistemas.

### ECS leve
- Entidades são structs que implementam `ecs.Entidade`.
- Cada entidade carrega um mapa `Componentes` com dados específicos.
- O sistema de componentes é dinâmico e usa `string` => `interface{}`.
- Componentes principais:
  - `CORPO`: retângulo de colisão.
  - `VIDA`, `NIVEL`, `SUB_TIPO`, `LIBERDADE`.
- Heurística de tipos utiliza `GetTipo()` e constantes de enum.

### Câmera e mini mapa
- `Camera` ajusta a visualização em torno do jogador.
- `OrganizarCameraPeloJogador()` centraliza a câmera e aplica limites.
- `MiniMapa` desenha uma vista reduzida do mundo e posiciona a área visível.

### Configurações de mundo
- `config.PROPORCAO_MUNDO` multiplica a área do mundo em relação à janela.
- `config.PROPORCAO_MAPA` reduz o coordenado para mini mapa.
- `MUNDO_LARGURA` / `MUNDO_ALTURA` criam um mundo maior do que a janela.
- `MAPA_LARGURA` / `MAPA_ALTURA` definem o mini mapa na tela.

## 11. Pontos de atenção e sugestões

- Vários sistemas (`SistemaInput`, `SistemaEntidades`, `SistemaIA`) estão definidos, mas não possuem lógica implementada. São ótimos pontos para estender o jogo.
- A colisão do jogador é feita pixel-a-pixel com `EstaNaMargemInterna`, o que é consistente, mas pode ser simplificada para desempenho.
- O labirinto e os portais já fornecem um bom ambiente, mas a lógica de spawn poderia usar componentes de tipo em vez de strings diretas em alguns pontos.
- `src/assets/sprits/portal/teste.go` parece ser um protótipo não usado pelo jogo atual. Ele pode ser removido ou movido para uma branch de experimentos.
- O método `Atira()` do jogador ainda não implementa um projétil; essa é uma funcionalidade natural para expansão.

## 12. Resumo por arquivo

| Arquivo | Função principal |
|---|---|
| `main.go` | Inicializa `ebiten` e roda o jogo. |
| `src/game.go` | Gerencia cena atual e transições. |
| `src/config/configuracoes.go` | Constantes de janela, mundo e geradores aleatórios. |
| `src/interfaces/*.go` | Contratos para cenas, jogo, ECS e sistemas. |
| `src/cenas/cenaMenuIniciar.go` | Menu inicial com `ENTER` e `ESC`. |
| `src/cenas/cenaMenuPause.go` | Menu de pausa com `SPACE` e `ESC`. |
| `src/cenas/cenaJogo.go` | Cena do jogo, inicializa mundo, entidades e sistemas. |
| `src/ecs/ecs.go` | Interface base de entidade. |
| `src/ecs/camera.go` | Câmera que centraliza no jogador. |
| `src/ecs/miniMapa.go` | Mini mapa e projeção da câmera. |
| `src/sistema/*.go` | Lógica de atualização, desenho, spawn e colisão. |
| `src/entidades/personagens/jogador.go` | Lógica do jogador e controle do teclado. |
| `src/entidades/personagens/bot.go` | Bots com movimento e vida. |
| `src/entidades/objeto/parede.go` | Obstáculo de parede. |
| `src/entidades/objeto/portalEntrada.go` | Portal de entrada que guarda bot. |
| `src/entidades/objeto/portalSaida.go` | Portal de saída que libera bot. |
| `src/entidades/outros/time.go` | Agrupamento de jogadores em time. |
| `src/entidades/funcionalidades/combate.go` | Lógica de combate entre jogador e bot. |
| `src/entidades/funcionalidades/teletransporte.go` | Lógica de teletransporte entre portais. |
| `src/entidades/geometria/*` | Retângulos e pontos para posições e colisão. |
| `src/enum/*` | Tipos, cores e constantes string usadas em entidades e componentes. |
| `src/utils/constantes.go` | Tamanhos e constantes de velocidade. |
| `src/assets/*.go` | Carrega fonte embutida e desenho de texto. |
| `src/assets/sprits/portal/teste.go` | Protótipo de portal por sprite. |

## 13. Conclusão

O projeto tem uma estrutura organizada e bem segmentada. A arquitetura de cena e o ECS leve deixam o código extensível. Há várias áreas preparadas para expansão, como IA, input global e armas/projéteis.

Se quiser, posso também gerar um diagrama de arquitetura simplificado ou sugerir melhorias específicas de refatoração em uma segunda etapa.
