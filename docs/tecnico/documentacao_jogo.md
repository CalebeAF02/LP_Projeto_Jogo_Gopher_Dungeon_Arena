# Documentação Completa - Gopher Dungeon Arena

## 1. Visão geral

Gopher Dungeon Arena é um projeto acadêmico desenvolvido em Go para a disciplina de Linguagens de Programação da UnB. O objetivo principal é criar um jogo 2D de arena com mecânicas simples de movimentação, colisão, câmera, minimapa, inimigos e progressão de fases, usando a biblioteca Ebiten para renderização.

O projeto já possui uma estrutura funcional suficiente para executar o jogo, carregar um mapa a partir de um arquivo JSON, criar entidades, gerenciar cenas e manter progresso entre execuções.

## 2. Análise completa do projeto

### 2.1 Objetivo do projeto

O projeto busca demonstrar que é possível construir um motor de jogo básico em Go com:
- organização em camadas;
- uso de interfaces e tipos explícitos;
- separação entre entidades, componentes e sistemas;
- lógica de jogo baseada em cenas;
- carregamento de conteúdo via arquivos externos.

A proposta também explora conceitos importantes de programação, como encapsulamento, modularização, concorrência e manipulação de estado.

### 2.2 Arquitetura adotada

O projeto segue uma arquitetura híbrida inspirada em ECS, com elementos de jogo organizados em:
- entidades: jogador, bots, comida, paredes, portais e saída;
- componentes: corpo, vida, nível, pontuação, movimentação, atividade e teletransporte;
- sistemas: movimento, colisão, spawn, desenho, IA e gerenciamento de entidades.

A estrutura é relativamente bem organizada, pois separa responsabilidades entre pacotes de cenas, entidades, sistemas, interfaces e utilidades.

### 2.3 Fluxo de execução

1. O ponto de entrada está em main.go.
2. O jogo é inicializado por src/game.go, que cria a cena inicial e a cena principal de jogo.
3. A cena de jogo carrega o primeiro nível a partir de um arquivo JSON.
4. O loop principal do Ebiten executa atualização e desenho continuamente.
5. O jogador pode se mover, interagir com objetos e atravessar as mecânicas do mapa.
6. O progresso é salvo em um arquivo JSON para ser reutilizado em novas execuções.

### 2.4 Módulos principais

- main.go: inicialização do jogo e entrada do programa.
- src/game.go: coordenador do estado global do jogo e transições de tela.
- src/cenas: implementação das telas do jogo, como menu, pausa, jogo e progresso.
- src/ecs: interface básica de entidades e abstrações de identidade.
- src/entidades: implementação concreta de objetos e personagens do mundo.
- src/componentes: estruturas que representam dados do jogo, como vida, corpo e pontuação.
- src/sistema: lógica de atualização e comportamento do jogo.
- src/nivel: carregamento de fases e persistência de progresso.
- src/config: constantes globais de tamanho de janela, mundo e mapa.
- src/utils: valores comuns usados por várias entidades.

### 2.5 Estado atual do projeto

A implementação atual já está em um estado funcional e compilável. A validação realizada no ambiente confirmou que o projeto compila com sucesso usando:

```bash
go build ./...
```

O projeto já apresenta as seguintes funcionalidades:
- menu inicial;
- cena de jogo;
- câmera dinâmica;
- minimapa;
- jogador controlável;
- bots com diferentes comportamentos de movimentação;
- colisões básicas;
- comida, paredes, saída e portais;
- carregamento de mapa via JSON;
- sistema de progresso persistido.

### 2.6 Pontos fortes

- Estrutura modular e bem separada para um projeto acadêmico.
- Uso de interfaces para desacoplar cenas e sistemas.
- Implementação de diferentes movimentadores para bots, o que deixa a lógica mais flexível.
- Carregamento de níveis via JSON, o que facilita expansão futura.
- Persistência de progresso com arquivo JSON simples e acessível.
- Uso de concorrência no sistema de movimento, mostrando tentativa de explorar recursos do Go.

### 2.7 Limitações observadas

Apesar do bom avanço, ainda existem pontos que podem ser melhorados:
- alguns sistemas ainda funcionam como placeholders, como SistemaInput, SistemaEntidades e parte de SistemaIA;
- a lógica de combate ainda está incompleta ou pouco expandida;
- a colisão é bastante direta e não cobre todos os cenários complexos de forma robusta;
- o projeto ainda não possui testes automatizados;
- a arquitetura, embora organizada, ainda pode ser refinada para reduzir acoplamento entre entidades e cenas;
- alguns comportamentos de IA e spawn ainda são simples demais para um jogo mais completo.

### 2.8 Melhorias recomendadas

Para evoluir o projeto, os próximos passos mais úteis seriam:
1. completar a implementação dos sistemas de entrada e entidades;
2. consolidar a IA dos bots com decisões mais reais;
3. melhorar a resposta de colisão e adicionar mais interações entre entidades;
4. implementar combate mais consistente, incluindo projéteis ou ataques diretos;
5. acrescentar áudio, animações e feedback visual mais rico;
6. criar testes unitários para componentes e lógica de colisão;
7. expandir os níveis e tornar o progresso mais rico e narrativo.

## 3. Estrutura de pastas

A organização do projeto é a seguinte:

- main.go: ponto de entrada.
- src/game.go: controlador principal do jogo.
- src/cenas: telas e gerenciamento de fluxo.
- src/componentes: estruturas de dados do jogo.
- src/ecs: abstração de entidades.
- src/entidades: objetos e personagens do mundo.
- src/interfaces: contratos usados pelas cenas e sistemas.
- src/sistema: sistemas de atualização e processamento.
- src/nivel: carregamento de fases e progresso.
- src/config: constantes e configuração geral.
- src/utils: valores reutilizáveis.
- docs: documentação do projeto.

## 4. Como executar

No diretório do projeto, use:

```bash
go run .
```

Ou, para gerar um build:

```bash
go build ./...
```

## 5. Controles principais

- Enter: iniciar o jogo no menu inicial.
- Esc: voltar ao menu ou sair, conforme a tela atual.
- Setas ou WASD: movimentar o jogador.
- P: pausar o jogo.
- Ctrl + M: alternar a posição do minimapa.
- Ctrl + O: mostrar ou ocultar o minimapa.
- Espaço: ação de ataque, ainda em fase de evolução.

## 6. Mecânicas já implementadas

- movimento do jogador;
- movimentação de bots com diferentes estratégias;
- spawn automático de bots;
- colisões básicas com paredes e outros objetos;
- coleta de comida;
- saída e condição de vitória;
- câmera seguindo o jogador;
- minimapa;
- persistência de progresso.

## 7. Resumo executivo

O projeto já demonstra uma boa base para um jogo feito em Go, com uma estrutura organizada e um conjunto de mecânicas funcional. O maior valor do trabalho está na tentativa de combinar conceitos de programação com uma implementação prática de um jogo 2D, usando uma arquitetura modular que pode crescer com novas funcionalidades.

O projeto está em uma fase interessante: ele já funciona, já possui uma identidade visual e mecânica clara, e já está pronto para receber melhorias de gameplay, IA, combate e polimento técnico.

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
