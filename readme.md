# DNS-Updater ![build](https://github.com/triole/dns-updater/actions/workflows/build.yaml/badge.svg)

<!-- toc -->

- [Config](#config)
- [Help](#help)

<!-- /toc -->

A simple dns update tool that is used to regularly retrieve the current ip and send an update request to a dynamic dns provider. Currently only supports [spdns](https://www.spdyn.de).

## Config

The configuration is auto detected if not defined by the `-c` argument. The file is looked for in the following order. The first that exists will be taken.

- binary folder + dns-updater.toml
- ${HOME}/.conf/dns-updater/conf.yaml
- ${HOME}/.conf/dns-updater/conf.toml
- ${HOME}/.config/dns-updater/conf.yaml
- ${HOME}/.config/dns-updater/conf.toml

Config examples can be found in the `examples` folder.

## Help

```go mdox-exec="r -h"

Send update requests containing the current external ip to a dns service

Flags:
  -h, --help            Show context-sensitive help.
  -j, --info            just display connection information, no dyndns update at
                        all
  -c, --config="/home/ole/.conf/dns-updater/conf.toml"
                        config file to use
  -g, --list            list embedded configs
  -f, --force           force update request irrespective of the current ip
  -i, --ip=STRING       use a specific ip to update
  -l, --logfile="/home/ole/.var/log/dns-updater.log"
                        file to process, positional required
  -d, --debug           enable debug output
  -n, --dry-run         do not send update request
  -V, --version-flag    display version
```
