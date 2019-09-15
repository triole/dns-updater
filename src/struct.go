package main

import "time"

// IPData contains the information that the IP fetch returned
type IPData struct {
	Time time.Time
	IP   string
}
