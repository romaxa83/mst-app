.SILENT:
#=================================
# Command for docker

.PHONY: up
up: up_docker

.PHONY: info
info: ps info_domen

.PHONY: up_docker
up_docker:
	docker-compose up

.PHONY: up_docker_d
up_docker_d:
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
# Info for App

.PHONY: info_domen
info_domen:
	echo '---------------------------------';
	echo '----------DEV--------------------';
	echo MAILER http://192.168.141.1:8025
	echo FTP-SERVER http://192.168.141.1:8081
	echo '---------------------------------';
.DEFAULT_GOAL := init
