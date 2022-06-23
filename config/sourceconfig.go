package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func (c *Config) LoadSourceConfig(path string) (SourceConfig, error) {
	sourceConfigFile, err := ioutil.ReadFile(path)
	sourceConfig := SourceConfig{}
	if err != nil {
		return sourceConfig, err
	}
	err = yaml.UnmarshalStrict(sourceConfigFile, &sourceConfig)
	if err != nil {
		return sourceConfig, err
	}
	err = c.ValidateSourceConfig(sourceConfig)
	return sourceConfig, err
}

func (c *Config) ValidateSourceConfig(cfg SourceConfig) error {
	// TODO: add source config validation
	return nil
}
