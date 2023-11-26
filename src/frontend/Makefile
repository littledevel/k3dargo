FRONTEND_TAG ?= 1.0

build:
	docker build -t frontend:$(FRONTEND_TAG) .

push:
	docker tag frontend brokenleg/frontend:$(FRONTEND_TAG)
	docker push brokenleg/frontend:$(FRONTEND_TAG)