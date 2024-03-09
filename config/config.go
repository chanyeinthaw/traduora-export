package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const configFile = "traduora.yaml"

type Config struct {
	Host string `yaml:"host"`

	ProjectId    string `yaml:"projectId"`
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`

	Locales []string `yaml:"locales"`

	OutputDir string `yaml:"outputDir"`
}

var cfg Config

func Read() {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		generateCfg := &Config{}
		outBytes, _ := yaml.Marshal(generateCfg)

		err = os.WriteFile(configFile, outBytes, 0755)

		if err == nil {
			fmt.Println("The configuration file was not found. A new one has been generated.\nEdit configuration in", configFile)
			os.Exit(0)
		}
	}

	bytes, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Unable to read %s\n", configFile)
		os.Exit(1)
	}

	err = yaml.Unmarshal(bytes, &cfg)

	if err != nil {
		fmt.Printf("Invalid configuration file")
		os.Exit(1)
	}
}

func ProjectId() string {
	return cfg.ProjectId
}

func ClientId() string {
	return cfg.ClientId
}

func ClientSecret() string {
	return cfg.ClientSecret
}

func OutputDir() string {
	return cfg.OutputDir
}

func Locales() []string {
	return cfg.Locales
}

func ApiURL(path string) string {
	return fmt.Sprintf("%s%s", cfg.Host, path)
}
