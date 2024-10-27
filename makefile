# Makefile для управления миграциями

DB_DSN := "postgres://myuser:mypassword@127.0.0.1:5433/mydatabase?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Таргет для создания новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations -seq $(NAME)

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# Команда для запуска приложения
run:
	go run cmd/app/main.go
