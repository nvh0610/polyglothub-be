GO_CMD = go
include .env
export $(shell sed 's/=.*//' .env)


run:
	$(GO_CMD) run main.go

create-migration:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(name)

migrate-up:
	migrate -path $(MIGRATION_DIR) -database "$(DATABASE__TYPE)://$(DATABASE__USER):$(DATABASE__PASSWORD)@tcp($(DATABASE__HOST):$(DATABASE__PORT))/$(DATABASE__NAME)" up

migrate-down:
	migrate -path $(MIGRATION_DIR) -database "$(DATABASE__TYPE)://$(DATABASE__USER):$(DATABASE__PASSWORD)@tcp($(DATABASE__HOST):$(DATABASE__PORT))/$(DATABASE__NAME)" down

docker-run:
	docker run --name mysql-learn-language -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=learn-language -p 3306:3306 -d mysql:latest
	docker run --name redis-learn-language -p 6379:6379 -d redis:latest

docker-stop:
	docker stop mysql-learn-language
	docker rm mysql-learn-language
	docker stop redis-learn-language
	docker rm redis-learn-language


setup:
	make docker-run
	@echo "Waiting for MySQL to be ready..."
	@until docker exec mysql-learn-language mysql -u$(DATABASE__USER) -p$(DATABASE__PASSWORD) -e "SELECT 1" > /dev/null 2>&1; do \
		echo "Waiting for MySQL..."; \
		sleep 5; \
	done
	make migrate-up

teardown:
	make migrate-down
	make docker-stop
