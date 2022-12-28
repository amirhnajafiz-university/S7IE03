package worker

type Config struct {
	Timeout int `koanf:"timeout"`
	Workers int `koanf:"workers"`
}
