include .env

stop_containers:
	@echo "Stopping other docker container"
	@if [ $$(docker ps -q) ]; then \
		echo "found and stopped containers"; \
		docker stop $$(docker ps -q); \
	else \
		echo "no containers running..."; \
	fi

create_container:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres:16-alpine

start_container:
	docker start ${DB_DOCKER_CONTAINER}

create_db:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

drop_db:
	docker exec -it ${DB_DOCKER_CONTAINER} dropdb --username=${DB_USER} ${DB_NAME}

build_app:
	GOOS=linux GOARCH=amd64 go build -o webapp
