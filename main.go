package main

import (
	"context"
	"log"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
)

type Config struct {
	Host string `config:"host"`
	Port uint32 `config:"port"`
}

func main() {
	loader := confita.NewLoader(
		env.NewBackend(),
		flags.NewBackend(),
		file.NewOptionalBackend("config.json"),
		file.NewOptionalBackend("config.yaml"),
	)

	cfg := Config{
		Host: "127.0.0.1",
		Port: 5656,
	}

	if err := loader.Load(context.Background(), &cfg); err != nil {
		log.Fatalln(err)
	}

	log.Println(cfg.Host)
	log.Println(cfg.Port)
}
