package main

import "os"

func main() {
	argparse()

	currentIPData := getCurrentIPData()

	if currentIPData.IP == "" {
		println("None of the made requests did return a valid IP address.")
		os.Exit(1)
	} else {
		b := didIPChange(currentIPData)
		if b == true || *argsForce == true {
			if *argsForce == true {
				lg.Log("Force update request irrespective of the current ip")
			}
			writeIPDataJSON(currentIPData)
			updateDNS(currentIPData.IP)
		}
	}
}
