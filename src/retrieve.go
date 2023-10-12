package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func getCurrentIPData(conf tConf) (ipd tIPDataSet, err error) {
	var ip string
	ip, err = getMyIP(conf.Retrieval.URLs)
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
		go makeIPRequest(url, ch)
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

func makeIPRequest(url string, ch chan<- string) {
	var ip string
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err == nil {
		body, _ := io.ReadAll(resp.Body)
		ip = rxFindIP(string(body))
	} else {
		lg.LogError("request failed", logrus.Fields{
			"url": url,
		})
	}
	ch <- ip
}

func makeSimpleRequest(url string) (string, error) {
	var bytes []byte
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err == nil {
		bytes, _ = io.ReadAll(resp.Body)
	} else {
		lg.LogError("request failed", logrus.Fields{
			"url": url,
		})
	}
	return string(bytes), err
}

func displayConnectionInformation(conf tConf) {
	ipd, err := getCurrentIPData(conf)
	if err == nil {
		fmt.Printf("\nReponse time: %s\n", ipd.Time)
		fmt.Printf("External ip:  %s\n", ipd.IP)
	}

	body, err := makeSimpleRequest(conf.Retrieval.TorCheck)
	if err == nil {
		torEnabled := rxFind("You are not using Tor", body) == ""
		fmt.Printf("Tor enabled:  %v\n", torEnabled)
	}

	for _, url := range conf.Retrieval.MoreInfo {
		body, err := makeSimpleRequest(url)
		if err == nil {
			fmt.Printf("\n%s...\n", url)
			fmt.Printf("%s\n", body)
		}
	}

	fmt.Printf("")
}
