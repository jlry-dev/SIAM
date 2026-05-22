package config

import (
	"os"
	"strings"
)

func LoadEnv() {
	b, err := os.ReadFile(".env")
	if err != nil {
		return
	}

	for _, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)

		if len(parts) == 2 {
			os.Setenv(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}
}

