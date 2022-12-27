package auth

type Config struct {
	PrivateKey string `koanf:"private_key"`
	ExpireTime int    `koanf:"expire_time"`
}
