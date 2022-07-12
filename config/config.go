/*This file represents .automationthingy.yaml*/
package config

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Config struct {
	Rbac        Rbac       `json:"rbac"`
	Scripts     []Script   `json:"scripts"`
	Sources     Sources    `json:"sources"`
	Secretmgrs  SecretMgrs `json:"secretmgrs"`
	Logger      Logger
	ScriptIndex map[uuid.UUID]int
}

type SourceConfig struct {
	Scripts []Script `json:"scripts"`
}

func (c *Config) GetConfigAsJson() (string, error) {
	// returns the currently loaded config as a JSON string
	c.Logger.Debugf("Converting config into JSON string")
	configJson, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return "", err
	}
	return string(configJson), nil
}
