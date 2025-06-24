package main

import (
	"context"
	"fmt"

	"github.com/T2Knock/blog-aggregators/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("failed fetching user: %w", err)
		}

		return handler(s, cmd, user)
	}
}
