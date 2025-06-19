package main

import (
	"context"
	"log"
	"os"

	"github.com/T2Knock/blog-aggregators/internal/config"
	"github.com/jackc/pgx/v5"
)

type state struct {
	config config.Config
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Requires at least two arguments")
	}

	currentConfig, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("DB connection failed")
	}
	defer conn.Close(ctx)

	s := &state{
		config: currentConfig,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	cmd := command{
		Name:      os.Args[1],
		Arguments: os.Args[2:],
	}

	if err = cmds.run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
