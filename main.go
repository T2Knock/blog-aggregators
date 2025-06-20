package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/T2Knock/blog-aggregators/internal/config"
	"github.com/T2Knock/blog-aggregators/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
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

	db, err := sql.Open("postgres", currentConfig.DbURL)
	if err != nil {
		log.Fatalf("DB connection failed %v", err)
	}

	dbQuerries := database.New(db)

	s := &state{
		config: currentConfig,
		db:     dbQuerries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	cmd := command{
		Name:      os.Args[1],
		Arguments: os.Args[2:],
	}

	if err = cmds.run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
