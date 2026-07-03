# Resposta aos Critérios de Avaliação

## 1. Linguagem (histórico, versão)

O projeto foi desenvolvido na linguagem Go, utilizando a versão 1.26.2, conforme definido no arquivo de configuração do módulo. A escolha de Go se justifica pela sua legibilidade, pela produtividade no desenvolvimento e pela capacidade de organizar projetos de forma modular e eficiente.

Além disso, a linguagem foi aplicada em um contexto de desenvolvimento de jogos, mostrando que Go pode ser usada para criar aplicações mais complexas, com estrutura bem definida e boa manutenção do código.

## 2. Projeto: premissas, usuário, domínio

O projeto Gopher Dungeon Arena consiste em um jogo de arena com temática de dungeon, no qual o jogador interage com o ambiente, realiza movimentos, enfrenta inimigos e percorre fases. O objetivo principal é criar uma experiência de jogo funcional, com regras básicas de combate, colisão e progressão.

O domínio do projeto está relacionado à implementação de um motor de jogo simples, com entidades, componentes, sistemas e regras de interação, o que torna o trabalho adequado para demonstrar conceitos de programação orientada a objetos e organização de código em Go.

## 3. Construtores

No projeto, foram utilizados recursos próprios da linguagem Go de forma relevante, como:

- structs para representar entidades do jogo, como jogador, inimigos, objetos e cenários;
- métodos associados a essas estruturas para encapsular comportamentos;
- interfaces para definir contratos entre diferentes partes do sistema;
- organização em pacotes, separando responsabilidades como cenas, entidades, componentes, sistemas e interfaces;
- uso de funções construtoras, como as funções de criação de entidades e objetos do jogo.

Esses recursos mostram que o projeto aproveita características importantes de Go e não se limita a uma implementação simples baseada em estruturas básicas de linguagem.

## 4. Legibilidade

A organização do projeto contribui diretamente para a legibilidade do código. A estrutura foi separada em pastas com papéis bem definidos, como:

- src/cenas: gerenciamento das telas do jogo;
- src/entidades: definição das entidades do mundo;
- src/componentes: componentes que descrevem características e estados;
- src/sistema: lógica de atualização, colisão, entrada e spawn;
- src/interfaces: abstrações para comunicação entre módulos.

Essa divisão torna o código mais claro, facilita a manutenção e demonstra uma forma de desenvolvimento mais profissional e escalável em Go.

## 5. Capacidade de escrita

O projeto utiliza recursos da linguagem que mostram boa capacidade de escrita e organização, como:

- uso de maps e slices para organizar dados e estados;
- uso de interfaces para reduzir o acoplamento entre módulos;
- aplicação de composição de estruturas para representar comportamento e atributos;
- uso de funções e métodos para separar responsabilidades;
- manipulação de arquivos e dados para carregamento de níveis e persistência de progresso.

Esses elementos mostram que o trabalho não apenas implementa funcionalidades, mas também utiliza recursos de Go de maneira consciente e bem estruturada.

## 6. Confiabilidade

O projeto apresenta boa confiabilidade em relação à organização e ao fluxo de execução, principalmente por conta da separação lógica entre as partes do sistema. As estruturas foram organizadas para reduzir problemas de manutenção e facilitar a identificação de erros.

Também há tratamento de erros no processo de leitura e conversão de arquivos de dados, o que demonstra atenção à robustez do sistema. Esse ponto é importante, pois evidencia que o projeto não depende apenas de funcionamento básico, mas também de uma abordagem mais segura e estável.

## 7. Exemplos

Alguns exemplos que podem ser destacados na apresentação:

- criação de entidades com structs e funções construtoras;
- implementação de interfaces para padronizar o comportamento das entidades;
- organização em pacotes para separar a lógica do jogo em módulos;
- carregamento de níveis a partir de arquivos JSON;
- funcionamento do jogo com menus, fases, colisões e progressão.

Esses exemplos mostram na prática como o projeto utiliza conceitos importantes de Go e como eles foram aplicados em um contexto real.

## 8. Projeto

O projeto está funcional e apresenta uma estrutura completa para um jogo simples em Go. Entre os principais pontos implementados, destacam-se:

- menu inicial e telas de jogo;
- movimentação do jogador;
- colisões e interação com objetos do cenário;
- entidades como jogador, inimigos, comida, paredes e portais;
- sistema de fases e progresso;
- organização de código em pacotes e módulos.

Essa implementação demonstra que o trabalho não é apenas teórico, mas também resultou em um produto concreto e executável.

## 9. Considerações finais

O projeto evidencia boas práticas de desenvolvimento em Go, destacando principalmente a organização modular, a utilização de structs e interfaces, a separação de responsabilidades e a construção de uma estrutura de jogo mais profissional. Na apresentação, vale destacar esses pontos, sem perder tempo com conceitos básicos da linguagem que já existem em outras linguagens, como C, e focando nos recursos que realmente diferenciam Go.

Em resumo, o trabalho demonstra um bom uso de Go para desenvolvimento de software mais estruturado, organizado e próximo de uma aplicação real.
