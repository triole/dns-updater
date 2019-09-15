package main

import (
	"olibs/environment"
	"olibs/logging"
	"os"
	"regexp"
	"strings"
)

var (
	// BUILDTAGS are the ld flags that get injected during the build process
	BUILDTAGS      string
	appName        = "dns-updater"
	appMainversion = "0"
	appDescription = "Sends update requests containing the current external ip tp a dns service."
	env            = environment.Init(appName, appMainversion, appDescription, BUILDTAGS)

	lg = logging.Init(env.Logfile)

	fileIPDataJSON = "/tmp/dns-updater.json"

	// dns service settings
	hostname = "***REMOVED***"
	token    = "***REMOVED***"
	requrl   = "http://update.spdns.de/nic/update?hostname=" + hostname + "&myip=[IP]"

	err error
)

func argparse() {
	rxHelp, _ := regexp.Compile("(-h|--help)")
	rxVersion, _ := regexp.Compile("(-v|--version)")

	s := strings.Join(os.Args[1:], " ")
	if rxHelp.MatchString(s) == true {
		println("\nThere are only two available arguments:\n")
		println("\t -h, --help\tdisplay help")
		println("\t -v, --version\tdisplay version information")
		println("\n")
		os.Exit(0)
	}

	s = strings.Join(os.Args[1:], " ")
	if rxVersion.MatchString(s) == true {
		println(env.AppInfoString)
		os.Exit(0)
	}
}
