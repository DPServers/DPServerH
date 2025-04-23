package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// 设置路由
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// 获取端口，Vercel会提供PORT环境变量
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 启动服务器
	fmt.Printf("服务器启动在 http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>欢迎来到我的网站</h1>")
	fmt.Fprintf(w, "<p>这是一个使用Go语言构建的网站</p>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>关于我们</h1>")
	fmt.Fprintf(w, "<p>这是一个示例网站，用于展示Go语言的Web开发能力</p>")
}
