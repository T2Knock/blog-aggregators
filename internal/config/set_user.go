package config

import (
	"encoding/json"
	"os"
)

func (c *Config) SetCurrentUser(username string) error {
	c.CurrentUserName = username

	jsonData, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	if err = os.WriteFile(filePath, jsonData, 0644); err != nil {
		return err
	}

	return nil
}
