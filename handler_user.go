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

	user, err := s.db.GetUser(context.Background(), cmd.Arguments[0])
	if err != nil {
		return err
	}

	if err = s.config.SetUsers(user.UserName); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User %s switched sucessfully!\n", user.UserName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("missing arguments: usage %s <name>", cmd.Name)
	}

	ctx := context.Background()
	userName := cmd.Arguments[0]

	existUser, err := s.db.GetUser(ctx, userName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("failed to query user: %w", err)
	}

	if err == nil {
		return fmt.Errorf("user %q already exist", existUser.UserName)
	}

	newUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		UserID:   uuid.New(),
		UserName: userName,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	if err := s.config.SetUsers(newUser.UserName); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("New user registered!")
	return nil
}
