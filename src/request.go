package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/triole/logseal"
)

type tRequestResponse struct {
	Method       string
	URL          string
	Status       int
	Body         string
	ResponseTime time.Duration
	Errors       []error
}

func (conf *tConf) req(method, url string) (rr tRequestResponse) {
	var bytes []byte
	var err error
	var req *http.Request
	var response *http.Response

	client := http.Client{
		Timeout: conf.RequestsTimeout,
	}

	rr.URL = url
	rr.Method = strings.ToUpper(method)

	if err == nil {
		lg.Debug("fire request", logseal.F{
			"method": rr.Method,
			"url":    rr.URL,
		})

		req, err = http.NewRequest(rr.Method, rr.URL, nil)
		if err != nil {
			rr.Errors = append(rr.Errors, err)
			lg.Error("can not initialize request", conf.reqFields(rr))
		}
		start := time.Now()
		response, err = client.Do(req)
		rr.ResponseTime = time.Since(start)

		if err != nil {
			rr.Errors = append(rr.Errors, err)
			lg.Error("request failed", conf.reqFields(rr))
		}

		if response == nil {
			rr.Errors = append(rr.Errors, errors.New("response is empty"))
			lg.Error("request response is empty", conf.reqFields(rr))
		} else {
			rr.Status = response.StatusCode
			if rr.Status == 200 {
				defer response.Body.Close()
				bytes, err = io.ReadAll(response.Body)
				if err == nil {
					rr.Body = string(bytes)
				} else {
					rr.Errors = append(rr.Errors, err)
					lg.IfErrError("can not read body", conf.reqFields(rr))
				}
			} else {
				rr.Errors = append(rr.Errors, errors.New("status code not 200"))
			}

			fields := conf.reqFields(rr)
			if rr.Body != "" {
				lg.Info("request success", fields)
			} else {
				if lg.Logrus.Level > 4 {
					fields["body"] = rr.Body
				}
				lg.Error("request response body empty", fields)
			}
		}
	}
	return
}

func (conf *tConf) reqFields(rr tRequestResponse) logseal.F {
	return logseal.F{
		"method":        rr.Method,
		"url":           rr.URL,
		"status":        rr.Status,
		"errors":        rr.Errors,
		"response_time": rr.ResponseTime,
	}
}
