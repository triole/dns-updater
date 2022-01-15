package main

import (
	"fmt"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

type tConf struct {
	Hostname        string `toml:hostname`
	Token           string `toml:token`
	URL             string `toml:url`
	IpRetrievalURLs []string
}

func readConf(filename string) (conf tConf) {
	data, err := efs.ReadFile(filename)
	if err != nil {
		lg.LogFatal("Can not read config", logrus.Fields{
			"err":      err.Error(),
			"filename": filename,
		})
	}
	if err := toml.Unmarshal(data, conf); err != nil {
		if err != nil {
			lg.LogFatal("Can not unmarshal config", logrus.Fields{
				"filename": filename,
				"err":      err.Error(),
			})
		}
	}
	conf.IpRetrievalURLs = readIpURLs("ip_retrieval_urls.toml")
	return conf
}

func readIpURLs(filename string) (urls []string) {
	data, err := efs.ReadFile("embed/ip_retrieval_urls.toml")
	lg.LogFatal("Can not read ip retrieval urls", logrus.Fields{
		"err":      err.Error(),
		"filename": filename,
	})
	if err := toml.Unmarshal(data, urls); err != nil {
		lg.LogFatal("Can not unmarshal ip retrieval urls", logrus.Fields{
			"err":      err,
			"filename": filename,
		})
	}
	return urls
}

func listConfigs() {
	fmt.Printf("Available configurations:\n")
	arr, _ := efs.ReadDir("embed/conf")
	for _, el := range arr {
		fmt.Printf("  %s\n", strings.Replace(el.Name(), ".toml", "", -1))
	}
}
