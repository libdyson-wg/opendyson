package config

import "fmt"

func SetToken(t string) error {
	conf, err := readConfig()
	if err != nil {
		return fmt.Errorf("could not read config: %v", err)
	}

	conf.Token = t
	err = writeConfig(conf)
	if err != nil {
		err = fmt.Errorf("could not save config: %v", err)
	}

	return err
}

func GetToken() (string, error) {
	conf, err := readConfig()
	if err != nil {
		return "", fmt.Errorf("could not read token from config: %v", err)
	}

	return conf.Token, nil
}
