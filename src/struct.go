package main

import "time"

type tConf struct {
	DNSs               []tDNS   `toml:"dynamic_name_services"`
	RetrievalURLs      []string `toml:"retrieval_urls"`
	IPData             tIPData
	DataJSONFile       string
	IPChanged          bool
	RequestsTimeoutInt int64
	RequestsTimeout    time.Duration
	ForceUpdate        bool
	DryRun             bool
	Debug              bool
	ExitCode           int
}

type tDNS struct {
	Hostname string `toml:"hostname"`
	Token    string `toml:"token"`
	URL      string `toml:"url"`
	Method   string `toml:"method"`
}

type tIPData struct {
	Current tIPDataSet `json:"current"`
	Old     tIPDataSet `json:"old"`
}

type tIPDataSet struct {
	Time time.Time `json:"time"`
	IP   string    `json:"ip"`
}
