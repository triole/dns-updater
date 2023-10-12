package main

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func updateDNS(dns tDNS) (err error) {
	lg.LogInfo("fire update request", logrus.Fields{
		"url": dns.URL,
	})
	err = makeUpdateRequest(dns)
	return
}

func makeUpdateRequest(dns tDNS) (err error) {
	var client http.Client
	var req *http.Request
	var response *http.Response
	req, err = http.NewRequest("GET", dns.URL, nil)

	lg.LogIfError(err, "error initializing request", logrus.Fields{
		"err": err,
	})
	if err == nil {
		req.Header.Add("Authorization", "Basic "+basicAuth(dns.Hostname, dns.Token))
		response, err = client.Do(req)
		lg.LogIfError(
			err, "error during request", logrus.Fields{
				"err": err,
			})

		if response.StatusCode == 200 {
			lg.LogInfo("request success", logrus.Fields{
				"status_code": response.StatusCode,
			})

			defer response.Body.Close()
			bytes, err := io.ReadAll(response.Body)
			lg.LogIfError(
				err, "can not read body", logrus.Fields{
					"err": err,
				})

			lg.LogInfo("got response", logrus.Fields{
				"body": string(bytes),
			})

		} else {
			lg.LogFatal("update request failed", logrus.Fields{
				"status_code": response.StatusCode,
			})
		}
	}
	return
}
