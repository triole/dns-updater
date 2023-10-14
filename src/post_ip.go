package main

import (
	"bytes"
	"html/template"
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

func (conf *tConf) makeUpdateRequest(dns tDNS) {
	dns.URL = conf.execURLTemplate(dns)
	_ = conf.req(dns.Token, dns.URL, regexIPv4)
}
