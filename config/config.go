package config

import (
	"fmt"
	"os"
)

type Config struct{ DataFile, ListenAddr string }

func Load() Config {
	f := os.Getenv("LIBRARY_DATA")
	if f == "" {
		f = "library.db"
	}
	a := os.Getenv("LIBRARY_ADDR")
	if a == "" {
		a = ":8080"
	}
	return Config{f, a}
}
func Validate(c Config) error {
	if c.DataFile == "" {
		return fmt.Errorf("data file required")
	}
	if c.ListenAddr == "" {
		return fmt.Errorf("listen address required")
	}
	return nil
}
