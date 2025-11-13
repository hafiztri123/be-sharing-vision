package utils

import (
	"fmt"
	"os"
)

func GetMandatoryEnv(key string) string {
	if val, ok := os.LookupEnv(key); !ok {
		panic(fmt.Sprintf("Missing environment variable: %s", key))
	} else {
		return val
	}
}


func GetOptionalEnv(key string, defaultValue string) string {
	if val, ok := os.LookupEnv(key); !ok {
		return defaultValue
	} else {
		return val
	}
} 
