package main

import (
	"io"
	"net/http"
	"regexp"

	"github.com/sirupsen/logrus"
)

func (conf *tConf) req(url string) (string, error) {
	var bytes []byte
	client := http.Client{
		Timeout: conf.RequestsTimeout,
	}
	resp, err := client.Get(url)
	if err == nil {
		bytes, _ = io.ReadAll(resp.Body)
	} else {
		lg.Error("request failed", logrus.Fields{
			"url": url,
		})
	}
	return string(bytes), err
}

func rxFindIP(content string) (r string) {
	r = rxFind(
		"(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9])\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[0-9])",
		content,
	)
	return
}

func rxFind(rx string, content string) (r string) {
	temp, _ := regexp.Compile(rx)
	r = temp.FindString(content)
	return
}
