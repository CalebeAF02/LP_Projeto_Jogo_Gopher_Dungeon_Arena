# Wiki do Projeto: Gopher Dungeon Arena

Bem-vindo à documentação oficial do **Gopher Dungeon Arena**. Este espaço armazena guias aprofundados, decisões de design e detalhes técnicos sobre o desenvolvimento do jogo.

---

## 1. Visão Geral do Projeto
O Gopher Dungeon Arena é um jogo de arena com temática de dungeon desenvolvido como projeto universitário para a Universidade de Brasília (UnB), no âmbito da disciplina de Linguagens de Programação (8º semestre). 

O foco principal do desenvolvimento é demonstrar a viabilidade e a eficiência da linguagem Go na criação de jogos estruturados, aplicando conceitos avançados de computação gráfica e arquitetura de software.

---

## 2. Objetivos de Aprendizado e Engenharia
O projeto foi concebido para evidenciar competências práticas em engenharia de software com Go, focando nos seguintes pilares:
* **Concorrência Avançada:** Uso de goroutines e canais para gerenciar sistemas independentes, inteligência artificial de inimigos e lógica interna do jogo de forma assíncrona.
* **Sincronização de Estado:** Implementação de mecanismos seguros para evitar condições de corrida (*race conditions*) durante simulações rápidas de combate e movimento.
* **Tipagem Segura:** Organização rígida dos dados do jogo para mitigar erros em tempo de execução e garantir consistência estrutural.
* **Desempenho Consistente:** Otimização dos laços de repetição de renderização (*game loop*) e das rotinas de colisão.

---

## 3. Tecnologias e Ferramentas
* **Linguagem Principal:** Go (compondo 99.8% da base de código do repositório).
* **Biblioteca Gráfica/Motor:** Ebitengine (Ebiten), utilizada para o gerenciamento de janelas, captura de entradas do usuário (teclado/mouse) e renderização 2D de sprites e mapas.
* **Arquitetura de Código:** Implementação baseada no padrão **ECS (Entity Component System)** para garantir alta modularidade, permitindo que entidades da dungeon (jogadores, inimigos, projéteis) compartilhem comportamentos de forma desacoplada.

---

## 4. Recursos Implementados
* **Mundo e Renderização:** Base estrutural do motor para desenhar cenários e mapas no estilo dungeon.
* **Movimentação e Input:** Captura de comandos do jogador traduzidos em deslocamento em tempo real na arena.
* **Mecânicas de Combate:** Sistema de regras para troca de dano, geração de projéteis e controle de colisões em tempo real.
* **Inteligência Artificial:** Comportamento automatizado e controle de entidades inimigas presentes na arena.

---

## 5. Como Contribuir e Modificar a Wiki
Esta wiki armazena o conhecimento técnico acumulado do projeto. Caso possua permissão de escrita no repositório, você pode editar estas páginas diretamente pela interface web do GitHub ou clonar o repositório da wiki localmente para realizar alterações via Markdown.
