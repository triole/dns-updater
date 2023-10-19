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
	conf.readOldIPDataJSON()
	fmt.Printf("%+v\n", conf.OldIPDataJSON)

	lg.Debug("config layout and data json location", logseal.F{
		"conf":     fmt.Sprintf("%+v", conf),
		"datajson": conf.DataJSONFile,
	})

	if CLI.TestRetrieval {
		ch := make(chan tIPDataSet, 1)
		for _, url := range conf.RetrievalURLs {
			conf.fetchIP(url, ch)
			<-ch
		}
		conf.exit()
	}

	conf.getMyIP()
	conf.iterDNSServicesAndPost()
	conf.exit()
}

func (conf *tConf) iterDNSServicesAndPost() {
	for _, dns := range conf.DNSs {
		oldDNSsIP := conf.pickOldMatchingDNSsIP(dns)
		conf.IPChanged = dns.IPToSend.IP != oldDNSsIP

		lg.Info("ip comparison", logseal.F{
			"former":  fmt.Sprintf("%+v", oldDNSsIP),
			"current": fmt.Sprintf("%+v", dns.IPToSend.IP),
			"changed": conf.IPChanged,
		})
		if conf.IPChanged || conf.ForceUpdate {
			conf.writeIPDataJSON()
		} else {
			lg.Info("skip dns update, unchanged")
		}
		if !CLI.DryRun {
			conf.makeUpdateRequest(dns)
		} else {
			lg.Info("skip dns update, dry run")
		}
	}
}

func (conf *tConf) pickOldMatchingDNSsIP(dns tDNS) (ip string) {
	for _, oldDNS := range conf.OldIPDataJSON {
		if oldDNS.URL == dns.URL &&
			oldDNS.Hostname == dns.Hostname && oldDNS.IPv6 == dns.IPv6 {
			ip = oldDNS.IPToSend.IP
		}
	}
	return
}

func (conf *tConf) exit() {
	lg.Info("done", logseal.F{"exitcode": conf.ExitCode})
	os.Exit(conf.ExitCode)
}
