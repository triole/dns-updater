package main

import (
	"fmt"

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
	}

	// conf.IPData.Current, _ = conf.getCurrentIPData(conf)
	if conf.IPData.Current.IP == "" {
		lg.Fatal("ip retrieval failed", logseal.F{"ip": conf.IPData.Current.IP})
	}

	conf.IPData.Old = conf.readIPDataJSON()
	conf.IPChanged = conf.IPData.Old.IP != conf.IPData.Current.IP
	lg.Info("ip comparison", logseal.F{
		"former":  fmt.Sprintf("%+v", conf.IPData.Old),
		"current": fmt.Sprintf("%+v", conf.IPData.Current),
		"changed": conf.IPChanged,
	})
	if conf.IPChanged || CLI.Force {
		conf.iterDNSServicesAndPost()
	} else {
		lg.Info("skip dns update")
	}
	lg.Info("done", logseal.F{"exitcode": conf.ExitCode})
}

func (conf *tConf) iterDNSServicesAndPost() {
	conf.writeIPDataJSON(conf.IPData.Current)
	for _, dns := range conf.DNSs {
		err := conf.makeUpdateRequest(dns)
		lg.IfErrError(
			err,
			"update request failed", logrus.Fields{
				"current_ip": conf.IPData.Current.IP,
				"err":        err,
			},
		)
	}
}
