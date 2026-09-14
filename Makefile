GO      ?= go
NPM     ?= npm
TOOLDIR := tools/config-manager
BIN     ?= $(TOOLDIR)/bin/config-manager
ARGS    ?=

.DEFAULT_GOAL := help
.PHONY: all install dev preview web-build build tool-build tool-build-linux run test vet fmt fmt-check tidy clean help

all: build

# ---------------------------------------------------------------- 前端

install: ## 安装前端依赖
	$(NPM) install

dev: ## 启动前端开发服务器（http://localhost:5173）
	$(NPM) run dev

preview: ## 预览前端构建结果
	$(NPM) run preview

web-build: ## 构建前端到 dist/
	$(NPM) run build

# ---------------------------------------------------------------- 配置管理 TUI

tool-build: ## 编译配置管理 TUI 到 tools/config-manager/bin/
	$(GO) -C $(TOOLDIR) build -o bin/config-manager .

tool-build-linux: ## 交叉编译配置管理 TUI 到 linux/amd64
	GOOS=linux GOARCH=amd64 $(GO) -C $(TOOLDIR) build -o bin/config-manager-linux-amd64 .

run: tool-build ## 运行 TUI，可传参数：make run ARGS="-web dist"
	$(BIN) $(ARGS)

test: ## 运行 Go 程序全部测试
	$(GO) -C $(TOOLDIR) test ./...

vet: ## Go 静态检查
	$(GO) -C $(TOOLDIR) vet ./...

fmt: ## 格式化 Go 代码
	gofmt -w $(TOOLDIR)

fmt-check: ## 检查 Go 代码格式
	@unformatted=$$(gofmt -l $(TOOLDIR)); \
	if [ -n "$$unformatted" ]; then echo "以下文件需要 gofmt:"; echo "$$unformatted"; exit 1; fi

tidy: ## 整理 Go 模块依赖
	$(GO) -C $(TOOLDIR) mod tidy

# ---------------------------------------------------------------- 其它

build: web-build tool-build ## 构建前端 + 配置管理 TUI

clean: ## 清理 dist/ 与 TUI 构建产物
	rm -rf dist $(TOOLDIR)/bin

help: ## 显示本帮助
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "} {printf "  \033[36m%-12s\033[0m %s\n", $$1, $$2}'
