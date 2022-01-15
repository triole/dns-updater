package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pelletier/go-toml"
)

func readConf(filename string) (conf tConf) {
	if strings.HasSuffix(filename, ".toml") == false {
		filename = filename + ".toml"
	}
	filename = path.Join("embed/conf", filename)
	data, err := efs.ReadFile(filename)
	lg.LogIfFileError("read", filename, err, true)

	err = toml.Unmarshal(data, &conf)
	lg.LogIfFileError("unmarshal", filename, err, true)

	conf.RetrievalConf = readRetrievalConf()
	conf.IPDataJSON = path.Join(os.TempDir(), "dns-updater.json")
	conf.ForceUpdate = CLI.Force
	conf.DryRun = CLI.DryRun
	conf.Debug = CLI.Debug
	if conf.Debug == false {
		conf.URL = strings.Replace(conf.URL, "{{.HOSTNAME}}", conf.Hostname, 1)
	}
	return conf
}

func readRetrievalConf() tRetrievalConf {
	var ipru tRetrievalConf
	filename := "embed/retrieval_conf.toml"
	data, err := efs.ReadFile(filename)
	lg.LogIfFileError("read embedded", filename, err, true)

	err = toml.Unmarshal(data, &ipru)
	lg.LogIfFileError("unmarshal embedded", filename, err, true)
	return ipru
}

func listConfigs() {
	fmt.Printf("Available configurations:\n")
	arr, _ := efs.ReadDir("embed/conf")
	for _, el := range arr {
		fmt.Printf("  %s\n", strings.Replace(el.Name(), ".toml", "", -1))
	}
}
