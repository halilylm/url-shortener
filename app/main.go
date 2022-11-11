package main

import (
	"github.com/halilylm/url-shortner/config"
	"github.com/halilylm/url-shortner/server"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	cfg, err := config.Load()
	if err != nil {
		log.Fatalln(err)
	}
	srv := server.New(cfg)
	srv.Run()
}
