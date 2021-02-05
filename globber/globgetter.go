package globber

import (
	"path/filepath"
	"log"
)

// Globzipfiles globs zip files
func Globzipfiles() []string {
	matches, err := filepath.Glob("*.zip")
	if err != nil {
		log.Fatalf("Pattern malformed")
	}
	return matches
}