package omsv2_commons

import (
	"os"
)

func EnvString(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}
