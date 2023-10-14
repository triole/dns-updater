package main

import (
	"regexp"
)

var (
	// TODO: integrate ipv6 regex later, also think of conf entry
	rxIPAdresses = []string{
		"(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9])\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[0-9])",
		// "([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}",
	}
)

func rxFindByList(matchers []string, content string) (r string) {
	for _, rx := range matchers {
		r = rxFind(rx, content)
		if r != "" {
			break
		}
	}
	return
}

func rxFind(rx string, content string) (r string) {
	temp, _ := regexp.Compile(rx)
	r = temp.FindString(content)
	return
}
