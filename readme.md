# DNS-Updater ![build](https://github.com/triole/dns-updater/actions/workflows/build.yaml/badge.svg)

<!-- toc -->

- [Config](#config)

<!-- /toc -->

A simple dns update tool that is used to regularly retrieve the current ip and send an update request to a dynamic dns provider. Currently only supports [spdns](https://www.spdyn.de).

## Config

The configuration is in `toml` format and auto detected if not defined by the `-c` argument. The file is looked for in the following order. The first that exists will be taken.

- binary folder + dns-updater.toml
- ${HOME}/.conf/dns-updater/conf.toml
- ${HOME}/.config/dns-updater/conf.toml

This is how a config looks like:

```go mdox-exec="cat examples/conf.toml"
retrieval_urls = [
  "https://api.ipify.org?format=text",
  "https://www.whatismypublicip.com/ip-lookup/",
  "https://www.showmyip.com/",
  "https://www.expressvpn.com/what-is-my-ip",
  "https://iplocation.com/",
  "https://browserleaks.com/ip",
  "http://icanhazip.com",
  "https://ident.me",
  "http://ipecho.net/plain",
  "https://wtfismyip.com/text",
  "https://jsonip.com",
]

[[dynamic_name_services]]
method = "get"
url = "http://anydns.service/update?user={{.hostname}}&pass={{.token}}&hostname={{.hostname}}&myip={{.ip}}"
hostname = "hostname"
token = "token"
ipv6 = false

[[dynamic_name_services]]
method = "get"
url = "http://anydns.service/update?user={{.hostname}}&pass={{.token}}&hostname={{.hostname}}&myip={{.ip}}"
hostname = "hostname"
token = "token"
ipv6 = true
```

## Help

```go mdox-exec="r -h"

Send update requests containing the current external ip to a dns service

Flags:
  -h, --help                 Show context-sensitive help.
  -c, --config="/home/ole/.conf/dns-updater/conf.toml"
                             config file to use
  -f, --force                force update request irrespective of the current ip
  -p, --ip=STRING            use a specific ip to update
  -t, --timeout=5            web requests timeout in seconds
  -l, --log-file="stdout"    log file
  -e, --log-level="info"     log level
      --log-no-colors        disable output colours, print plain text
      --log-json             enable json log, instead of text one
      --test-retrieval       test configured retrieval urls only
  -d, --data-json="/tmp/dns-updater.json"
                             json file to store ip information to be read in
                             later runs
  -n, --dry-run              do not send update request
  -V, --version-flag         display version
```
