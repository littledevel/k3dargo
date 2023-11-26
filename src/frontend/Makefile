FRONTEND_TAG ?= 1.1

build:
	docker build -t frontend:$(FRONTEND_TAG) .

push: build
	docker tag frontend brokenleg/frontend:$(FRONTEND_TAG)
	docker push brokenleg/frontend:$(FRONTEND_TAG)
