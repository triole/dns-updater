package main

import "time"

type tConf struct {
	Hostname        string `toml:"hostname"`
	Token           string `toml:"token"`
	URL             string `toml:"url"`
	IPRetrievalURLs []string
	IPData          tIPData
	IPDataJSON      string
	IPChanged       bool
	ForceUpdate     bool
}

type tIPRetrievalURLs struct {
	URLs []string `toml:"urls"`
}

type tIPData struct {
	Current tIPDataSet `json:"current"`
	Old     tIPDataSet `json:"old"`
}

type tIPDataSet struct {
	Time time.Time `json:"time"`
	IP   string    `json:"ip"`
}
