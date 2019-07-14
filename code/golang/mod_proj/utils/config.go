package utils

import (
	"log"
	"os"
	"path"

	"gopkg.in/yaml.v2"

	"io/ioutil"
)

type Config struct {
	Mysql struct {
		Address      string `yaml:"Address"`
		Port         int    `yaml:"Port"`
		DB           string `yaml:"DB"`
		User         string `yaml:"User"`
		Password     string `yaml:"Password"`
		MaxIdleConns int    `yaml:"MaxIdleConns"`
		MaxOpenConns int    `yaml:"MaxOpenConns"`
	} `yaml:"Mysql"`

	Redis struct {
		Address string `yaml:"Address"`
		Port    string `yaml:"Port"`
	} `yaml:"Redis"`
}

func GetConfig() Config {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cfgPath := path.Join(cwd, "conf/config.yml")
	return loadConfig(cfgPath)
}

func loadConfig(cfgPath string) (c Config) {
	bytes, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.UnmarshalStrict(bytes, &c)
	if err != nil {
		log.Fatal(err)
	}
	return
}
