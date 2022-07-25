/*This file represents .automationthingy.yaml*/
package config

import (
	"encoding/json"
)

func (c *Config) GetConfigAsJson() (string, error) {
	// returns the currently loaded config as a JSON string
	c.Logger.Debugf("Converting config into JSON string")
	configJson, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return "", err
	}
	return string(configJson), nil
}
