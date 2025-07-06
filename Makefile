hello:
	echo "Hello, World"

server:
	go run ./cmd/server/main.go

build_cli:
	go build -C ./cmd/client/ -o ../../bin/client_cli.exe

cli:
	./bin//client_cli.exe

evans:
	evans --proto proto/v1/todo_service.proto -p 3000 repl    

unit_test: 
	go test -v -cover ./test

bump: 
	go get -u ./...

docker_up:
	docker compose up -d

docker_down:
	docker compose down

docker_prune:
	docker system prune

about: ## Display info related to the build
	@echo "OS: ${OS}"
	@echo "Shell: ${SHELL} ${SHELL_VERSION}"
	@echo "Protoc version: $(shell protoc --version)"
	@echo "Go version: $(shell go version)"
	@echo "Go package: ${PACKAGE}"
	@echo "Openssl version: $(shell openssl version)"

help: ## Show this help
	@${HELP_CMD}


# TODO : migration commands

# migration_up: 
# 	migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose up
# migration_down: 
#	migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" -verbose down
# migration_fix: 
# 	migrate -path database/migration/ -database "postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" force VERSION