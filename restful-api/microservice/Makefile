SHELL=/bin/bash
fmt:
	gofmt -w -s .

run:
	docker-compose up

build:
	docker-compose build

image-build:
	docker build -t redopsbay/sample-restful-api:product-service-v1.0.0 product-service/ && \
	docker build -t redopsbay/sample-restful-api:review-service-v1.0.0 review-service/ && \
	docker build -t redopsbay/sample-restful-api:ratings-service-v1.0.0 ratings-service/ && \
	docker build -t redopsbay/sample-restful-api:loadgen-v1.0.0 load-gen/

image-push:
	docker push redopsbay/sample-restful-api:product-service-v1.0.0 && \
	docker push redopsbay/sample-restful-api:review-service-v1.0.0 && \
	docker push redopsbay/sample-restful-api:ratings-service-v1.0.0 && \
	docker push redopsbay/sample-restful-api:loadgen-v1.0.0