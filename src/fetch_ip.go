package main

import (
	"fmt"
	"time"

	"github.com/triole/logseal"
)

func (conf *tConf) getMyIP() (err error) {
	err = conf.getMyIPWorker()
	lg.IfErrError(
		"can not fetch current ip", logseal.F{"error": err},
	)
	return err
}

func (conf *tConf) getMyIPWorker() (err error) {
	var ipdse tIPDataSet
	ch := make(chan tIPDataSet)
	for _, url := range conf.RetrievalURLs {
		go conf.fetchIP(url, ch)
	}
	for i := 0; i < len(conf.RetrievalURLs); i++ {
		ipdse = <-ch
		lg.Trace("received ip dataset entry from channel", logseal.F{
			"ipdse": fmt.Sprintf("%+v", ipdse),
		})
		for idx, dns := range conf.DNSs {
			if !isValidIP(dns.IPToSend.IP) && isValidIP(ipdse.IP) && dns.IPv6 == ipdse.IPv6 {
				dns.IPToSend = ipdse
				conf.DNSs[idx] = dns
				lg.Trace("set ip to send for dns entry", logseal.F{
					"updated_dns_entry": fmt.Sprintf("%+v", dns),
				})
			}
		}
		if conf.allIPsFilled() {
			lg.Trace("all required ips were updated, break loop")
			break
		}
	}
	return err
}

func (conf *tConf) allIPsFilled() (b bool) {
	b = true
	for _, dns := range conf.DNSs {
		if !isValidIP(dns.IPToSend.IP) {
			b = false
			break
		}
	}
	return
}

func (conf *tConf) fetchIP(url string, ch chan<- tIPDataSet) {
	var ipdse tIPDataSet
	req := conf.req("get", url)
	if len(req.Errors) == 0 {
		ipdse.Time = time.Now()
		ipdse.IP = rxFindIPv4(req.Body)
		if ipdse.IP == "" {
			ipdse.IP = rxFindIPv6(req.Body)
		}
		ipdse.IPv6 = isIPv6(ipdse.IP)
		if ipdse.IP != "" {
			lg.Info("ip found", logseal.F{
				"url":  url,
				"date": ipdse.Time,
				"ip":   ipdse.IP,
				"ipv6": ipdse.IPv6,
			})
		}
	}
	ch <- ipdse
}
