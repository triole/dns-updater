package main

import (
	"regexp"
)

// Find returns the substring that matches the given regex scheme
func rxFind(rx string, content string) (r string) {
	temp, _ := regexp.Compile(rx)
	r = temp.FindString(content)
	return
}
