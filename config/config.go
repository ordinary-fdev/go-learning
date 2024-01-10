package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server string `yaml:"Server"`
	User   string `yaml:"User"`
	Pwd    string `yaml:"Pwd"`
	Port   string `yaml:"Port"`
	Db     string `yaml:"Db"`
}

func ReadConfig() *Config {

	var config = &Config{}
	file, err := ioutil.ReadFile("./config.yaml")

	if err != nil {
		fmt.Println(err)
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config
}
