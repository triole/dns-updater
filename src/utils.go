package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

func displayConnectionInformation(conf tConf) {
	ipd := getCurrentIPData(conf)
	fmt.Printf("\nReponse time: %s\n", ipd.Time)
	fmt.Printf("External ip:  %s\n\n", ipd.IP)
}

func rxFindIP(content string) (r string) {
	r = rxFind(
		"(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9])\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[0-9])",
		content,
	)
	return
}

func rxFind(rx string, content string) (r string) {
	temp, _ := regexp.Compile(rx)
	r = temp.FindString(content)
	return
}

func pprint(i interface{}) {
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
}
