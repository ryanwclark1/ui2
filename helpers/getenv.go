package helpers

import (
	"os"
)

// Getenv gets the given environment variable. This is a more advanced version
// of the built-in os.Getenv function.
//
// If key is not found, Getenv sets it to a fallback value.
//
// Example:

//	// Get a value of the environment variable 'BACKEND_PORT' or sets it to a fallback value '5000'.
//	helpers.Getenv("BACKEND_PORT", "5000")
func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
