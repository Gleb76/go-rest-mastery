.PHONY: help deps tidy test lint docker-up docker-down migrate-up migrate-down

help:
	@echo "Go REST Mastery — команды для обучения"
	@echo ""
	@echo "  make deps        — скачать зависимости"
	@echo "  make tidy        — go mod tidy"
	@echo "  make test        — запустить все тесты в workspace/"
	@echo "  make docker-up   — PostgreSQL + Redis"
	@echo "  make docker-down — остановить контейнеры"
	@echo "  make migrate-up  — применить миграции (capstone)"
	@echo "  make run-ch01    — запустить главу 01 starter"

deps:
	go mod download

tidy:
	go mod tidy

test:
	@find workspace -name go.mod -execdir go test ./... \; 2>/dev/null || echo "Пока нет проектов в workspace/ — создай их по главам"

docker-up:
	docker compose up -d

docker-down:
	docker compose down

migrate-up:
	cd capstone/starter && go run ./cmd/migrate up

run-ch01:
	go run ./chapters/01-http-tcp-basics/starter/main.go

run-ch02:
	go run ./chapters/02-net-http-server/starter/main.go

run-ch05:
	go run ./chapters/05-json-and-encoding/starter/main.go
