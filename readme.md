# DNS-Updater ![build](https://github.com/triole/dns-updater/actions/workflows/build.yaml/badge.svg)

<!--- mdtoc: toc begin -->

1. [Synopsis](#synopsis)
2. [Config](#config)
3. [Help](#help)<!--- mdtoc: toc end -->

## Synopsis

A simple dns update tool that is used to regularly retrieve the current ip and send an update request to a dynamic dns provider. Currently only supports [spdns](https://www.spdyn.de).

## Config

Configuration files are embedded on build for the ease of shipping. Examples are to be found in the `conf` folder.

## Help

```go mdox-exec="r -h"

Send update requests containing the current external ip to a dns service

Flags:
  -h, --help                Show context-sensitive help.
  -j, --info                just display connection information, no dyndns
                            update at all
  -c, --config="default"    config file to use
  -g, --list                list embedded configs
  -f, --force               force update request irrespective of the current ip
  -i, --ip=STRING           use a specific ip to update
  -l, --logfile="/home/ole/.var/log/dns-updater.log"
                            file to process, positional required
  -d, --debug               enable debug output
  -n, --dry-run             do not send update request
  -V, --version-flag        display version
```
