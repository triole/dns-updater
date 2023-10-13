package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/triole/logseal"
)

func main() {
	parseArgs()
	lg = logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	lg.Info("start dns-updater", logseal.F{
		"loglevel": CLI.LogLevel,
		"logfile":  CLI.LogFile,
	})

	conf := readConf(CLI.Config)
	lg.Debug("config layout and data json location", logseal.F{
		"conf":     fmt.Sprintf("%+v", conf),
		"datajson": conf.DataJSONFile,
	})
	_ = conf.getMyIP()

	if CLI.IP != "" {
		conf.IPData.Current.IP = CLI.IP
		CLI.Force = true
		os.Exit(0)
	}

	// conf.IPData.Current, _ = conf.getCurrentIPData(conf)
	if conf.IPData.Current.IP == "" {
		lg.Fatal("ip retrieval failed", logseal.F{"ip": conf.IPData.Current.IP})
	}

	conf.IPData.Old = conf.readIPDataJSON()
	conf.IPChanged = conf.IPData.Old.IP != conf.IPData.Current.IP
	if conf.IPChanged || CLI.Force {
		conf.iterDNSServicesAndPost()
	}
}

func (conf *tConf) iterDNSServicesAndPost() {
	conf.writeIPDataJSON(conf.IPData.Current)
	for _, dns := range conf.DNSs {
		dns.URL = strings.Replace(
			dns.URL, "{{.IP}}", conf.IPData.Current.IP, 1,
		)
		lg.Info("fire update request", logseal.F{
			"url": dns.URL,
		})
		err := makeUpdateRequest(dns)
		lg.IfErrError(
			err,
			"update request failed", logrus.Fields{
				"current_ip": conf.IPData.Current.IP,
				"err":        err,
			},
		)
	}
}
