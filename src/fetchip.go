package main

import (
	"errors"
	"io"
	"net/http"
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
	}
	return err
}

func (conf *tConf) getMyIPWorker() (ip string, err error) {
	ch := make(chan string)
	for _, url := range conf.Retrieval.URLs {
		go conf.fetchIP(url, ch)
	}
	for i := 0; i < len(conf.Retrieval.URLs); i++ {
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
	var ip string
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	lg.Debug("fetch current ip", logseal.F{
		"url": url,
	})
	resp, err := client.Get(url)
	if err == nil {
		body, _ := io.ReadAll(resp.Body)
		ip = rxFindIP(string(body))
		if ip != "" {
			lg.Info("fetch current ip success", logseal.F{
				"url": url,
				"ip":  ip,
			})
		}
	} else {
		lg.Error("request failed", logseal.F{
			"url": url,
		})
	}
	ch <- ip
}
