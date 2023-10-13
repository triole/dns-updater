package main

import (
	"encoding/base64"
	"io"
	"net/http"

	"github.com/triole/logseal"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func updateDNS(dns tDNS) (err error) {
	lg.Info("fire update request", logseal.F{
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
	lg.IfErrError("can not initialize request", logseal.F{
		"err": err,
	})

	if err == nil {
		req.Header.Add("Authorization", "Basic "+basicAuth(dns.Hostname, dns.Token))
		response, err = client.Do(req)
		lg.IfErrError(
			"update request failed", logseal.F{
				"err": err,
			})

		if response.StatusCode == 200 {
			lg.Info("request success", logseal.F{
				"status_code": response.StatusCode,
			})

			defer response.Body.Close()
			bytes, err := io.ReadAll(response.Body)
			lg.IfErrError(
				"can not read body", logseal.F{
					"err": err,
				})

			lg.Info("got response", logseal.F{
				"body": string(bytes),
			})

		} else {
			lg.Fatal("update request failed", logseal.F{
				"status_code": response.StatusCode,
			})
		}
	}
	return
}
