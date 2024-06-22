package environment

import "os"

const PREFIX string = "TINY_TRAIL"

func withPrefix(key string) string {
	return PREFIX + "_" + key
}

func GetConfigDir() string {
	return os.Getenv(withPrefix("CONFIG_DIR"))
}
