package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/triole/logseal"
)

var (
	// BUILDTAGS are the ld flags that get injected during the build process
	BUILDTAGS      string
	appName        = "dns-updater"
	appDescription = "Send update requests containing the current external ip to a dns service"
	appMainversion = "0.2"
	lg             logseal.Logseal
)

var CLI struct {
	Config        string `help:"config file to use" default:"${config}" short:"c"`
	Force         bool   `help:"force update request irrespective of the current ip" short:"f"`
	IP            string `help:"use a specific ip to update" short:"p"`
	Timeout       int64  `help:"web requests timeout in seconds" default:"5" short:"t"`
	LogFile       string `help:"log file" default:"${logfile}" short:"l"`
	LogLevel      string `help:"log level" short:"e" default:"info" enum:"debug,info,error"`
	LogNoColors   bool   `help:"disable output colours, print plain text"`
	LogJSON       bool   `help:"enable json log, instead of text one"`
	TestRetrieval bool   `help:"test configured retrieval urls only"`
	DataJSON      string `help:"json file to store ip information to be read in later runs" default:"${datajson}" short:"d"`
	DryRun        bool   `help:"do not send update request" short:"n"`
	VersionFlag   bool   `help:"display version" short:"V"`
}

func parseArgs() {
	homeFolder := getHome()
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
			"logfile":  "stdout",
			"datajson": path.Join(os.TempDir(), "dns-updater.json"),
			"config": returnFirstExistingFile(
				[]string{
					path.Join(getBindir(), appName+".toml"),
					path.Join(homeFolder, ".conf", appName, "conf.yaml"),
					path.Join(homeFolder, ".conf", appName, "conf.toml"),
					path.Join(homeFolder, ".config", appName, "conf.yaml"),
					path.Join(homeFolder, ".config", appName, "conf.toml"),
				},
			),
		},
	)
	_ = ctx.Run()

	if CLI.VersionFlag {
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

func returnFirstExistingFile(arr []string) (s string) {
	for _, el := range arr {
		if isFile(el) {
			s = el
			break
		}
	}
	return
}

func makeAbs(filename string) string {
	filename, err := filepath.Abs(filename)
	if err != nil {
		fmt.Printf("can not assemble absolute filename %q\n", err)
		os.Exit(1)
	}
	return filename
}

func isFile(filePath string) bool {
	stat, err := os.Stat(makeAbs(filePath))
	if !os.IsNotExist(err) && !stat.IsDir() {
		return true
	}
	return false
}

func getBindir() (s string) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	s = filepath.Dir(ex)
	return
}

func getHome() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("unable to retrieve user's home folder")
	}
	return usr.HomeDir
}
