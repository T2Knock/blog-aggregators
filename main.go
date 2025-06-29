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

	currentConfig, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", currentConfig.DbURL)
	if err != nil {
		log.Fatalf("DB connection failed %v", err)
	}

	dbQueries := database.New(db)

	s := &state{
		config: currentConfig,
		db:     dbQueries,
	}

	cmds := newCommands()

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	cmd := command{
		Name:      os.Args[1],
		Arguments: os.Args[2:],
	}

	if err := cmds.run(s, cmd); err != nil {
		log.Fatal(err)
	}
}
