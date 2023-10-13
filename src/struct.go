package main

import "time"

type tConf struct {
	DNSs          []tDNS   `toml:"dynamic_name_services"`
	RetrievalURLs []string `toml:"retrieval_urls"`
	IPData        tIPData
	DataJSONFile  string
	IPChanged     bool
	ForceUpdate   bool
	DryRun        bool
	Debug         bool
}

type tDNS struct {
	Hostname       string            `toml:"hostname"`
	Token          string            `toml:"token"`
	URL            string            `toml:"url"`
	RequestHeaders map[string]string `toml:"request_headers"`
}

type tIPData struct {
	Current tIPDataSet `json:"current"`
	Old     tIPDataSet `json:"old"`
}

type tIPDataSet struct {
	Time time.Time `json:"time"`
	IP   string    `json:"ip"`
}
