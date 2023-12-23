SHELL=/bin/bash
DOCKERIMAGE="ratings:latest"
run:
	go mod tidy && \
	export DATABASE_USER=admin && \
	export DATABASE_PASSWORD=admin && \
	export DATABASE_HOSTNAME=172.17.0.1 && \
	export DATABASE_PORT=3306 && \
	export DATABASE_NAME=microservice && \
	export SERVICE_NAME=ratings-service && \
	export OTEL_EXPORTER_OTLP_ENDPOINT=172.26.126.224:4317 && \
	export INSECURE_MODE=true && \
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
	--env SERVICE_NAME=ratings-service && \
	--env OTEL_EXPORTER_OTLP_ENDPOINT=172.26.126.224:4317 && \
	--env INSECURE_MODE=true && \

	ratings:latest