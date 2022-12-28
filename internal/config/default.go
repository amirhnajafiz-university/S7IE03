package config

import (
	"github.com/ceit-aut/policeman/internal/storage"
	"github.com/ceit-aut/policeman/internal/worker"
	"github.com/ceit-aut/policeman/pkg/auth"
)

func Default() Config {
	return Config{
		HttpPort:      8080,
		UserEndpoints: 20,
		Threshold:     20,
		JWT: auth.Config{
			PrivateKey: "",
			ExpireTime: 5,
		},
		Storage: storage.Config{
			Database: "",
			Host:     "",
			Port:     5000,
		},
		Worker: worker.Config{
			Timeout: 5,
			Workers: 10,
		},
	}
}
