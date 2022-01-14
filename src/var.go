package main

import (
	"olibs/environment"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	// BUILDTAGS are the ld flags that get injected during the build process
	BUILDTAGS      string
	appName        = "dns-updater"
	appMainversion = "0"
	appDescription = "Sends update requests containing the current external ip to a dns service."
	env            = environment.Init(appName, appMainversion, appDescription, BUILDTAGS)

	fileIPDataJSON = "/tmp/dns-updater.json"

	lg Logging

	// dns service settings
	hostname = "***REMOVED***"
	token    = "***REMOVED***"
	requrl   = "***REMOVED***"

	app         = kingpin.New(appName, appDescription)
	argsForce   = app.Flag("force", "force update request irrespective of the current ip").Short('f').Default("False").Bool()
	argsIP      = app.Flag("ip", "use specific up to update").Short('i').String()
	argsLogfile = app.Flag("logfile", "logfile location").Short('l').Default(env.Logfile).String()
)

func argparse() {
	app.Version(env.AppInfoString)
	app.VersionFlag.Short('V')
	app.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))
	// if *argsLogfile != env.Logfile {
	lg = initLogging(*argsLogfile)
	// }
}
