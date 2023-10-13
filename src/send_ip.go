package main

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/triole/logseal"
)

func (conf *tConf) execURLTemplate(dns tDNS) (s string) {
	buf := &bytes.Buffer{}
	templ, err := template.New("url").Parse(dns.URL)
	if err == nil {
		templ.Execute(buf, map[string]interface{}{
			"ip":       conf.IPData.Current.IP,
			"hostname": dns.Hostname,
			"token":    dns.Token,
		})
		s = buf.String()
	}
	return
}

func (conf *tConf) makeUpdateRequest(dns tDNS) (err error) {
	var req *http.Request
	var response *http.Response
	client := http.Client{
		Timeout: conf.RequestsTimeout,
	}

	dns.Method = strings.ToUpper(dns.Method)
	dns.URL = conf.execURLTemplate(dns)
	req, err = http.NewRequest(dns.Method, dns.URL, nil)
	lg.IfErrError("can not initialize request", logseal.F{
		"err": err,
	})

	if err == nil {
		lg.Info("fire update request", logseal.F{
			"method": dns.Method,
			"url":    dns.URL,
		})
		response, err = client.Do(req)
		lg.IfErrError(
			"update request failed", logseal.F{
				"err": err,
			})

		if response == nil {
			lg.Error("request response is empty")
			conf.ExitCode = 1
		} else {
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
				lg.Error("update request failed", logseal.F{
					"status_code": response.StatusCode,
				})
				conf.ExitCode = 1
			}
		}
	}
	return
}
