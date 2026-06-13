# Revisão Geral do Jogo Gopher Dungeon Arena

Este documento descreve a arquitetura, organização, funcionalidades e arquivos do projeto `Gopher_Dungeon_Arena`.

## 1. Visão geral do projeto

O jogo é construído com Go e usa a biblioteca `ebiten` para renderização 2D. A estrutura principal é centrada em um sistema de cenas (`cenas`) e um pequeno ECS (entidade-componentes). O arquivo `main.go` inicia o jogo, cria a janela e executa o loop principal através de `ebiten.RunGame`.

## 2. Organização de pastas

- `main.go`: ponto de entrada do jogo.
- `go.mod`, `go.sum`: dependências e módulo Go.
- `src/`: código-fonte principal.
  - `assets/`: fontes e recursos de desenho.
  - `cenas/`: cenas do jogo (menu inicial, palco do jogo, pause).
  - `config/`: constantes de configuração e utilitários aleatórios.
  - `ecs/`: abstrações de entidade e sistemas de câmera/mini mapa.
  - `entidades/`: implementação das entidades do jogo.
  - `enum/`: tipos enumerados de entidades, componentes e cores.
  - `interfaces/`: contratos para cenas, jogo, sistemas e movimentação.
  - `sistema/`: lógica de atualização e desenho do jogo.
  - `utils/`: constantes, desenho e utilitários de tamanho.

## 3. Fluxo principal

1. `main.go` chama `src.NovoGame()`.
2. `src/game.go` cria a cena inicial `CenaMenuIniciar` e a cena de jogo `CenaJogo`.
3. A cena corrente é atualizada e desenhada no loop de `ebiten`.
4. Quando o jogador pressiona `ENTER`, o fluxo troca para `CenaJogo`.
5. Durante o jogo, o jogador pode pausar com `P` e voltar com `SPACE` no menu de pause.

## 4. Arquivos principais e funcionalidades

### `main.go`
- Define a janela do jogo com `ebiten.SetWindowSize` e `SetWindowTitle`.
- Executa `ebiten.RunGame(game)`.
- Usa constantes de `config` para largura, altura e nome.

### `src/game.go`
- Representa o jogo com as cenas `CenaCorrente` e `CenaJogo`.
- Contém o método `Update()` que delega para a cena corrente.
- Contém o método `Draw()` que desenha a cena corrente.
- Métodos de transição:
  - `IniciarJogo()` troca para a cena de jogo.
  - `Pausar()` cria e define a cena de pausa.
  - `Voltar()` retorna à cena de jogo.
  - `Sair()` encerra o processo.

### `src/config/configuracoes.go`
- Define constantes de janela, mundo e mini mapa.
- Fornece geradores aleatórios com `GeradorAleatorio()`.
- `XAleatorio` e `YAleatorio` usam a largura/altura da tela.

### `src/interfaces/`
- `i_game.go`: define a interface do jogo, com transições de cena.
- `i_cena.go`: interface básica de cena com `Update`, `Draw` e `GetNome`.
- `i_cenajogo.go`: interface de cena de jogo, expõe entidades, câmera, mini mapa, colisões e métodos de criação.
- `i_sistema_atualizar.go`: contrato para sistemas que atualizam lógica.
- `i_sistema_desenhar.go`: contrato para sistemas de desenho.
- `i_sistema_colisao.go`: contrato para sistema de colisão.
- `i_movimentador.go`: contrato para comportamentos de movimento (ANALISAR EM ARQUIVO separado, não lido ainda). 

## 5. Cenas do jogo

### `src/cenas/cenaMenuIniciar.go`
- Cena inicial com menu de título.
- Detecta `ENTER` para iniciar jogo e `ESC` para sair.
- Desenha texto na tela com `ebiten/text/v2`.
- Usa `assets.Fonte` para renderizar as fontes.

### `src/cenas/cenaMenuPause.go`
- Cena de pausa simples.
- Detecta `SPACE` para voltar ao jogo e `ESC` para sair.
- Exibe instruções de controle na tela.

### `src/cenas/cenaJogo.go`
- Cena principal do jogo.
- Cria o mundo, câmera, mini mapa e sistema de colisão.
- Instancia sistemas de atualização e desenho.
- Executa spawn de jogadores, portais, paredes e bots.
- `Update()` chama `Input`, `OrganizarCamera`, todos os sistemas e depois remove entidades mortas.
- `Draw()` passa para os sistemas de desenho.
- `OrganizaPosicaoAleatoriaBot()` busca uma posição livre para bots.
- Tem contagem de bots mortos e método `ContarEntidadesMortas()`.

## 6. ECS e entidades

### `src/ecs/ecs.go`
- Define `EntidadeID` como `int`.
- Define interface `Entidade` com:
  - `GetID()`
  - `GetTipo()`
  - `GetComponente(id string)`
  - `ExisteComponente(id string)`
  - `Atualizar()`
  - `Desenhar(screen *ebiten.Image)`
  - `DesenharMapa(screen *ebiten.Image, mapaX float64, mapaY float64)`

### `src/enum/entidades/entidade_tipo.go`
- Enumeração de tipos: `BOT`, `JOGADOR`, `TIME`, `PAREDE`, `PORTAL_ENTRADA`, `PORTAL_SAIDA`.

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
