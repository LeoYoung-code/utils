#!/bin/bash

# 迁移工具库文件到新的目录结构
# 作者：Claude
# 日期：2024-07-12

echo "开始迁移文件到新的目录结构..."

# 确保目标目录存在
mkdir -p pkg/algorithm/{set,map} \
         pkg/aigc \
         pkg/cache/bigcache \
         pkg/common/{trace,gc,addr,template,error} \
         pkg/concurrent/{pool,atomic,async,channel} \
         pkg/crypto \
         pkg/database/sql \
         pkg/design/generate_builder \
         pkg/encoding/{json,transcode} \
         pkg/http \
         pkg/io/pdf \
         pkg/math \
         pkg/string \
         pkg/test \
         pkg/time \
         pkg/version \
         cmd \
         internal/example

# 复制文件时检查源文件是否存在
safe_copy() {
    if [ -f "$1" ]; then
        mkdir -p "$(dirname "$2")"
        cp -v "$1" "$2"
    else
        echo "警告: 源文件 $1 不存在，跳过"
    fi
}

safe_copy_dir() {
    if [ -d "$1" ]; then
        mkdir -p "$2"
        cp -rv "$1"/* "$2"/ 2>/dev/null || true
    else
        echo "警告: 源目录 $1 不存在，跳过"
    fi
}

# 移动根目录下的通用工具函数到 pkg/common
echo "移动通用工具函数..."
safe_copy util.go pkg/common/util.go
safe_copy util_test.go pkg/common/util_test.go
safe_copy gen_uid.go pkg/common/gen_uid.go
safe_copy gen_uid_test.go pkg/common/gen_uid_test.go
safe_copy transaction.go pkg/database/transaction.go

# 移动设计模式相关文件
echo "移动设计模式相关文件..."
safe_copy desgin/visitor.go pkg/design/visitor.go
safe_copy_dir desgin/generate_builder pkg/design/generate_builder

# 移动各个功能模块
echo "移动各个功能模块..."
safe_copy_dir http pkg/http
safe_copy_dir aigc pkg/aigc
safe_copy_dir encrypt pkg/crypto
safe_copy_dir cache pkg/cache
safe_copy_dir time pkg/time
safe_copy_dir string pkg/string
safe_copy_dir math pkg/math
safe_copy_dir version pkg/version
safe_copy_dir test pkg/test
safe_copy_dir json pkg/encoding/json
safe_copy_dir bigcache pkg/cache/bigcache
safe_copy_dir pool pkg/concurrent/pool
safe_copy_dir atomic pkg/concurrent/atomic
safe_copy_dir async pkg/concurrent/async
safe_copy_dir channel pkg/concurrent/channel
safe_copy_dir trace pkg/common/trace
safe_copy_dir common pkg/common
safe_copy_dir sql pkg/database/sql
safe_copy_dir gc pkg/common/gc
safe_copy_dir set pkg/algorithm/set
safe_copy_dir map pkg/algorithm/map
safe_copy_dir addr pkg/common/addr
safe_copy_dir transcode pkg/encoding/transcode
safe_copy_dir pdf pkg/io/pdf
safe_copy_dir template pkg/common/template
safe_copy_dir error pkg/common/error
safe_copy_dir main cmd

echo "迁移完成！"
echo "请检查新结构下的代码是否正常工作，然后考虑是否删除旧的文件和目录。"
echo "可以执行以下命令编译项目检查是否有问题："
echo "go mod tidy && go build -o /tmp/utils_test ./cmd/..." 