package common

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// ParseJSONConfig parse configuration file or environment variables, receiver must be a pointer
func ParseJSONConfig(configFile string, receiver interface{}, prefix string) error {
	if configFile != "" {
		b, err := os.ReadFile(filepath.Clean(configFile))
		if err != nil && !os.IsNotExist(err) {
			return err
		}
		if b != nil {
			if err := json.Unmarshal(b, receiver); err != nil {
				return err
			}
		}
	}
	return nil
}
