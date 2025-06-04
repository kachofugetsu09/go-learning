#!/bin/bash

# Go 代码格式化和检查脚本
echo "🚀 开始格式化 Go 代码..."

# 1. 格式化所有 Go 文件
echo "📝 格式化代码..."
go fmt ./...

# 2. 整理导入包
echo "📦 整理导入包..."
for file in *.go; do
    if [ -f "$file" ]; then
        goimports -w "$file"
    fi
done

# 3. 运行 go vet 检查
echo "🔍 运行代码检查..."
go vet ./...

# 4. 运行 golint (如果安装了)
if command -v golint &> /dev/null; then
    echo "✨ 运行代码风格检查..."
    golint ./...
fi

echo "✅ 格式化完成!"
