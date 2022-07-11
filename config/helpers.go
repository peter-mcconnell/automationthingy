package config

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/peter-mcconnell/automationthingy/scm"
	"github.com/peter-mcconnell/automationthingy/types"
	"gopkg.in/yaml.v2"
)

func LoadConfig(logger types.Logger) (Config, error) {
	// load base config
	configFilePath := ".automationthingy.yaml"
	logger.Debugf("Reading config file: %s", configFilePath)
	configFile, err := ioutil.ReadFile(configFilePath)
	config := Config{
		Logger: logger,
	}
	if err != nil {
		return config, err
	}
	err = yaml.UnmarshalStrict(configFile, &config)
	if err != nil {
		return config, err
	}
	if err = LoadScriptSources(logger, config); err != nil {
		return config, err
	}
	logger.Debugf("Validating config file")
	err = ValidateConfig(config)
	// load script sources
	return config, err
}

func ValidateConfig(cfg Config) error {
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

func LoadScriptSources(logger types.Logger, config Config) error {
	logger.Debug("Loading script sources")
	git := scm.Git{
		Logger: logger,
	}
	for _, source := range config.Scriptsources.Git {
		dest := fmt.Sprintf("scripts/%s", base64.StdEncoding.EncodeToString([]byte(source.Repo)))
		// ensure source has been pulled
		if _, err := os.Stat(dest); os.IsNotExist(err) {
			logger.Debugf("%s directory '%s' doesn't exist. Creating it", source.Repo, dest)
			if err := git.Clone(source, dest); err != nil {
				return err
			}
		} else {
			logger.Debugf("%s directory '%s' already exists. Pulling latest", source.Repo, dest)
			logger.Info("TODO - ensure latest?")
		}
		// load source config
		logger.Info("TODO - load source config")
	}
	return nil
}
