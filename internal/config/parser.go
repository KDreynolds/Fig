package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ParseConfig reads and parses the YAML configuration file
func ParseConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	if err := validateConfig(&config); err != nil {
		return nil, fmt.Errorf("config validation failed: %v", err)
	}

	return &config, nil
}

// validateConfig performs basic validation on the parsed configuration
func validateConfig(config *Config) error {
	if len(config.ServerGroups) == 0 {
		return fmt.Errorf("no server groups defined")
	}

	if len(config.Tasks) == 0 {
		return fmt.Errorf("no tasks defined")
	}

	if len(config.Configurations) == 0 {
		return fmt.Errorf("no configurations defined")
	}

	// Add more validation as needed
	return nil
}

// GetServerGroup retrieves a server group by name
func (c *Config) GetServerGroup(name string) (*ServerGroup, error) {
	for _, sg := range c.ServerGroups {
		if sg.Name == name {
			return &sg, nil
		}
	}
	return nil, fmt.Errorf("server group '%s' not found", name)
}

// GetTask retrieves a task by name
func (c *Config) GetTask(name string) (*Task, error) {
	for _, t := range c.Tasks {
		if t.Name == name {
			return &t, nil
		}
	}
	return nil, fmt.Errorf("task '%s' not found", name)
}

// GetConfiguration retrieves a configuration by name
func (c *Config) GetConfiguration(name string) (*Configuration, error) {
	for _, conf := range c.Configurations {
		if conf.Name == name {
			return &conf, nil
		}
	}
	return nil, fmt.Errorf("configuration '%s' not found", name)
}
