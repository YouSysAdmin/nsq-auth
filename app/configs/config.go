package configs

import (
	"fmt"
	"github.com/yousysadmin/nsq-auth/app/models"
	"gopkg.in/yaml.v2"
	"os"
)

type Identity struct {
	Identity       string                 `json:"identity" yaml:"identity"`
	Secret         string                 `json:"secret" yaml:"secret"`
	Authorizations []models.Authorization `json:"authorizations" yaml:"authorizations"`
}

// Config is the application config structure
type Config struct {
	// IP address of the HTTP server binding
	BindIP string `json:"bind_addr" yaml:"bind_addr"`
	// HTTP port of the HTTP server binding
	BindPort string `json:"bind_port" yaml:"bind_port"`
	// Full bid address of the HTTP server
	BindAddr string `json:"_" yaml:"_"`
	// Identity list
	Identities []Identity `json:"identities" yaml:"identities"`
}

// Load load HSQAuth config file
func (c *Config) Load(configFilePath string) (*Config, error) {
	yamlFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}
	// setup full address of the HTTPS server binding
	c.BindAddr = c.MakeBindAddr()

	return c, err
}

// FindIdentityBySecret find identity by `secret` field
func (c *Config) FindIdentityBySecret(secret string) (*Identity, error) {
	for _, ident := range c.Identities {
		if ident.Secret == secret {
			return &ident, nil
		}
	}
	return &Identity{}, fmt.Errorf("identity with secret '%s' not found in the configs file", secret)
}

// MakeBindAddr preparing the HTTP server binding address from the BindIP and BindPort fields
func (c *Config) MakeBindAddr() string {
	if c.BindPort == "" {
		c.BindPort = "4181"
	}

	if c.BindIP == "" {
		c.BindIP = "0.0.0.0"
	}

	return fmt.Sprintf("%s:%s", c.BindIP, c.BindPort)
}
