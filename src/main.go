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
			lg.LogFatal("ip retrieval failed", nil)
		} else if conf.Debug == true {
			lg.LogInfo("debug mode, exit here", logrus.Fields{
				"err": conf.IPData,
			},
			)
		} else {
			conf.IPData.Old = readIPDataJSON()
			conf.IPChanged = conf.IPData.Old.IP != conf.IPData.Current.IP
			if conf.IPChanged == true || CLI.Force == true {
				writeIPDataJSON(conf.IPData.Current)
				conf.URL = strings.Replace(
					conf.URL, "{{.IP}}", conf.IPData.Current.IP, 1,
				)
				err = updateDNS(conf)
				lg.LogIfError(
					err,
					"update request failed", logrus.Fields{
						"err": err,
					},
				)
			} else {
				lg.LogStatus(
					"skip update", conf.IPData, conf.ForceUpdate,
				)
			}
		}
	}
}
