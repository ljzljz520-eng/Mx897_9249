package config

import "strings"

func (c Config) IsLocal() bool {
	return strings.HasPrefix(c.DataFile, "/") || strings.HasPrefix(c.DataFile, ".")
}
func (c Config) AddressParts() []string { return strings.Split(c.ListenAddr, ":") }
func DefaultDataFile() string           { return "library.db" }
func DefaultListenAddr() string         { return ":8080" }
func Merge(base, override Config) Config {
	if override.DataFile != "" {
		base.DataFile = override.DataFile
	}
	if override.ListenAddr != "" {
		base.ListenAddr = override.ListenAddr
	}
	return base
}
