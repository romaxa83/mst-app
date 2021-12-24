.SILENT:
#=================================
# Variables

prometheus_container = prometheus
#=================================
# Command for docker

.PHONY: up
up: up_docker info

.PHONY: info
info: ps info_domen

.PHONY: up_docker
up_docker:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down --remove-orphans

.PHONY: rebuild
rebuild: down build up_docker info

# флаг -v удаляет все volume (очищает все данные)
.PHONY: down-clear
down-clear:
	docker-compose down -v --remove-orphans

.PHONY: build
build:
	docker-compose build

.PHONY: ps
ps:
	docker-compose ps

#=================================
# Command for App

# run app
.PHONY: run
run:
	go run cmd/app/main.go

# generate swagger documentation
.PHONY: swagger
swagger:
	swag init -g internal/app/app.go

# migrate
.PHONY: migrate_up
migrate_up:
	migrate -path ./migrations -database 'postgres://root:root@192.168.141.1:5432/db?sslmode=disable' up

.PHONY: migrate_down
migrate_down:
	migrate -path ./migrations -database 'postgres://root:root@192.168.141.1:5432/db?sslmode=disable' down
#=================================
# Command for test

#.PHONY: test-run
#test-run: test-running

test-run:
	go test -v ./...

.PHONY: test
test:
	go test -v ./...
	#go test --short -coverprofile=cover.out -v ./...
#go test -coverpkg=./... -coverprofile=c.out ./...
.PHONY: test-coverage
test-coverage:
	go tool cover -func=cover.out
#=================================
# Command for into container

.PHONY: prometheus_bash
prometheus_bash:
	docker exec -it $(prometheus_container) sh
#=================================
# Info for App

.PHONY: info_domen
info_domen:
	echo '---------------------------------';
	echo '----------DEV--------------------';
	echo http://127.0.0.1:8080/swagger/index.html
	echo http://127.0.0.1:8080/metrics
	echo MAILER http://192.168.141.1:8025
	echo PROMETHEUS - http://localhost:9090
	echo GRAFANA - http://localhost:3000
	echo '---------------------------------';
.DEFAULT_GOAL := init
