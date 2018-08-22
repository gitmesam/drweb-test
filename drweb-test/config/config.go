package config

import (
	"io/ioutil"
	"log"
	"os"

	"drweb-test/google-client"

	"gopkg.in/yaml.v2"
)

// Config for service
type Config struct {
	Service Service              `yaml:"service"`
	Google  google_client.Config `yaml:"google"`
}

// Service config
type Service struct {
	Port int `yaml:"port"`
}

// ReadConfig reads given yaml and parses it to config
func ReadConfig(path string, config interface{}) {
	if err := ParseYamlFile(path, config); err == nil {
		return
	} else if os.IsNotExist(err) { // if 'name' file exists and cannot be parsed, do not check other locations.
		path = "config/" + path
		err := ParseYamlFile(path, config)
		if err == nil {
			return
		} else if !os.IsNotExist(err) {
			log.Fatalf("Failed to open config '%s': %v", path, err)
		}
	} else {
		log.Fatalf("Failed to open config '%s': %v", path, err)
	}
	log.Fatalf("Unable to find config file '%s'", path)
}

// ParseYamlFile reads file with the given name and parses its content as yaml.
func ParseYamlFile(fullName string, data interface{}) error {
	bytes, err := ioutil.ReadFile(fullName)
	if err == nil {
		err = yaml.Unmarshal(bytes, data)
	}
	return err
}
