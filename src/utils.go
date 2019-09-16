package main

import (
	"olibs/netop"
	"strings"
	"time"
)

func didIPChange(currentIPData IPData) (b bool) {
	b = false
	oldIPData := readIPDataJSON()
	if currentIPData.IP != oldIPData.IP {
		b = true
		lg.Log("IP did change.")
	} else {
		t := oldIPData.Time
		lg.Log("IP did not change. It is stable since " + t.Format("Monday, Jan _2 15:04:05 2006"))
	}
	return
}

func getCurrentIPData() (ipd IPData) {
	ip, err := netop.GetMyIP()
	if err == nil {
		ipd = IPData{
			Time: time.Now(),
			IP:   ip,
		}
		lg.Log("Current IP is " + ipd.IP + ".")
	}
	return
}

func updateDNS(ip string) {
	lg.Logf("Updating ip at dns service. New ip is %v", ip)
	url := strings.Replace(requrl, "[IP]", ip, 1)
	makeUpdateRequest(url)
}
