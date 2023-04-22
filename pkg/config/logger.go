package config

import (
	"io"
	"log"
	"os"

	"github.com/ghodss/yaml"
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger = GenerateLog()

func GenerateLog() *zap.SugaredLogger {
	configFile, err := os.Open("config/logger.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	os.MkdirAll("logs", 0764)
	config, err := io.ReadAll(configFile)
	if err != nil {
		log.Fatal(err)
	}

	var cfg *zap.Config
	if err := yaml.Unmarshal(config, &cfg); err != nil {
		log.Fatal(err)
	}

	logger, err := cfg.Build()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()
	return sugar
}
