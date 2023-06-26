#!/bin/bash

# 定义默认参数-d的值
dirs=(
    "./core"
    "./global/common/response"
    "./global/common/request"
    "./storage/relational/model"
    "./api"
)

# 保存旧的IFS值
OLD_IFS=$IFS
IFS=,

# 执行初始化
~/go/bin/swag init -g server.go -d "${dirs[*]}"

# 恢复
IFS=$OLD_IFS
