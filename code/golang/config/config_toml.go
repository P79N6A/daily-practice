package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

// TomlConfig struct
type TomlConfig struct {
	Env map[string]EnvConfig `toml:"enviroment"`
}

// EnvConfig struct
type EnvConfig struct {
	Adapter  string `toml:"adapter"`
	Encoding string `toml:"encoding"`
	Database string `toml:"database"`
	Pool     int    `toml:"pool"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
}

var (
	CurrentConfig TomlConfig
)

func init() {
	var err error
	CurrentConfig, err = GetTomlConfig("./config/database.yml")
	if err != nil {
		panic(err)
	}
}

// GetTomlConfig (fpath string) (*TomlConfig, error)
func GetTomlConfig(fpath string) (*TomlConfig, error) {
	absPath, err := filepath.Abs(fpath)
	if err != nil {
		return nil, err
	}
	blob, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	tomlData := string(blob)
	var config TomlConfig
	if _, err := toml.Decode(tomlData, &config); err != nil {
		log.Fatal(err)
	}

	return &config, nil
}
