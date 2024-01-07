package pkg

import (
	"os"
	"strconv"
)

// START: loading environment function
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvInt(key string, fallback int) int {
	env := GetEnv(key, strconv.Itoa(fallback))
	retval, err := strconv.Atoi(env)
	if err != nil {
		return fallback
	}
	return retval
}

func GetEnvBool(key string, fallback bool) bool {
	env := GetEnv(key, strconv.FormatBool(fallback))
	retval, err := strconv.ParseBool(env)
	if err != nil {
		return fallback
	}
	return retval
}

// END: loading environment function
