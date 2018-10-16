package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Counter struct {
	Value       int
	DisplayName string
	Regex       string
}

//Config is the struct of the config
type Config struct {
	Counters []Counter
}

// NewConfig creates config with the config file path
func NewConfig(configFile string) *Config {
	var config Config
	str, _ := ioutil.ReadFile(configFile)
	if err := json.Unmarshal(str, &config); err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	return &config
}
