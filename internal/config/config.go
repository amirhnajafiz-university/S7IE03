package config

import (
	"encoding/json"
	"log"

	"github.com/ceit-aut/S7IE03/internal/storage"
	"github.com/ceit-aut/S7IE03/internal/worker"
	"github.com/ceit-aut/S7IE03/pkg/auth"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/tidwall/pretty"
)

type Config struct {
	HttpPort      int            `koanf:"http_port"`
	Threshold     int            `koanf:"threshold"`
	UserEndpoints int            `koanf:"user_endpoints"`
	JWT           auth.Config    `koanf:"jwt"`
	Storage       storage.Config `koanf:"mongodb"`
	Worker        worker.Config  `koanf:"worker"`
}

// Load reads configuration with koanf.
func Load() Config {
	var instance Config

	k := koanf.New(".")

	// load default configuration from file
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load configuration from file
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Printf("error loading config.yml: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	indent, err := json.MarshalIndent(instance, "", "\t")
	if err != nil {
		log.Fatalf("error marshaling config to json: %s", err)
	}

	indent = pretty.Color(indent, nil)
	tmpl := `
	================ Loaded Configuration ================
	%s
	======================================================
	`
	log.Printf(tmpl, string(indent))

	return instance
}
