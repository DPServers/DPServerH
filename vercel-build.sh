#!/bin/bash

# 设置Go环境变量
export GO111MODULE=on
export GOPATH=$PWD

# 构建项目
go build -o vercel-go main.go 