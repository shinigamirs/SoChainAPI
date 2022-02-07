SHELL = bash
COMPOSE_FILE_RUNTIME = docker-compose.yml

build:
	 docker build -t nuri-sochainapi .

up:
	docker-compose -f ${COMPOSE_FILE_RUNTIME} up -d --remove-orphans

down:
	docker-compose -f ${COMPOSE_FILE_RUNTIME} down --remove-orphans