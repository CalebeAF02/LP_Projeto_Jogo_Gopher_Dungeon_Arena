# Resposta aos Critérios de Avaliação

### 🎮 Gopher Dungeon Arena
**Aluno:** Calebe (Desenvolvedor Solo)

---

## 1. Linguagem (Histórico, Versão)
O projeto foi desenvolvido em Go (Golang), utilizando a versão **1.22+** (especificada no `go.mod`). Go foi criada pelo Google em 2009 por Robert Griesemer, Rob Pike e Ken Thompson para resolver problemas de escalabilidade, concorrência e lentidão de compilação. No contexto acadêmico de LPs, o projeto prova a eficiência de Go no desenvolvimento de sistemas gráficos de tempo real e simulações complexas, operando de forma nativa e performática sem a necessidade de máquinas virtuais pesadas.

---

## 2. Projeto: Premissas, Usuário, Domínio
*   **Domínio:** Desenvolvimento de jogos casuais/arcade 2D no estilo *dungeon crawler*, processado a 60 FPS estáveis através do framework gráfico **Ebitengine**.
*   **Usuário-alvo:** Entusiastas de jogos de arena ágeis com mecânicas de sobrevivência e pontuação.
*   **Premissas:** Demonstrar o uso prático de um motor de jogo estruturado puramente sob a arquitetura **ECS (Entity Component System)**, validando a viabilidade de Go em gerenciar estados geométricos, renderização dinâmica, câmera móvel, mini-mapas e inteligência artificial assíncrona.

---

## 3. Construtores de Tipos
A ausência de herança de classes tradicional em Go foi totalmente suprida através de construtores ortogonais da linguagem:

1.  **Composição de Structs:** Dados primitivos puros encapsulados em structs específicas como `Vida`, `Nivel`, `Pontuacao` e `Corpo` (geometria de retângulos).
2.  **Polimorfismo por Interfaces:** Contratos rígidos como `Entidade` (para os personagens), `ICena` (para gerenciamento de telas) e `Movimentador` (para algoritmos de IA).
3.  **Mapeamento Dinâmico:** Uso de mapas flexíveis de tipos genéricos `map[string]interface{}` acoplados dentro de cada entidade para armazenar dinamicamente seus componentes ECS, permitindo mutação e consulta de propriedades em tempo de execução.

---

## 4. Legibilidade
A legibilidade do código é garantida pela simplicidade sintática nativa de Go (apenas 25 palavras-chave) e pela padronização rígida imposta pelo formatador universal `gofmt`. Arquiteturalmente, a legibilidade foi otimizada pela separação modular estrita em pacotes:

*   `src/cenas/`: Controla a máquina de estados finitos das telas (Menu Iniciar, Menu Pause, Cena Jogo, Telas Informativas).
*   `src/ecs/`: Contratos base, sistema de câmera e projeção de coordenadas para o Mini Mapa.
*   `src/sistema/`: Módulos de lógica isolada (Input, Movimento, Colisão, Spawn, Desenho, Debug).
*   `src/enum/` e `src/utils/`: Constantes estritas, cores e tamanhos físicos do mundo.

---

## 5. Capacidade de Escrita (Writability)
A expressividade da linguagem permitiu codificar um motor modular complexo de forma ágil através de:

*   **Composição sobre Herança:** Facilidade em injetar novos comportamentos em structs sem criar árvores de dependência complexas.
*   **Uso de Slices e Maps Dinâmicos:** Manipulação rápida da coleção global de entidades no mapa de frames da `CenaJogo`.
*   **Modularidade de Pacotes:** Namespaces limpos que reduzem o acoplamento. A sintaxe direta facilitou a escrita solo de 9 lógicas de movimentação independentes delegadas sob uma única interface.

---

## 6. Confiabilidade (Reliability)
A robustez do jogo apoia-se fortemente nas garantias de segurança de Go:

*   **Tipagem Estática Forte:** Erros de atribuição de dados ou incompatibilidade geométrica são capturados estritamente em tempo de compilação, blindando a execução do loop de jogo.
*   **Tratamento de Erro Explícito:** Funções críticas como o ciclo `Update() error` e os carregadores de fontes/assets via diretiva embutida `//go:embed` implementam o padrão idiomático `if err != nil`, tratando falhas imediatamente na inicialização antes que gerem pânicos silenciosos na memória.
*   **Segurança de Estado:** Métodos atômicos (`EstaVivo()`, `PerdeSangue()`) encapsulados nas structs garantem a integridade das variáveis durante os frames de combate.

---

## 7. Exemplos Práticos no Código
Os pontos altos da implementação que devem ser demonstrados na apresentação incluem:

*   **O Loop Principal ECS:** O método `CenaJogo.Update()` executando ordenadamente os sistemas isolados (`SistemaInput`, `SistemaSpawn`, `SistemaMovimento`, `SistemaColisao`) e limpando entidades mortas.
*   **Algoritmos de Movimentação (IA):** A implementação da interface `Movimentador` ramificada em 9 variações matemáticas (Vertical, Horizontal, Diagonal, Lógico Linha, Lógico Duplo, etc.), identificadas por cores na arena.
*   **Física de Colisão AABB:** O método `VaiColidir` do `SistemaColisao` calculando as caixas delimitadoras nos eixos X e Y separadamente para evitar travamentos ou tunelamento em paredes e portais.
*   **Barramento de Teletransporte:** O fluxo coordenado entre os componentes `ENVIANDO_TELETRANSPORTE` (Portal de Entrada) e `RECEBENDO_TELETRANSPORTE` (Portal de Saída) que rotaciona, retém e transfere os bots no espaço geométrico.

---

## 8. Projeto
O software está 100% functional, concretizado em um produto executável completo. A engenharia solo entregou:

*   Mundo expandido de jogo (2560x1440 pixels) com câmera dinâmica inteligente que segue o jogador.
*   Mini Mapa HUD com projeção em escala reduzida em tempo real.
*   Geração procedural/estática de labirintos, portais funcionais de teletransporte e comida colecionável.
*   Combate em tempo real com redução matemática de HP baseada em nível.
*   Sistema avançado de Inspeção via Reflexão (`reflect`) que permite varrer a memória de todas as entidades ativas pressionando as teclas `F1` e `F2` (Debug).

---

## 9. Considerações Finais
O projeto consolida as melhores práticas de Go para o desenvolvimento de software estruturado. Ao focar em recursos que realmente diferenciam Go de linguagens tradicionais de baixo nível — como a formatação nativa, o tempo de compilação instantâneo, a flexibilidade da composição por structs/interfaces e a segurança em tempo de execução —, o trabalho prova que é possível conceber e arquitetar sistemas altamente complexos e interativos com alto desempenho através de um desenvolvimento limpo, seguro e produtivo.
