package main

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/triole/logseal"
)

func main() {
	parseArgs()
	lg = logseal.Init(CLI.LogLevel, CLI.LogFile, CLI.LogNoColors, CLI.LogJSON)
	conf := readConf(CLI.Config)

	_ = conf.getMyIP(conf.Retrieval.URLs)

	if !CLI.Info {
		if CLI.IP != "" {
			conf.IPData.Current.IP = CLI.IP
			CLI.Force = true
			os.Exit(0)
		}

		// conf.IPData.Current, _ = getCurrentIPData(conf)
		if conf.IPData.Current.IP == "" {
			lg.Fatal("ip retrieval failed", nil)
		}

		conf.IPData.Old = conf.readIPDataJSON()
		conf.IPChanged = conf.IPData.Old.IP != conf.IPData.Current.IP
		if conf.IPChanged || CLI.Force {
			conf.iterDNSServicesAndPost()
		}
	}
}

func (conf tConf) iterDNSServicesAndPost() {
	conf.writeIPDataJSON(conf.IPData.Current)
	for _, dns := range conf.DNSs {
		dns.URL = strings.Replace(
			dns.URL, "{{.IP}}", conf.IPData.Current.IP, 1,
		)
		err := updateDNS(dns)
		lg.IfErrError(
			err,
			"update request failed", logrus.Fields{
				"current_ip": conf.IPData.Current.IP,
				"err":        err,
			},
		)
	}
}
