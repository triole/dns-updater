package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/triole/logseal"
)

func (conf tConf) getMyIP(ipRetrievalURLs []string) (err error) {
	var ip string
	ip, err = conf.getMyIPWorker(ipRetrievalURLs)
	lg.IfErrError(
		"can not fetch current ip", logseal.F{"error": err},
	)
	if err == nil {
		lg.Info(
			"fetch current ip success", logseal.F{"ip": ip},
		)
	}
	conf.IPData.Current.IP = ip
	return err
}

func (conf tConf) getMyIPWorker(urlList []string) (ip string, err error) {
	ch := make(chan string)
	for _, url := range urlList {
		go conf.fetchIP(url, ch)
	}
	for i := 0; i < len(urlList); i++ {
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

func (conf tConf) fetchIP(url string, ch chan<- string) {
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
		lg.Debug("current ip fetch success", logseal.F{
			"url": url,
			"ip":  ip,
		})
	} else {
		lg.Error("request failed", logseal.F{
			"url": url,
		})
	}
	ch <- ip
}

func (conf tConf) torCheck() {
	body, err := req(conf.Retrieval.TorCheck)
	if err == nil {
		torEnabled := rxFind("You are not using Tor", body) == ""
		fmt.Printf("Tor enabled:  %v\n", torEnabled)
	}

}

func moreInformation(conf tConf) {
	for _, url := range conf.Retrieval.MoreInfo {
		body, err := req(url)
		if err == nil {
			lg.Info("more information", logseal.F{
				"url":  url,
				"body": body,
			})
		}
	}
	fmt.Printf("")
}
