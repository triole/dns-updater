package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func getCurrentIPData(conf tConf) (ipd tIPDataSet) {
	ip, err := getMyIP(conf.IPRetrievalURLs)
	if err == nil {
		ipd = tIPDataSet{
			Time: time.Now(),
			IP:   ip,
		}
	}
	return
}

func getMyIP(ipRetrievalURLs []string) (ip string, err error) {
	ip, err = getMyIPWorker(ipRetrievalURLs)
	return ip, err
}

func getMyIPWorker(urlList []string) (ip string, err error) {
	ch := make(chan string)
	for _, url := range urlList {
		go makeRequest(url, ch)
	}
	for i := 0; i < len(urlList); i++ {
		ip = <-ch
		if ip != "" {
			break
		}
	}
	if ip == "" {
		err = errors.New(
			"none if the fetched urls was able to provide a valid IP address",
		)
	}
	return ip, err
}

func makeRequest(url string, ch chan<- string) {
	var ip string
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err == nil {
		body, _ := ioutil.ReadAll(resp.Body)
		ip = rxFindIP(string(body))
	} else {
		lg.LogError("request failed", logrus.Fields{
			"url": url,
		})
	}
	ch <- ip
}
