package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const configFilePath = "libdyson/config.yml"

// init sets up the file path for config, panics if there are any problems trying to do so
func init() {
	confDir, err := os.UserConfigDir()
	if err != nil {
		panic(fmt.Errorf("could not get user config dir: %w", err))
	}
	fullFilePath = filepath.Clean(fmt.Sprintf("%s/%s", confDir, configFilePath))

	// Make sure the directory is created. Return any error other than an "already exists" error
	err = os.MkdirAll(filepath.Dir(fullFilePath), os.ModePerm)
	if err != nil && !errors.Is(err, os.ErrExist) {
		panic(fmt.Errorf("could not create config dir: %w", err))
	}

	// Stat the file, if it doesn't exist try to create it
	_, err = os.Stat(fullFilePath)
	if errors.Is(err, os.ErrNotExist) {
		_, err = os.Create(fullFilePath)
	}

	if err != nil {
		panic(fmt.Errorf("problem with config file: %w", err))
	}
}

var fullFilePath string

type Config struct {
	Token string

	Devices []Device
}

type Device struct {
	Serial string
}

// writeConfig is a variable so it can be replaced with a mock in unit tests
var writeConfig = func(config Config) error {
	file, err := os.Create(fullFilePath)
	if err != nil {
		return fmt.Errorf("unable to open config file for writing: %w", err)
	}

	err = json.NewEncoder(file).Encode(config)
	if err != nil {
		err = fmt.Errorf("unable to parse config file: %w", err)
	}

	return err
}

// readConfig is a variable so it can be replaced with a mock in unit tests
var readConfig = func() (Config, error) {
	conf := Config{}

	f, err := os.Open(fullFilePath)
	if err != nil {
		return conf, fmt.Errorf("unable to open config file for reading: %w", err)
	}

	err = json.NewDecoder(f).Decode(&conf)
	if err != nil {
		err = fmt.Errorf("unable to parse config file: %w", err)
	}

	return conf, nil
}

func GetFilePath() string {
	return fullFilePath
}
