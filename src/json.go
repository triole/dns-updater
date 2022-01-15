package main

import (
	"encoding/json"
	"io/ioutil"

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
	data, err := fs.ReadFile(filename)
	lg.LogFatal("Can not read config", logrus.Fields{
		"err":      err,
		"filename": filename,
	})
	if err := toml.Unmarshal(data, conf); err != nil {
		lg.LogFatal("Can not unmarshal config", logrus.Fields{
			"err":      err,
			"filename": filename,
		})
	}
	conf.IpRetrievalURLs = readIpURLs("ip_retrieval_urls.toml")
	return conf
}

func readIpURLs(filename string) (urls []string) {
	data, err := fs.ReadFile(filename)
	lg.LogFatal("Can not read ip retrieval urls", logrus.Fields{
		"err":      err,
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

func readIPDataJSON() (ipd IPData) {
	raw, err := ioutil.ReadFile(fileIPDataJSON)
	lg.LogError("Error reading ip data json", err)
	err = json.Unmarshal(raw, &ipd)
	lg.LogError("Can not unmarshal ip data json", err)
	return
}

func writeIPDataJSON(ipd IPData) {
	var err error
	JSONstring, _ := json.Marshal(ipd)
	err = ioutil.WriteFile(fileIPDataJSON, JSONstring, 0644)
	lg.LogError("Error writing json file", err)

}
