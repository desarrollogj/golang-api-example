package system

import "os"

// GetEnv gets an enviroment variable value. If it is not found, a fallback value is returned.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
