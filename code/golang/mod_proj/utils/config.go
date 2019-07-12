package utils

import (
	"fmt"
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"

	"io/ioutil"
)

type Config struct {
	Mysql struct {
		Address  string `yaml:"Address"`
		Port     string `yaml:"Port"`
		DB       string `yaml:"DB"`
		User     string `yaml:"User"`
		Password string `yaml:"Password"`
	}
}

func GetConfig() Config {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cfgPath := path.Join(cwd, "conf/config.yml")
	fmt.Println(cfgPath)
	return loadConfig(cfgPath)
}

func loadConfig(cfgPath string) Config {
	c := Config{}

	bytes, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
