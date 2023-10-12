package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

func readConf(filename string) (conf tConf) {
	var err error
	raw, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading config %q, %q", filename, err)
	}
	ext := filepath.Ext(filename)
	if ext == ".toml" {
		err = toml.Unmarshal(raw, &conf)
	}
	if ext == ".yaml" {
		err = yaml.Unmarshal(raw, &conf)
	}
	if err != nil {
		log.Fatalf("unmarshal error %q, %q", filename, err)
	}

	// conf.IPDataJSON = path.Join(os.TempDir(), "dns-updater.json")
	// conf.ForceUpdate = CLI.Force
	// conf.DryRun = CLI.DryRun
	// conf.Debug = CLI.Debug
	// if !conf.Debug {
	// 	conf.URL = strings.Replace(conf.URL, "{{.HOSTNAME}}", conf.Hostname, 1)
	// }
	return conf
}
