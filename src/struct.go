package main

import "time"

type tConf struct {
	DNSs               tDNSs    `toml:"dynamic_name_services"`
	RetrievalURLs      []string `toml:"retrieval_urls"`
	DataJSONFile       string
	IPChanged          bool
	RequestsTimeoutInt int64
	RequestsTimeout    time.Duration
	ForceUpdate        bool
	DryRun             bool
	Debug              bool
	ExitCode           int
	OldIPDataJSON      tDNSs
}

type tDNSs []tDNS

type tDNS struct {
	Hostname string `toml:"hostname"`
	Token    string `toml:"token"`
	URL      string `toml:"url"`
	Method   string `toml:"method"`
	IPv6     bool   `toml:"ipv6"`
	IPToSend tIPDataSet
}

type tIPDataSet struct {
	Time time.Time `json:"time"`
	IP   string    `json:"ip"`
	IPv6 bool      `json:"ipv6"`
}
