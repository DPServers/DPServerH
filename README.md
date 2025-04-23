# Go Web 静态页面示例

这是一个使用Go语言提供静态页面的简单示例。

## 功能特点

- 使用Go标准库`net/http`提供静态文件服务
- 简单的Hello World静态页面
- 响应式设计，适配各种设备

## 项目结构

```
go_web/
├── main.go          # Go服务器主程序
├── static/          # 静态文件目录
│   └── index.html   # 静态页面
└── README.md        # 项目说明文档
```

## 本地运行方法

1. 确保已安装Go环境
2. 在项目目录下运行：
   ```bash
   go run main.go
   ```
3. 在浏览器中访问：http://localhost:8080

## 部署说明

本项目可以部署到任何支持Go的服务器或平台，如：
- Vercel
- Heroku
- 普通Linux服务器

部署时需要注意：
1. 确保`static`目录被正确部署
2. 设置正确的环境变量`PORT`
3. 确保服务器有足够的权限访问静态文件 