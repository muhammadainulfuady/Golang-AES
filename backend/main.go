package main

import (
	"crypto-backend/internal/handler"
	"crypto-backend/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi router Gin
	r := gin.Default()

	// Pasang middleware CORS agar frontend Svelte bisa mengakses API
	r.Use(middleware.CORSMiddleware())

	// Kelompokkan route di bawah /api/crypto
	api := r.Group("/api/crypto")
	{
		api.POST("/generate-key", handler.GenerateKeyHandler)
		api.POST("/encrypt", handler.EncryptHandler)
		api.POST("/decrypt", handler.DecryptHandler)
	}

	log.Println("Server backend berjalan di port :8080 (http://localhost:8080)")

	// Jalankan server di port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
