package main

import (
	"olibs/netop"
	"time"
)

func didIPChange(currentIPData IPData) (b bool) {
	b = false
	oldIPData := readIPDataJSON()
	if currentIPData.IP != oldIPData.IP {
		b = true
		lg.Logf("IP changed %q -> %q", oldIPData.IP, currentIPData.IP)
	} else {
		t := oldIPData.Time
		lg.Log("IP stable since " + t.Format("Monday, Jan _2 15:04:05 2006"))
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
		lg.Logf("Current IP is %q", ipd.IP)
	}
	return
}
