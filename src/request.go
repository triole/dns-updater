package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func updateDNS(conf tConf, ip string) (err error) {
	url := strings.Replace(conf.URL, "[IP]", ip, 1)
	lg.LogInfo("Updating ip at dns service. Request", url)
	err = makeUpdateRequest(conf)
	return
}

func makeUpdateRequest(conf tConf) (err error) {
	var client http.Client
	var req *http.Request
	var response *http.Response
	req, err = http.NewRequest("GET", conf.URL, nil)
	lg.LogError("Error initializing request", err)
	if err == nil {
		req.Header.Add("Authorization", "Basic "+basicAuth(conf.Hostname, conf.Token))
		response, err = client.Do(req)
		lg.LogError("Error during request", err)

		if response.StatusCode == 200 {
			lg.LogInfo("Update success. Response code 200", nil)

			defer response.Body.Close()
			bytes, err := ioutil.ReadAll(response.Body)
			lg.LogError("Can not read body", err)

			lg.LogError("Reponse body", string(bytes))

		} else {
			lg.LogFatal("Update failed. Response code", response.StatusCode)
		}
	}
	return
}
