.SILENT:
.PHONY:
#=================================
# Command for App

info: ps info_domen

ps:
	docker-compose ps

# run app
run:
	go run cmd/app/main.go

#=================================
# Generate

# proto
gen_proto: proto_author proto_author_message

proto_author:
	@echo Generating library-author microservice proto
	cd proto/author && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. author.proto

proto_author_message:
	@echo Generating library-author messages microservice proto
	cd proto/author && protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. author_messages.proto

swagger:
	swag init -g internal/app/app.go
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
# Info for App

.PHONY: info_domen
info_domen:
	echo '---------------------------------';
	echo '----------DEV--------------------';
	echo http://127.0.0.1:8060/swagger/index.html
	echo '---------------------------------';
.DEFAULT_GOAL := init
