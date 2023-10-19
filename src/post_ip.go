package main

import (
	"bytes"
	"html/template"
	"net/url"
)

func (conf *tConf) execURLTemplate(dns tDNS) (s string) {
	buf := &bytes.Buffer{}
	templ, err := template.New("url").Parse(dns.URL)
	if err == nil {
		templ.Execute(buf, map[string]interface{}{
			"ip":       url.QueryEscape(dns.IPToSend.IP),
			"hostname": url.QueryEscape(dns.Hostname),
			"token":    url.QueryEscape(dns.Token),
		})
		s = buf.String()
	}
	return
}

func (conf *tConf) makeUpdateRequest(dns tDNS) {
	dns.URL = conf.execURLTemplate(dns)
	_ = conf.req(dns.Method, dns.URL)
}
