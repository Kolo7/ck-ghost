PROJECT="ck-ghost"
BUILDTIME=`date '+%Y%m%d%H%M'`

run: ## run --help
	go run main.go -h


build: ## Build by linux
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/${PROJECT}.${BUILDTIME} main.go

help: ## Display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

.SILENT: build help

.PHONY: all build
