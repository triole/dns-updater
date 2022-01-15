package main

import (
	"embed"
	"strings"

	"github.com/sirupsen/logrus"
)

//go:embed embed/**
var efs embed.FS

func main() {
	parseArgs()
	lg = initLogging(CLI.Logfile)
	conf := readConf(CLI.Config)

	if CLI.List == true {
		listConfigs()
	} else if CLI.Info == true {
		displayConnectionInformation(conf)
	} else {
		if CLI.IP != "" {
			conf.IPData.Current.IP = CLI.IP
			CLI.Force = true
		} else {
			conf.IPData.Current = getCurrentIPData(conf)
		}

		if conf.IPData.Current.IP == "" {
			lg.LogFatal("ip retrieval failed", nil)
		} else {
			conf.IPData.Old = readIPDataJSON()
			conf.IPChanged = conf.IPData.Old.IP != conf.IPData.Current.IP
			lg.LogDebug("debug output", conf)
			if conf.IPChanged == true || CLI.Force == true {
				if conf.DryRun == true {
					lg.LogStatus(
						"dry run, skip update request", conf.IPData, conf.ForceUpdate,
					)
				} else {
					writeIPDataJSON(conf.IPData.Current)
					conf.URL = strings.Replace(
						conf.URL, "{{.IP}}", conf.IPData.Current.IP, 1,
					)
					err := updateDNS(conf)
					lg.LogIfError(
						err,
						"update request failed", logrus.Fields{
							"err": err,
						},
					)
				}
			} else {
				lg.LogStatus(
					"skip update request", conf.IPData, conf.ForceUpdate,
				)
			}
		}
	}
}
