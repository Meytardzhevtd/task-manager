# Переменные
DB_URL = postgres://taskuser:taskpass@localhost:5432/taskdb?sslmode=disable
MIGRATE = migrate -path migrations -database "$(DB_URL)"

# Команды
.PHONY: migrate-up migrate-down migrate-create migrate-force migrate-version

migrate-up:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down 1

migrate-down-all:
	$(MIGRATE) down

migrate-create:
	@read -p "Enter migration name: " name; \
	$(MIGRATE) create -ext sql -seq -dir migrations $$name

migrate-force:
	@read -p "Enter version number: " version; \
	$(MIGRATE) force $$version

migrate-version:
	$(MIGRATE) version

# Запуск БД
db-up:
	docker-compose up -d

db-down:
	docker-compose down

db-reset: db-down db-up migrate-up

# Dev
run:
	go run cmd/main.go

build:
	go build -o bin/task-manager cmd/main.go

tidy:
	go mod tidy