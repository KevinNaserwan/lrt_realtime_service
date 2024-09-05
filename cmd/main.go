package main

import (
	"os"

	App "github.com/kevinnaserwan/lrt-realtime-service/internal/app"
	"github.com/kevinnaserwan/lrt-realtime-service/internal/util/env"
)

func main() {
	// Digunakan untuk mengambil konfigurasi dari environment variable
	config := env.LoadConfig()

	// Set timezone ke Asia/Jakarta
	os.Setenv("TZ", "Asia/Jakarta")

	// Inisialisasi aplikasi
	app := App.NewApp(config)

	// Start server
	app.StartServer()
}
