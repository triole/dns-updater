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

func updateDNS(ip string) (err error) {
	url := strings.Replace(requrl, "[IP]", ip, 1)
	lg.Logf("Updating ip at dns service. Request %q", url)
	err = makeUpdateRequest(url)
	return
}

func makeUpdateRequest(url string) (err error) {
	var client http.Client
	var req *http.Request
	var response *http.Response
	req, err = http.NewRequest("GET", url, nil)
	lg.LogIfErr("Error initializing request: %q", err)
	if err == nil {
		req.Header.Add("Authorization", "Basic "+basicAuth(hostname, token))
		response, err = client.Do(req)
		lg.LogIfErr("Error during request: %q", err)

		if response.StatusCode == 200 {
			lg.Logf("Update success. Response code 200")

			defer response.Body.Close()
			bytes, err := ioutil.ReadAll(response.Body)
			lg.LogIfErr("Can not read body: %q", err)

			lg.Logf("Reponse body: %q", string(bytes))

		} else {
			lg.Logf("Update failed. Response code %d", response.StatusCode)
		}
	}
	return
}
