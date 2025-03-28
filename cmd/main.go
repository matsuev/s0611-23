package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Config ..
type Config struct {
	DBUser string `env:"DB_USER"`
	DBPass string `env:"DB_PASS"`
	DBName string `env:"DB_NAME"`
	DBAddr string `env:"DB_ADDR"`
}

func main() {
	var cfg Config

	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(cfg)

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?%s",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBAddr,
		cfg.DBName,
		"sslmode=disable",
	)

	fmt.Println(connString)

	connCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalln(err)
	}

	connCfg.MaxConns = 70
	connCfg.MinConns = 25

	conn, err := pgxpool.NewWithConfig(context.Background(), connCfg)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	for range 500 {
		go func() {
			if err := conn.Ping(context.Background()); err != nil {
				log.Println(err)
			}
		}()
	}

	var s string

	fmt.Scanln(&s)
}
