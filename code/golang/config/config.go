package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/olebedev/config"
)

// Config struct
type Config struct {
	cfg *config.Config
}

// Get (path string) string
func (cfg Config) Get(path string) string {
	v, err := cfg.cfg.String(path)
	if err != nil {
		log.Printf("Config.Get failed, cause: %v\n", err)
		return ""
	}
	return v
}

// GetConfig (fpath string) (*Config, error)
func GetConfig(fpath string) (*Config, error) {
	absPath, err := filepath.Abs(fpath)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	yamlString := string(bytes)
	cfg, err := config.ParseYaml(yamlString)

	if err != nil {
		return nil, err
	}
	return &Config{cfg}, nil
}
