package environment

import "os"

// PREFIX is the prefix for all environment variables used in the application.
const PREFIX string = "TINY_TRAIL"

// WithPrefix adds the prefix to the key.
func WithPrefix(key string) string {
	return PREFIX + "_" + key
}

// Gets the directory string where the configuration files are stored.
func GetConfigDir() string {
	return os.Getenv(WithPrefix("CONFIG_DIR"))
}
