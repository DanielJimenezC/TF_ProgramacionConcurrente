package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Psql configuration struct
type Psql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// SetConfiguration PSQL enviroment
func (config *Psql) SetConfiguration() {
	file, err := ioutil.ReadFile("Configuration.yaml")
	if err != nil {
		log.Printf("File ger err %v", err)
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	os.Setenv("POSTGRES_HOST", config.Host)
	os.Setenv("POSTGRES_PORT", config.Port)
	os.Setenv("POSTGRES_USER", config.User)
	os.Setenv("POSTGRES_PASSWORD", config.Password)
	os.Setenv("POSTGRES_DATABASE", config.Database)
}
