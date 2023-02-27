package main

import (
	"log"

	"github.com/faqihyugos/pengaduan-api/server"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
