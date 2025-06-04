# Go 学习项目 Makefile
.PHONY: run fmt check clean install-tools

# 运行程序
run:
	@echo "🚀 运行程序..."
	go run .

# 格式化代码
fmt:
	@echo "📝 格式化代码..."
	go fmt ./...
	@echo "✅ 格式化完成!"

# 检查代码
check:
	@echo "🔍 检查代码..."
	go vet ./...
	@echo "✅ 检查完成!"

# 完整的格式化和检查
format-all: fmt check
	@echo "✨ 完整格式化和检查完成!"

# 清理构建文件
clean:
	@echo "🧹 清理构建文件..."
	go clean
	@echo "✅ 清理完成!"

# 安装必要的工具
install-tools:
	@echo "🔧 安装 Go 开发工具..."
	go install golang.org/x/tools/cmd/goimports@latest
	go install golang.org/x/lint/golint@latest
	go install golang.org/x/tools/gopls@latest
	@echo "✅ 工具安装完成!"

# 快速开发循环：格式化 + 运行
dev: fmt run

# 帮助信息
help:
	@echo "可用命令:"
	@echo "  run         - 运行程序"
	@echo "  fmt         - 格式化代码"
	@echo "  check       - 检查代码"
	@echo "  format-all  - 完整格式化和检查"
	@echo "  clean       - 清理构建文件"
	@echo "  install-tools - 安装开发工具"
	@echo "  dev         - 快速开发（格式化+运行）"
	@echo "  help        - 显示帮助信息"
