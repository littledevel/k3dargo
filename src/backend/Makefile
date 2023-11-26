BACKEND_TAG ?= 2.1

build:
	docker build -t backend:$(BACKEND_TAG) .

push: build
	docker tag backend:$(BACKEND_TAG) brokenleg/backend:$(BACKEND_TAG)
	docker push brokenleg/backend:$(BACKEND_TAG)
