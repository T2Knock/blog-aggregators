package main

import (
	"fmt"
	"log"
	"os"

	"github.com/T2Knock/blog-aggregators/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	handlerMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlerMap[cmd.name]
	if !ok {
		return fmt.Errorf("command %v did not exist", cmd.name)
	}

	handler(s, cmd)

	return nil
}

func (c *commands) register(name string, handler func(*state, command) error) {
	c.handlerMap[name] = handler
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("missing arguments on command %v", cmd.name)
	}

	s.config.SetUsers(cmd.arguments[0])

	fmt.Println("user login success!")

	return nil
}

func main() {
	arguments := os.Args

	if len(arguments) < 2 {
		log.Fatalf("Requires at least two arguments")
	}

	currentConfig, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	s := state{
		config: &currentConfig,
	}

	cmds := commands{
		handlerMap: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
}
