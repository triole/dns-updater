package main

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/triole/logseal"
)

type tRequestResponse struct {
	Method string
	URL    string
	Match  string
	Status int
	Body   string
	Errors []error
}

func (conf *tConf) req(method, url string, matchers []string) (rr tRequestResponse) {
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
			lg.Error("can not initialize request", reqFields(rr))
		}

		response, err = client.Do(req)
		if err != nil {
			rr.Errors = append(rr.Errors, err)
			lg.Error("request failed", reqFields(rr))
		}

		if response == nil {
			rr.Errors = append(rr.Errors, errors.New("response is empty"))
			lg.Error("request response is empty", reqFields(rr))
		} else {
			rr.Status = response.StatusCode
			if rr.Status == 200 {
				defer response.Body.Close()
				bytes, err = io.ReadAll(response.Body)
				if err != nil {
					rr.Errors = append(rr.Errors, err)
					lg.IfErrError("can not read body", reqFields(rr))
				} else {
					rr.Body = string(bytes)
					if err == nil {
						rr.Match = rxFindByList(rxIPAdresses, rr.Body)
						if rr.Match == "" {
							rr.Errors = append(rr.Errors, errors.New("regex did not match"))
						}
					}
				}
			} else {
				rr.Errors = append(rr.Errors, errors.New("status code not 200"))
			}

			fields := reqFields(rr)
			if rr.Match != "" {
				lg.Info("request success", fields)
			} else {
				if lg.Logrus.Level > 4 {
					fields["body"] = rr.Body
					fields["matchers"] = matchers
				}
				lg.Error("request fail", fields)
			}
		}
	}
	return
}

func reqFields(rr tRequestResponse) logseal.F {
	return logseal.F{
		"method": rr.Method,
		"url":    rr.URL,
		"match":  rr.Match,
		"status": rr.Status,
		"errors": rr.Errors,
	}
}
