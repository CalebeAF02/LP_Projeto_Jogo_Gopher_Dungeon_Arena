# Relatório Final de Entrega – Gopher Dungeon Arena

Este documento foi organizado para servir como material de entrega final ao professor, reunindo os critérios de avaliação, as evidências concretas no projeto e os principais links de acesso.

## 1. Informações gerais
- Projeto: Gopher Dungeon Arena
- Linguagem utilizada: Go
- Repositório: https://github.com/CalebeAF02/LP_Projeto_Jogo_Gopher_Dungeon_Arena
- Wiki do projeto: https://github.com/CalebeAF02/LP_Projeto_Jogo_Gopher_Dungeon_Arena/wiki/Wiki
- Vídeo de apresentação: https://youtu.be/iEJp9QhSvJs

## 2. Descrição do projeto
O projeto Gopher Dungeon Arena consiste em um jogo 2D desenvolvido em Go, com arquitetura modular inspirada em ECS, câmera, minimapa, entidades, colisões, IA e progressão de fases. A implementação demonstra o uso prático da linguagem em um contexto de desenvolvimento de jogos, abordando organização de código, modularização e concorrência.

## 3. Análise da linguagem Go
A linguagem Go foi escolhida por sua simplicidade, legibilidade, compilação rápida, tipagem estática e suporte nativo à concorrência. O projeto evidencia o uso de structs, interfaces, pacotes, funções e estruturas organizadas para facilitar manutenção e expansão.

## 4. Critérios de avaliação e comprovação

### 4.1. Histórico e versão da linguagem
- Evidência:
  - Arquivo de configuração do módulo: [go.mod](../../go.mod)
  - Descrição geral do projeto: [README.md](../../README.md)

### 4.2. Premissas, usuário e domínio
- Evidência:
  - Visão geral do jogo: [README.md](../../README.md)
  - Documentação técnica: [docs/tecnico/documentacao_jogo.md](../tecnico/documentacao_jogo.md)

### 4.3. Construtores
- Evidência:
  - Estruturas e entidades do jogo: [src/entidades](../../src/entidades)
  - Interfaces do projeto: [src/interfaces](../../src/interfaces)
  - Exemplo de criação de cena e entidades: [src/cenas/cenaJogo.go](../../src/cenas/cenaJogo.go)

### 4.4. Legibilidade
- Evidência:
  - Estrutura modular do projeto: [src](../../src)
  - Organização por responsabilidades: [src/cenas](../../src/cenas), [src/sistema](../../src/sistema), [src/entidades](../../src/entidades)

### 4.5. Capacidade de escrita
- Evidência:
  - Uso de estruturas de dados e composição: [src/ecs/ecs.go](../../src/ecs/ecs.go)
  - Lógica de movimentação e entidades: [src/componentes](../../src/componentes)
  - Carregamento e persistência de fases: [src/nivel](../../src/nivel)

### 4.6. Confiabilidade
- Evidência:
  - Leitura e processamento de dados: [src/nivel/carregar.go](../../src/nivel/carregar.go)
  - Estrutura do fluxo principal do jogo: [src/cenas/cenaJogo.go](../../src/cenas/cenaJogo.go)
  - Observação: o projeto já possui execução e compilação válidas, mas ainda não há testes automatizados implementados.

### 4.7. Custo e outros
- Evidência:
  - Documentação técnica: [docs/tecnico/documentacao_jogo.md](../tecnico/documentacao_jogo.md)
  - Documentação sobre concorrência e paralelismo: [docs/tecnico/paralelismo_go.md](../tecnico/paralelismo_go.md)

### 4.8. Exemplos
- Evidência:
  - Exemplo de concorrência na IA: [src/sistema/sistema_de_ia.go](../../src/sistema/sistema_de_ia.go)
  - Exemplo de concorrência no movimento: [src/sistema/sistema_de_movimento.go](../../src/sistema/sistema_de_movimento.go)

### 4.9. Projeto
- Evidência:
  - Ponto de entrada do jogo: [main.go](../../main.go)
  - Fluxo principal do jogo: [src/game.go](../../src/game.go)
  - Compilação verificada com sucesso pelo comando: `go build ./...`

### 4.10. Qualidade do vídeo
- Evidência:
  - Roteiro do vídeo: [docs/apresentacao/roteiro_video.md](../apresentacao/roteiro_video.md)
  - Guia de slides: [docs/apresentacao/guia_slides_video.md](../apresentacao/guia_slides_video.md)

### 4.11. Apresentação
- Evidência:
  - Material de apoio para apresentação: [docs/apresentacao](../apresentacao)
  - Resposta aos critérios avaliativos: [docs/avaliacao/Resposta_aos_criterios_de_avaliacao.md](Resposta_aos_criterios_de_avaliacao.md)

### 4.12. Site/Demonstração
- Evidência:
  - Repositório GitHub: https://github.com/CalebeAF02/LP_Projeto_Jogo_Gopher_Dungeon_Arena
  - Wiki do projeto: https://github.com/CalebeAF02/LP_Projeto_Jogo_Gopher_Dungeon_Arena/wiki/Wiki
  - Instruções de execução: [README.md](../../README.md)

## 5. Atividades desenvolvidas
- Planejamento e organização do projeto
- Implementação do jogo e estrutura modular
- Desenvolvimento de sistemas de movimento, colisão, câmera e minimapa
- Implementação de lógica de IA e concorrência em Go
- Documentação técnica e elaboração da wiki
- Produção do vídeo de apresentação e organização da entrega final

## 6. Conclusão
O projeto Gopher Dungeon Arena demonstra a aplicação prática da linguagem Go no desenvolvimento de um jogo com arquitetura organizada, modularidade, lógica de IA e uso de concorrência. A implementação cumpre os principais requisitos da disciplina e está preparada para servir como material de apresentação e avaliação final.
