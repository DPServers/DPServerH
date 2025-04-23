package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// 设置静态文件服务
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// 获取端口
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 启动服务器
	log.Printf("服务器启动在 http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
