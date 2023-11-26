FRONTEND_TAG ?= 1.3

build:
	docker build -t frontend:$(FRONTEND_TAG) .

push: build
	docker tag frontend:$(FRONTEND_TAG) brokenleg/frontend:$(FRONTEND_TAG)
	docker push brokenleg/frontend:$(FRONTEND_TAG)
