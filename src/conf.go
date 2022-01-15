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
	lg.LogIfFileError("can not read config", filename, err, true)

	err = toml.Unmarshal(data, &conf)
	lg.LogIfFileError("can not unmarshal config", filename, err, true)

	conf.IPRetrievalURLs = readIpURLs()
	conf.IPDataJSON = path.Join(os.TempDir(), "dns-updater.json")
	conf.ForceUpdate = CLI.Force
	conf.URL = strings.Replace(conf.URL, "{{.HOSTNAME}}", conf.Hostname, 1)

	return conf
}

func readIpURLs() []string {
	var ipru tIPRetrievalURLs
	filename := "embed/ip_retrieval_urls.toml"
	data, err := efs.ReadFile(filename)
	lg.LogIfFileError("can not read embedded file", filename, err, true)

	err = toml.Unmarshal(data, &ipru)
	lg.LogIfFileError("can not unmarshal embedded file", filename, err, true)
	return ipru.URLs
}

func listConfigs() {
	fmt.Printf("Available configurations:\n")
	arr, _ := efs.ReadDir("embed/conf")
	for _, el := range arr {
		fmt.Printf("  %s\n", strings.Replace(el.Name(), ".toml", "", -1))
	}
}
