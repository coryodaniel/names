DOCKERFILE=./Dockerfile
REPO_ROOT=quay.io/coryodaniel
IMAGE_NAME=names:latest
IMAGE_URL=${REPO_ROOT}/${IMAGE_NAME}
NAMES_PORT=5050

.PHONY: help
help: ## Show this help
help:
	@grep -E '^[a-zA-Z0-9._-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all
all: names publish

.PHONY: start
start: names
	./names

names:
	go build

.PHONY: build
build:
	docker build . -f ${DOCKERFILE} -t ${IMAGE_NAME} -t ${IMAGE_URL}

.PHONY: publish
publish:
	docker push ${IMAGE_URL}

.PHONY: run
run:
	docker run -e NAMES_PORT=${NAMES_PORT} -p ${NAMES_PORT}:${NAMES_PORT} ${IMAGE_NAME}

.PHONY: deploy
deploy:
	kubectl apply -f ./deploy.yaml

.PHONY: proxy
proxy:
	kubectl port-forward -n default $(shell kubectl get pod -n default --selector="app=names" --output jsonpath='{.items[0].metadata.name}') ${NAMES_PORT}
