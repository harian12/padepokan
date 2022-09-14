package helpers

import (
	"os"
)

func GetEnv(key, static string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return static
}
