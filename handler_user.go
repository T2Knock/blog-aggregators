package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/T2Knock/blog-aggregators/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("missing arguments on command %s <name>", cmd.Name)
	}

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, cmd.Arguments[0])
	if err != nil {
		return err
	}

	if err = s.config.SetCurrentUser(user.Name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User %s switched successfully!\n", user.Name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("missing arguments: usage %s <name>", cmd.Name)
	}

	ctx := context.Background()
	name := cmd.Arguments[0]

	existUser, err := s.db.GetUser(ctx, name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to query user: %w", err)
	}

	if err == nil {
		return fmt.Errorf("user %q already exists", existUser.Name)
	}

	newUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		UserID: uuid.New(),
		Name:   name,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := s.config.SetCurrentUser(newUser.Name); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("New user registered!")
	return nil
}

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()

	err := s.db.DeleteUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete all users: %w", err)
	}

	fmt.Println("User table reset!")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	ctx := context.Background()

	users, err := s.db.GetUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}

	for _, user := range users {
		indicator := ""

		if user.Name == s.config.CurrentUserName {
			indicator = " (current)"
		}

		fmt.Printf("* %s%s\n", user.Name, indicator)
	}

	return nil
}
