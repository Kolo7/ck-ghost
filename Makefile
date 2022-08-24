PROJECT="ck-ghost"
BUILDTIME=`date '+%Y%m%d%H%M'`

run: ## run --help
	go run main.go -h


build_linux: ## Build by linux
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/${PROJECT}.${BUILDTIME}/linux main.go

build: ## Build by local
	go build -o bin/${PROJECT}.${BUILDTIME} main.go

help: ## Display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# 默认目标
.DEFAULT_GOAL := help

# 不打印执行过程
.SILENT: build help

# 忽略目标冲突文件
.PHONY: all build build_linux
