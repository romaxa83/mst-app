.SILENT:
.PHONY:
#=================================
# Run service

run_gateway:
	go run gateway/cmd/main.go -config=./gateway/config/config.yaml

run_reader:
	go run reader_service/cmd/main.go -config=./reader_service/config/config.yaml

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

#=================================
# Swagger

swagger:
	@echo Starting swagger generating
	swag init -g **/**/*.go

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

#=================================
# Info for App

info_domen:
	echo '---------------------------------';
	echo '----------DEV--------------------';
	echo MAILER http://192.168.141.1:8025
	echo FTP-SERVER http://192.168.141.1:8081
	echo MINIO 192.168.141.1:9000
	echo JAEGER http://localhost:16686
	echo SWAGGER http://127.0.0.1:5001/swagger/index.html
	echo METRICS http://http://127.0.0.1:8001/metrics
	echo '---------------------------------';
.DEFAULT_GOAL := init
