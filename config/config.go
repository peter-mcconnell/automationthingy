/*This file represents .automationthingy.yaml*/
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

func (c *Config) LoadConfig() (Config, error) {
	// load base config
	configFile, err := ioutil.ReadFile(".automationthingy.yaml")
	config := Config{}
	if err != nil {
		return config, err
	}
	err = yaml.UnmarshalStrict(configFile, &config)
	if err != nil {
		return config, err
	}
	err = c.ValidateConfig(config)
	// load script sources
	return config, err
}

func (c *Config) GetConfigAsJson() string {
	// prints the
	configJson, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(configJson)
}

func (c *Config) LoadScriptSources() {
	// if err := scm.CloneScriptRepos("/tmp/x"); err != nil {
	// 	panic(err)
	// }
}

func (c *Config) ValidateConfig(cfg Config) error {
	// validate scripts
	var errs []string
	for _, script := range cfg.Scripts {
		if script.ID.String() == "00000000-0000-0000-0000-000000000000" {
			errs = append(errs, fmt.Sprintf("script %s has no id. please add it.", script.Name))
		}
	}
	if len(errs) != 0 {
		return fmt.Errorf(strings.Join(errs, "\n"))
	}
	return nil
}
