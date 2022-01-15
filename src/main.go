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

	if CLI.List == true {
		listConfigs()
	} else {

		conf := readConf(CLI.Config)
		var err error

		if CLI.IP != "" {
			conf.IPData.Current.IP = CLI.IP
			CLI.Force = true
		} else {
			conf.IPData.Current = getCurrentIPData(conf)
		}

		if conf.IPData.Current.IP == "" {
			lg.LogError("None of the made requests did return a valid IP address.", nil)
			os.Exit(1)
		} else {
			b := didIPChange(conf.IPData.Current)
			if b == true || CLI.Force == true {
				if CLI.Force == true {
					lg.LogInfo("Force update request irrespective of the current ip", nil)
				}
				writeIPDataJSON(conf.IPData.Current)
				err = updateDNS(conf, conf.IPData.Current.IP)
				if err != nil {
					os.Exit(1)
				}
			}
		}
	}
}
