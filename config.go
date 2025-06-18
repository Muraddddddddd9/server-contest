package main

import (
	"contest/constants"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ORIGIN_URL string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf(constants.ErrLoadEnv)
	}

	return &Config{
		ORIGIN_URL: GetEnv("ORIGIN_URL"),
	}, nil
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		return ""
	}
	return value
}
