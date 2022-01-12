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
		lg.Log("None of the made requests did return a valid IP address.")
		os.Exit(1)
	} else {
		b := didIPChange(currentIPData)
		if b == true || *argsForce == true {
			if *argsForce == true {
				lg.Log("Force update request irrespective of the current ip")
			}
			writeIPDataJSON(currentIPData)
			err = updateDNS(currentIPData.IP)
			if err != nil {
				os.Exit(1)
			}
		}
	}
}
