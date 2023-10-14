package main

import (
	"errors"
	"time"

	"github.com/triole/logseal"
)

func (conf *tConf) getMyIP() (err error) {
	var ip string
	ip, err = conf.getMyIPWorker()
	lg.IfErrError(
		"can not fetch current ip", logseal.F{"error": err},
	)
	if err == nil {
		conf.IPData.Current.IP = ip
		conf.IPData.Current.Time = time.Now()
	}
	return err
}

func (conf *tConf) getMyIPWorker() (ip string, err error) {
	ch := make(chan string)
	for _, url := range conf.RetrievalURLs {
		go conf.fetchIP(url, ch)
	}
	for i := 0; i < len(conf.RetrievalURLs); i++ {
		ip = <-ch
		if ip != "" {
			break
		}
	}
	if ip == "" {
		err = errors.New(
			"none of the fetch requests provided a valid IP address",
		)
	}
	return ip, err
}

func (conf *tConf) fetchIP(url string, ch chan<- string) {
	req := conf.req("get", url, rxIPAdresses)
	ch <- req.Match
}
