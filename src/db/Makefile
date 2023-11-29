DB_TAG ?= 3.0
REGISTRY_URL ?= brokenleg

build:
	docker build -t db:$(DB_TAG) .

push: build
	docker tag db:$(DB_TAG) $(REGISTRY_URL)/db:$(DB_TAG)
	docker push $(REGISTRY_URL)/db:$(DB_TAG)
