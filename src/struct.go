package main

import "time"

type tConf struct {
	Hostname        string `toml:"hostname"`
	Token           string `toml:"token"`
	URL             string `toml:"url"`
	IPRetrievalURLs []string
	IPData          tIPData
}

type tIPRetrievalURLs struct {
	URLs []string `toml:"urls"`
}

type tIPData struct {
	Current tIPDataSet
}

type tIPDataSet struct {
	Time time.Time
	IP   string
}
