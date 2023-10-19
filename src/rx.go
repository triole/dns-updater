package main

import (
	"regexp"
)

var (
	regexIPMatchers = map[string]string{
		"v4": "(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9])\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[0-9])",
		"v6": "([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}",
	}
)

func rxFind(rx string, str string) (r string) {
	temp, _ := regexp.Compile(rx)
	r = temp.FindString(str)
	return
}

func rxFindIPv4(str string) (r string) {
	return rxFind(regexIPMatchers["v4"], str)
}

func rxFindIPv6(str string) (r string) {
	return rxFind(regexIPMatchers["v6"], str)
}

func isIPv6(str string) bool {
	re, _ := regexp.Compile(regexIPMatchers["v6"])
	return re.MatchString(str)
}

func isIPv4(str string) bool {
	re, _ := regexp.Compile(regexIPMatchers["v4"])
	return re.MatchString(str)
}

func isValidIP(str string) (b bool) {
	b = isIPv4(str)
	if !b {
		b = isIPv6(str)
	}
	return
}
