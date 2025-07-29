package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Liveness probe
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// トップページにもハンドラを追加
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to jpgg-backend!"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // デフォルト値
	}

	// 重要: 0.0.0.0 で公開しないと Docker/WSL2 から外に出ない
	addr := "0.0.0.0:" + port
	log.Printf("starting server on %s", addr)

	if err := r.Run(addr); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
