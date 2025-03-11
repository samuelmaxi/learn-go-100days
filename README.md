# Desafio 100 dias Golang

## **Dias 1 a 10 – Fundamentos da Linguagem**

- [x]  **Hello, World!** – Crie um programa que imprime "Hello, World!" na tela.
- [x]  **Variáveis e Tipos** – Declare e imprima diferentes tipos de variáveis.
- [x]  **Operadores Aritméticos** – Implemente uma calculadora básica com soma, subtração, multiplicação e divisão.
- [x]  **Entrada do Usuário** – Leia e imprima o nome e a idade do usuário.
- [x]  **Condicionais (if/else)** – Verifique se um número é par ou ímpar.
- [x]  **Switch Case** – Implemente um programa que imprime o nome do mês com base em um número de 1 a 12.
- [x]  **Laço for** – Imprima a tabuada de 1 a 10.
- [x]  **Arrays e Slices** – Crie uma lista de frutas e mostre como adicionar e remover itens.
- [x]  **Funções** – Crie uma função para calcular o fatorial de um número.
- [x]  **Defer e Panic** – Explore como o `defer` e `panic` funcionam em diferentes cenários.

---

## **Dias 11 a 20 – Estruturas de Dados e Funções Avançadas**

- [x]  **Maps** – Crie um dicionário que armazena a capital de cada país.
- [x]  **Structs** – Crie uma struct para representar um aluno (nome, idade, notas).
- [x]  **Métodos em Structs** – Adicione um método na struct `Aluno` para calcular a média.
- [x]  **Funções Variádicas** – Crie uma função que calcula a soma de números variáveis.
- [x]  **Funções Anônimas** – Use uma função anônima para imprimir números pares até 20.
- [x]  **Ponteiros** – Crie um exemplo de como manipular valores com ponteiros.
- [x]  **Funções como Parâmetro** – Passe uma função como argumento para outra função.
- [x]  **Recursividade** – Implemente a sequência de Fibonacci usando recursão. _
- [x]  **Interfacing** – Crie uma interface para representar diferentes formas geométricas. _
- [x]  **Type Assertions** – Implemente um exemplo de `type assertion` em uma interface.

---

## **Dias 21 a 30 – Manipulação de Arquivos e JSON**

- [x]  **Escrever em Arquivo** – Crie e escreva em um arquivo texto.
- [x]  **Ler Arquivo** – Leia e imprima o conteúdo de um arquivo.
- [x]  **Excluir Arquivo** – Crie uma função que exclui um arquivo.
- [x]  **Manipulação de JSON** – Converta uma struct para JSON.
- [x]  **Ler JSON de Arquivo** – Implemente a leitura de um arquivo JSON para uma struct.
- [x]  **API Simples** – Crie uma API que retorna um JSON simples.
- [x]  **Unmarshal JSON** – Decodifique um JSON para uma struct.
- [x]  **Marshal JSON** – Converta uma struct em JSON e imprima no terminal.
- [x]  **HTTP Client** – Realize uma requisição HTTP GET para uma API pública.
- [x]  **HTTP Server Básico** – Crie um servidor HTTP simples.

---

## **Dias 31 a 40 – Conexões HTTP e API Restful**

- [x]  **Roteamento HTTP** – Use o pacote `net/http` para criar rotas personalizadas.
- [x]  **Query Params** – Crie uma rota que aceita parâmetros na query string.
- [x]  **Middleware** – Implemente um middleware básico.
- [x]  **API CRUD** – Implemente uma API para CRUD de tarefas (sem banco de dados).
- [x]  **Gin Gonic Básico** – Crie uma API usando o framework Gin Gonic.
- [x]  **Tratamento de Erros** – Adicione erros personalizados na API.
- [x]  **Post com Body JSON** – Crie uma rota POST que recebe um JSON como entrada.
- [x]  **Autenticação Simples** – Adicione uma autenticação básica na API.
- [x]  **CORS** – Configure CORS na API.
- [x]  **Rate Limiting** – Adicione limite de requisições por IP.

---

## **Dias 41 a 50 – Concorrência e Goroutines**

- [x]  **Goroutines Básicas** – Crie duas goroutines que imprimem mensagens diferentes.
- [x]  **WaitGroup** – Use `sync.WaitGroup` para esperar o fim de várias goroutines.
- [x]  **Mutex** – Implemente um contador usando goroutines e mutex.
- [x]  **Channels** – Crie uma comunicação entre duas goroutines usando canais.
- [x]  **Buffered Channels** – Teste canais bufferizados.
- [x]  **Select Statement** – Use `select` para alternar entre canais.
- [x]  **Timeout com Channels** – Adicione timeout usando canais e `time.After`.  _
- [x]  **Workers Pool** – Implemente um pool de workers com goroutines.
- [x]  **Context API** – Adicione cancelamento com `context`.
- [x]  **Pipeline de Dados** – Crie um pipeline de processamento usando canais.

---

## **Dias 51 a 60 – Bancos de Dados e ORM**

- [x]  **SQLite Básico** – Conecte-se a um banco SQLite e crie uma tabela.
- [x]  **Inserção de Dados** – Insira registros em uma tabela SQLite.
- [x]  **Leitura de Dados** – Leia e imprima os dados do banco SQLite.
- [x]  **Atualização e Exclusão** – Adicione operações de update e delete.
- [x]  **PostgreSQL Básico** – Conecte-se a um banco PostgreSQL.
- [x]  **ORM com GORM** – Use a biblioteca GORM para manipular registros.
- [x]  **Migrações com GORM** – Crie migrações automáticas.
- [x]  **Relacionamento 1**– Implemente relacionamento entre duas tabelas.
- [x]  **Autenticação com Banco** – Crie um login usando banco de dados.
- [x]  **Cache Simples** – Adicione um cache na aplicação usando um map.

---

## **Dias 61 a 70 – Testes e Deploy**

- [x]  **Testes Unitários** – Escreva testes para uma função básica.
- [x]  **Testes com Mock** – Implemente testes com mocks.
- [x]  **Cobertura de Testes** – Gere um relatório de cobertura de testes.
- [x]  **Benchmarking** – Adicione testes de benchmark.
- [x]  **Dockerização** – Crie um Dockerfile para a aplicação Go.
- [x]  **Docker Compose** – Use Docker Compose para subir a aplicação e banco.
- [x]  **Configuração por Ambiente** – Adicione variáveis de ambiente.
- [x]  **Deploy Local** – Realize o deploy da aplicação Go localmente.
- [x]  **Deploy na Nuvem** – Suba a aplicação para uma plataforma como Heroku.
- [x]  **Usar um Makefile** – Criar um `Makefile` para rodar testes e build automaticamente.

---

## **Dias 71 a 80 – Desafios Avançados**

- [x]  **JWT Authentication** – Implemente autenticação JWT na API.
- [x]  **OAuth2 Local** – Simule um fluxo de autenticação OAuth2 sem depender de provedores externos.
- [x]  **WebSocket** – Crie um chat simples com WebSocket.
- [x]  **Testes Automatizados** – Escreva testes unitários
- [x]  **Eventos com Channels** – Utilize `chan` para implementar comunicação assíncrona entre goroutines.
- [x]  **Middleware Personalizado** – Crie um middleware para log de requisições e auditoria.
- [x]  **API GraphQL** – Adicione uma API GraphQL.
- [x]  **Documentação com Swagger** – Adicione documentação Swagger à API.
- [x]  **Mock de API Externa** – Simule o consumo de uma API pública utilizando um servidor local.
- [x]  **Refatoração e Otimização** – Revise seu código e aplique boas práticas de organização.

---

## **Dias 81 a 90 – Desafios Arquivos de Configuração**

- [x]  **Configuração via Arquivo** – Carregue configurações da aplicação a partir de um arquivo JSON ou YAML.
- [x]  **Monitoramento com Logs** – Registre logs detalhados de requisições e eventos do sistema.
- [x]  **Tracing Local** – Meça tempos de execução e trace chamadas internas da API.
- [x]  **Orquestração com Goroutines** – Gerencie múltiplas goroutines em um fluxo de trabalho concorrente.
- [x]  **Agendamento de Tarefas** – Implemente um job scheduler local para executar funções periodicamente.
- [x]  **Testes de Carga com Benchmarking** – Utilize `testing.B` para medir a performance de funções críticas.
- [x]  **Webhooks Locais** – Simule a recepção e envio de eventos via webhooks sem servidores externos.
- [x]  **Boas Práticas de Segurança** – Aplique validações e proteções contra injeção de dados.
- [x]  **Feature Flags Locais** – Implemente um sistema de ativação de recursos baseado em arquivos de configuração.
- [x]  **Gerenciamento de Erros** – Estruture um sistema robusto de tratamento de erros

---

## **Dias 91 a 100 – Revisão**

- [x]  **Automação com Makefile** – Crie um Makefile para automatizar tarefas comuns no projeto.
- [x]  **Testes Unitários** – Escreva testes unitários para funções principais da aplicação.
- [x]  **Testes de Integração** – Valide a comunicação entre componentes internos sem dependências externas.
- [x]  **Concurrency Patterns** – Implemente padrões avançados como fan-in, fan-out e worker pools.
- [x]  **Modularização do Código** – Estruture o projeto em pacotes organizados para melhor manutenção.
- [x]  **Construção de CLI** – Desenvolva uma ferramenta de linha de comando com `cobra` ou `flag`.
- [x]  **Cache em Memória** – Implemente um sistema de cache simples usando `map` e `sync.Mutex`.
- [x]  **Persistência de Dados Local** – Salve e carregue dados em arquivos JSON ou SQLite sem bancos externos.
- [x]  **Manipulação de Strings** – Escreva funções para formatar e modificar strings.
- [x]  **Leitura de Arquivos** – Crie um programa que leia e exiba o conteúdo de um arquivo `.txt`.
