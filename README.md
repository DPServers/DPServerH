# Go Web 项目

这是一个使用Go语言构建的简单Web服务器项目。

## 功能特点

- 使用Go标准库`net/http`构建
- 支持多个路由处理
- 简单的HTML页面展示
- 支持Vercel部署

## 本地运行方法

1. 确保已安装Go环境
2. 在项目目录下运行：
   ```bash
   go run main.go
   ```
3. 在浏览器中访问：http://localhost:8080

## Vercel部署

1. 将代码推送到GitHub仓库
2. 访问 [Vercel官网](https://vercel.com)
3. 使用GitHub账号登录
4. 点击"New Project"
5. 选择您的GitHub仓库
6. 在配置页面：
   - Framework Preset: 选择"Other"
   - Build Command: 留空
   - Output Directory: 留空
7. 点击"Deploy"

部署完成后，Vercel会提供一个域名，您可以通过该域名访问您的网站。

## 项目结构

- `main.go`: 主程序文件，包含服务器配置和路由处理
- `vercel.json`: Vercel部署配置文件
- `README.md`: 项目说明文档

## 路由说明

- `/`: 首页
- `/about`: 关于页面 