package env

import "syscall"

func EnvString(key string, defaultVal string) string {
	if val, ok := syscall.Getenv(key); ok {
		return val
	}
	return defaultVal
}
