package main

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
)

//go:embed conf/*
var fs embed.FS

func main() {
	argparse()

	v, _ := fs.ReadDir("/")
	fmt.Printf("%+v\n", v)

	conf := readConf("conf/default.toml")
	var err error
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
			err = updateDNS(conf, currentIPData.IP)
			if err != nil {
				os.Exit(1)
			}
		}
	}
}
