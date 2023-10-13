package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
	"github.com/triole/logseal"
	"gopkg.in/yaml.v3"
)

func readConf(filename string) (conf tConf) {
	var err error
	lg.Info("read config", logseal.F{"file": filename})
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
	conf.ForceUpdate = CLI.Force
	conf.DryRun = CLI.DryRun
	for idx, el := range conf.Retrieval.URLs {
		conf.Retrieval.URLs[idx] = os.ExpandEnv(el)
	}
	return conf
}
