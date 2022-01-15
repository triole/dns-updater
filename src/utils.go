package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func didIPChange(currentIPData IPData) (b bool) {
	b = false
	oldIPData := readIPDataJSON()
	if currentIPData.IP != oldIPData.IP {
		b = true
		lg.LogInfo(
			"IP changed", map[string]string{
				"oldIP":     oldIPData.IP,
				"currentIP": currentIPData.IP,
			},
		)
	} else {
		t := oldIPData.Time
		lg.LogInfo("IP stable since", t.Format("Monday, Jan _2 15:04:05 2006"))
	}
	return
}

func getCurrentIPData(conf tConf) (ipd IPData) {
	ip, err := getMyIP(conf.IpRetrievalURLs)
	if err == nil {
		ipd = IPData{
			Time: time.Now(),
			IP:   ip,
		}
		lg.LogInfo("Current IP is", ipd.IP)
	}
	return
}

// GetMyIP is the main function used to call the ip retrieval methods
func getMyIP(ipRetrievalURLs []string) (ip string, err error) {
	ip, err = getMyIPWorker(ipRetrievalURLs)
	return ip, err
}

// GetMyIP fetches the external ip and returns it, channeled and asynchronously
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
		err = errors.New("none if the fetched urls was able to provide a valid IP address")
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
		ip = rxFind(
			`(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])`,
			string(body),
		)
	} else {
		println("An error occured during the request to " + url)
	}
	ch <- ip
}

// Find returns the substring that matches the given regex scheme
func rxFind(rx string, content string) (r string) {
	temp, _ := regexp.Compile(rx)
	r = temp.FindString(content)
	return
}
