package frt

import (
	"strconv"
)

func GetInt64(value string, defaultValue int64) int64 {
	if value == "" {
		return defaultValue
	}
	if parsedValue, err := strconv.ParseInt(value, 10, 64); err == nil {
		return parsedValue
	}
	return defaultValue
}

func GetInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if parsedValue, err := strconv.Atoi(value); err == nil {
		return parsedValue
	}
	return defaultValue
}

func GetString(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
