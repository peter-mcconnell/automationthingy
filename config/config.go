/*This file represents .automationthingy.yaml*/
package config

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/peter-mcconnell/automationthingy/types"
)

type Config struct {
	Rbac          types.Rbac          `json:"rbac"`
	Scripts       []types.Script      `json:"scripts"`
	Scriptsources types.ScriptSources `json:"scriptsources"`
	Secretmgrs    types.SecretMgrs    `json:"secretmgrs"`
	Logger        types.Logger
	ScriptIndex   map[uuid.UUID]int
}

type SourceConfig struct {
	Sourcescripts []types.SourceScriptData `json:"sourcescripts"`
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
