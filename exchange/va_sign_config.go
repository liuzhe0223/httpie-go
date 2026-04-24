package exchange

import (
	"fmt"
	"os"

	"github.com/liuzhe0223/vasign"
	"gopkg.in/yaml.v2"
)

type VASignConfig struct {
	ClientID   string `yaml:"client_id"`
	KeyID      string `yaml:"key_id"`
	PrivateKey string `yaml:"private_key"`
}

func LoadVASigner(path string) (*vasign.Signer, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read va sign config: %w", err)
	}

	var cfg VASignConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parse va sign config: %w", err)
	}

	signer, err := vasign.NewSignerFromBase64(cfg.ClientID, cfg.KeyID, cfg.PrivateKey)
	if err != nil {
		return nil, err
	}
	return signer, nil
}
