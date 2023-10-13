package main

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pelletier/go-toml"
	"github.com/triole/logseal"
	"gopkg.in/yaml.v3"
)

func readConf(filename string) (conf tConf) {
	var err error
	lg.Info("read config", logseal.F{"file": filename})
	raw, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error reading config %q, %q", filename, err)
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
	conf.RequestsTimeoutInt = CLI.Timeout
	conf.RequestsTimeout = time.Duration(
		time.Duration(conf.RequestsTimeoutInt) * time.Second,
	)
	conf.ForceUpdate = CLI.Force
	conf.DryRun = CLI.DryRun
	conf.DataJSONFile = CLI.DataJSON

	for idx, el := range conf.RetrievalURLs {
		conf.RetrievalURLs[idx] = os.ExpandEnv(el)
	}
	return conf
}
