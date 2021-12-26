.SILENT:
#=================================
# Command for App

.PHONY: info
info: ps info_domen

.PHONY: ps
ps:
	docker-compose ps

# run app
.PHONY: run
run:
	go run cmd/app/main.go

# generate swagger documentation
.PHONY: swagger
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
