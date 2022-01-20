package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/sirupsen/logrus"
)

var (
	// BUILDTAGS are the ld flags that get injected during the build process
	BUILDTAGS      string
	appName        = "dns-updater"
	appDescription = "Send update requests containing the current external ip to a dns service"
	appMainversion = "0.1"

	fileIPDataJSON = "/tmp/dns-updater.json"

	lg Logging
)

var CLI struct {
	Info        bool   `help:"just display connection information, no dyndns update at all" short:j`
	Config      string `help:"config file to use" short:c default:default`
	List        bool   `help:"list embedded configs" short:l`
	Force       bool   `help:"force update request irrespective of the current ip" short:f`
	IP          string `help:"use a specific ip to update" short:i`
	Logfile     string `help:"file to process, positional required" default:${logfile}`
	Debug       bool   `help:"enable debug output" short:d`
	DryRun      bool   `help:"do not send update request" short:n`
	VersionFlag bool   `help:"display version" short:V`
}

func parseArgs() {
	user, err := user.Current()
	if err != nil {
		lg.LogFatal("unable to detect user", logrus.Fields{
			"err": err,
		})
	}
	homeDir := user.HomeDir
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:      true,
			Summary:      true,
			NoAppSummary: true,
			FlagsLast:    false,
		}),
		kong.Vars{
			"logfile": path.Join(homeDir, ".var", "log", appName+".log"),
		},
	)
	_ = ctx.Run()

	if CLI.VersionFlag == true {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
	// ctx.FatalIfErrorf(err)
}

type tPrinter []tPrinterEl
type tPrinterEl struct {
	Key string
	Val string
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "version: "+appMainversion+".", -1)
	fmt.Printf("\n%s\n%s\n\n", appName, appDescription)
	arr := strings.Split(s, "\n")
	var pr tPrinter
	var maxlen int
	for _, line := range arr {
		if strings.Contains(line, ":") {
			l := strings.Split(line, ":")
			if len(l[0]) > maxlen {
				maxlen = len(l[0])
			}
			pr = append(pr, tPrinterEl{l[0], strings.Join(l[1:], ":")})
		}
	}
	for _, el := range pr {
		fmt.Printf("%"+strconv.Itoa(maxlen)+"s\t%s\n", el.Key, el.Val)
	}
	fmt.Printf("\n")
}
