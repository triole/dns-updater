package main

import (
	"fmt"
	"os"

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

	if CLI.TestRetrieval {
		for _, url := range conf.RetrievalURLs {
			res := conf.req("get", url, regexIPv4)
			conf.ExitCode += len(res.Errors)
		}
		conf.exit()
	}

	err := conf.getMyIP()
	if err != nil {
		lg.Fatal("ip retrieval failed", logseal.F{"ip": conf.IPData.Current.IP})
	}

	if CLI.IP != "" {
		conf.IPData.Current.IP = CLI.IP
		CLI.Force = true
	}

	conf.IPData.Old = conf.readIPDataJSON()
	conf.IPChanged = conf.IPData.Old.IP != conf.IPData.Current.IP
	lg.Info("ip comparison", logseal.F{
		"former":  fmt.Sprintf("%+v", conf.IPData.Old),
		"current": fmt.Sprintf("%+v", conf.IPData.Current),
		"changed": conf.IPChanged,
	})
	if conf.IPChanged || CLI.Force {
		conf.writeIPDataJSON(conf.IPData.Current)
		conf.iterDNSServicesAndPost()
	} else {
		lg.Info("skip dns update")
	}
	conf.exit()
}

func (conf *tConf) iterDNSServicesAndPost() {
	for _, dns := range conf.DNSs {
		conf.makeUpdateRequest(dns)
	}
}

func (conf *tConf) exit() {
	lg.Info("done", logseal.F{"exitcode": conf.ExitCode})
	os.Exit(conf.ExitCode)
}
