package main

import (
	"embed"
	"os"
)

//go:embed embed/**
var efs embed.FS

func main() {
	parseArgs()
	lg = initLogging(CLI.Logfile)

	conf := readConf(CLI.Config)
	var err error
	var currentIPData IPData

	os.Exit(0)

	if CLI.IP != "" {
		currentIPData.IP = CLI.IP
		CLI.Force = true
	} else {
		currentIPData = getCurrentIPData(conf)
	}

	if currentIPData.IP == "" {
		lg.LogError("None of the made requests did return a valid IP address.", nil)
		os.Exit(1)
	} else {
		b := didIPChange(currentIPData)
		if b == true || CLI.Force == true {
			if CLI.Force == true {
				lg.LogInfo("Force update request irrespective of the current ip", nil)
			}
			writeIPDataJSON(currentIPData)
			err = updateDNS(conf, currentIPData.IP)
			if err != nil {
				os.Exit(1)
			}
		}
	}
}
