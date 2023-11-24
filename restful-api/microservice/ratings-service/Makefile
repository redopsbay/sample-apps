SHELL=/bin/bash
run: 
	export DATABASE_USER=admin && \
	export DATABASE_PASSWORD=admin && \
	export DATABASE_HOSTNAME=172.17.0.1 && \
	export DATABASE_PORT=3306 && \
	export DATABASE_NAME=microservice && \
	go run main.go

fmt:
	gofmt -w -s .


build:
	gofmt -w -s . && \
	go mod tidy && \
	go build ./main.go && \
	chmod +x main

image-build:
	docker build -t $(DOCKERIMAGE) .


run-docker:
	docker run \
	-it \
	--rm \
	--name server \
	-p 8080:8080 \
	--env DATABASE_USER=admin \
	--env DATABASE_PASSWORD=admin \
	--env DATABASE_HOSTNAME=172.17.0.1 \
	--env DATABASE_PORT=3306 \
	--env DATABASE_NAME=microservice \
	ratings:latest