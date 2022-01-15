package main

import "time"

type tConf struct {
	Hostname        string            `toml:"hostname"`
	Token           string            `toml:"token"`
	URL             string            `toml:"url"`
	RequestHeaders  map[string]string `toml:"request_headers"`
	IPRetrievalURLs []string
	IPData          tIPData
	IPDataJSON      string
	IPChanged       bool
	ForceUpdate     bool
	DryRun          bool
	Debug           bool
	RetrievalConf   tRetrievalConf
}

type tRetrievalConf struct {
	URLs     []string `toml:"urls"`
	MoreInfo []string `toml:"more_info"`
	TorCheck string   `toml:"torcheck"`
}

type tIPData struct {
	Current tIPDataSet `json:"current"`
	Old     tIPDataSet `json:"old"`
}

type tIPDataSet struct {
	Time time.Time `json:"time"`
	IP   string    `json:"ip"`
}
