.SILENT:
.PHONY:
#=================================
# Run service

run_gateway:
	go run gateway/cmd/main.go -config=./gateway/config/config.yaml

run_reader:
	go run reader_service/cmd/main.go -config=./reader_service/config/config.yaml

run_writer:
	go run writer_service/cmd/main.go -config=./writer_service/config/config.yaml

run_library:
	go run library-app/cmd/app/main.go

#=================================
# Command for docker

up: up_docker

info: ps info_domen

up_docker:
	docker-compose up

up_docker_d:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

rebuild: down build up_docker info

# флаг -v удаляет все volume (очищает все данные)
down-clear:
	docker-compose down -v --remove-orphans

build:
	docker-compose build

ps:
	docker-compose ps

# ================================
# Modules support

tidy:
	go mod tidy

deps-reset:
	git checkout -- go.mod
	go mod tidy

deps-upgrade:
	go get -u -t -d -v ./...
	go mod tidy

deps-cleancache:
	go clean -modcache

#=================================
# Swagger

swagger:
	@echo Starting swagger generating
	swag init -g **/**/*.go
# ================================
# PPROF

pprof_heap:
	go tool pprof -http :8006 http://localhost:6060/debug/pprof/heap?seconds=10

pprof_cpu:
	go tool pprof -http :8006 http://localhost:6060/debug/pprof/profile?seconds=10

pprof_allocs:
	go tool pprof -http :8006 http://localhost:6060/debug/pprof/allocs?seconds=10
# ================================
# MongoDB

mongo:
	cd ./scripts && mongo admin -u admin -p admin < mongo_init.js
# ================================
# Go migrate postgresql https://github.com/golang-migrate/migrate

DB_NAME = products
DB_HOST = localhost
DB_PORT = 5432
SSL_MODE = disable

force_db:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path ./migrations force 1

version_db:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path ./migrations version

migrate_up:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path ./migrations up

migrate_down:
	migrate -database postgres://postgres:postgres@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE) -path ./migrations down
# ================================
# Linters https://golangci-lint.run/usage/install/

run-linter:
	@echo Starting linters
	golangci-lint run ./...
# ================================
# Proto

proto_kafka:
	@echo Generating kafka proto
	cd proto/kafka && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. kafka.proto

proto_reader:
	@echo Generating product reader microservice proto
	cd reader_service/proto/product_reader && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. product_reader.proto

proto_reader_message:
	@echo Generating product reader messages microservice proto
	cd reader_service/proto/product_reader && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. product_reader_messages.proto

proto_writer:
	@echo Generating product writer microservice proto
	cd writer_service/proto/product_writer && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. product_writer.proto

proto_writer_message:
	@echo Generating product writer messages microservice proto
	cd writer_service/proto/product_writer && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. product_writer_messages.proto

#=================================
# Info for App

info_domen:
	echo '---------------------------------';
	echo '----------DEV--------------------';
	echo MAILER - http://192.168.141.1:8025
	echo FTP-SERVER - http://192.168.141.1:8081
	echo MINIO - 192.168.141.1:9000
	echo JAEGER - http://localhost:16686
	echo SWAGGER - http://127.0.0.1:5001/swagger/index.html
	echo METRICS - http://127.0.0.1:8001/metrics
	echo KAFKA UI - http://127.0.0.1:9000
	echo Prometheus UI - http://localhost:9090
	echo Grafana UI - http://localhost:3005
	echo '---------------------------------';
.DEFAULT_GOAL := init
