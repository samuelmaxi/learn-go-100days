GO ?= go
BIN_NAME = main
SRC= ./main.go

.PHONY: build run test fmt lint clean help

build:
	@echo "Compilando o projeto..."
	@$(GO) build -o $(BIN_NAME) $(SRC)
	@echo "Build concluído!"

run: build
	@echo "Executando a aplicação..."
	@./$(BIN_NAME)

test:
	@echo "Executando testes..."
	@$(GO) test ./... -cover

fmt:
	@echo "Formatando código..."
	@$(GO) fmt ./...

lint:
	@echo "Limpando arquivos compilados..."
	@rm -f $(BIN_NAME)

help:
	@echo "	 Comandos disponíveis:"
	@echo "  make build   - Compila o projeto"
	@echo "  make run     - Compila e executa o projeto"
	@echo "  make test    - Executa os testes"
	@echo "  make fmt     - Formata o código"
	@echo "  make lint    - Verifica lint do código"
	@echo "  make clean   - Remove arquivos compilados"
	@echo "  make help    - Mostra esta ajuda"