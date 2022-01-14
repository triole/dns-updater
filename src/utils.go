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
		lg.LogInfo(
			"IP changed", map[string]string{
				"oldIP":     oldIPData.IP,
				"currentIP": currentIPData.IP,
			},
		)
	} else {
		t := oldIPData.Time
		lg.LogInfo("IP stable since", t.Format("Monday, Jan _2 15:04:05 2006"))
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
		lg.LogInfo("Current IP is", ipd.IP)
	}
	return
}
