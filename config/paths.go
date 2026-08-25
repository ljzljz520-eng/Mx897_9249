package config

import "path/filepath"

func DatabasePath(c Config, base string) string {
	if filepath.IsAbs(c.DataFile) {
		return c.DataFile
	}
	return filepath.Join(base, c.DataFile)
}
func IsDatabaseFile(path string) bool    { return filepath.Ext(path) == ".db" }
func CleanPath(path string) string       { return filepath.Clean(path) }
func Join(base, name string) string      { return filepath.Join(base, name) }
func SamePath(a, b string) bool          { return filepath.Clean(a) == filepath.Clean(b) }
func HasExtension(path, ext string) bool { return filepath.Ext(path) == ext }
