package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 {
		return fmt.Errorf("missing arguments on command %s <name>", cmd.Name)
	}

	err := s.config.SetUsers(cmd.Arguments[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched sucessfully!")
	return nil
}
