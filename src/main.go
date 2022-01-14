package main

import (
	"os"
)

func main() {
	var err error
	argparse()

	var currentIPData IPData

	if *argsIP != "" {
		currentIPData.IP = *argsIP
		*argsForce = true
	} else {
		currentIPData = getCurrentIPData()
	}

	if currentIPData.IP == "" {
		lg.LogError("None of the made requests did return a valid IP address.", nil)
		os.Exit(1)
	} else {
		b := didIPChange(currentIPData)
		if b == true || *argsForce == true {
			if *argsForce == true {
				lg.LogInfo("Force update request irrespective of the current ip", nil)
			}
			writeIPDataJSON(currentIPData)
			err = updateDNS(currentIPData.IP)
			if err != nil {
				os.Exit(1)
			}
		}
	}
}
