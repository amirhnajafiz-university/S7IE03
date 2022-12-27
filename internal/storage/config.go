package storage

// Config
// mongodb config parameters.
type Config struct {
	Database string `koanf:"database"`
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
}
