/*This file represents .automationthingy.yaml*/
package config

import (
	"encoding/json"

	"github.com/peter-mcconnell/automationthingy/types"
)

type Config struct {
	Rbac          types.Rbac          `json:"rbac"`
	Scripts       []types.ScriptData  `json:"scripts"`
	Scriptsources types.ScriptSources `json:"scriptsources"`
	Secretmgrs    types.SecretMgrs    `json:"secretmgrs"`
	Logger        types.Logger
}

/**
sourcescripts:
  # python example
  - id: "268c55ac-6e2b-4c99-b84c-535b7d7e6cbc"
    name: "some random ops job 1"
    desc: "this is a little script"
    workdir: "./mythingy/"
    command: |-
      python3 main.py
    categories:
      - "production ops / testing"
      - "admin stuff"
*/

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
