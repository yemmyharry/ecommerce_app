run:
	go run .
up:
	go run ./internal/migration/migration.go

start:
	docker-compose up --build

stop:
	docker-compose down

