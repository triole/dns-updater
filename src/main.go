package main

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	parseArgs()
	lg = initLogging(CLI.Logfile)
	conf := readConf(CLI.Config)

	if CLI.Info {
		displayConnectionInformation(conf)
		os.Exit(0)
	}

	if CLI.IP != "" {
		conf.IPData.Current.IP = CLI.IP
		CLI.Force = true
		os.Exit(0)
	}

	conf.IPData.Current, _ = getCurrentIPData(conf)
	if conf.IPData.Current.IP == "" {
		lg.LogFatal("ip retrieval failed", nil)
	}

	conf.IPData.Old = readIPDataJSON()
	conf.IPChanged = conf.IPData.Old.IP != conf.IPData.Current.IP
	if conf.IPChanged || CLI.Force {
		if conf.DryRun {
			lg.LogStatus(
				"dry run, skip update request", conf.IPData, conf.ForceUpdate,
			)
		} else {
			conf.iterDNSServicesAndPost()
		}
	} else {
		lg.LogStatus(
			"skip update request", conf.IPData, conf.ForceUpdate,
		)
	}
}

func (conf tConf) iterDNSServicesAndPost() {
	writeIPDataJSON(conf.IPData.Current)
	for _, dns := range conf.DNSs {
		dns.URL = strings.Replace(
			dns.URL, "{{.IP}}", conf.IPData.Current.IP, 1,
		)
		err := updateDNS(dns)
		lg.LogIfError(
			err,
			"update request failed", logrus.Fields{
				"current_ip": conf.IPData.Current.IP,
				"err":        err,
			},
		)
	}
}
