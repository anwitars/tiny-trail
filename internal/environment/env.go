package environment

import "os"

// PREFIX is the prefix for all environment variables used in the application.
const PREFIX string = "TINY_TRAIL"

// withPrefix adds the prefix to the key.
func withPrefix(key string) string {
	return PREFIX + "_" + key
}

// Gets the directory string where the configuration files are stored.
func GetConfigDir() string {
	return os.Getenv(withPrefix("CONFIG_DIR"))
}
