package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

func readConf(filename string) (conf tConf) {
	if strings.HasSuffix(filename, ".toml") == false {
		filename = filename + ".toml"
	}
	filename = path.Join("embed/conf", filename)
	data, err := efs.ReadFile(filename)
	if err != nil {
		lg.LogFatal("Can not read config", logrus.Fields{
			"err":      err.Error(),
			"filename": filename,
		})
	}
	if err := toml.Unmarshal(data, &conf); err != nil {
		if err != nil {
			lg.LogFatal("Can not unmarshal config", logrus.Fields{
				"filename": filename,
				"err":      err.Error(),
			})
		}
	}
	conf.IPRetrievalURLs = readIpURLs()
	return conf
}

func readIpURLs() []string {
	var ipru tIPRetrievalURLs
	filename := "embed/ip_retrieval_urls.toml"
	data, err := efs.ReadFile(filename)
	if err != nil {
		lg.LogFatal("Can not read ip retrieval urls", logrus.Fields{
			"err":      err.Error(),
			"filename": filename,
		})
	}
	if err := toml.Unmarshal(data, &ipru); err != nil {
		if err != nil {
			lg.LogFatal("Can not unmarshal ip retrieval urls", logrus.Fields{
				"err":      err,
				"filename": filename,
			})
		}
	}
	return ipru.URLs
}

func listConfigs() {
	fmt.Printf("Available configurations:\n")
	arr, _ := efs.ReadDir("embed/conf")
	for _, el := range arr {
		fmt.Printf("  %s\n", strings.Replace(el.Name(), ".toml", "", -1))
	}
}
