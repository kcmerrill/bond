[![Build Status](https://travis-ci.org/kcmerrill/bond.svg?branch=master)](https://travis-ci.org/kcmerrill/bond)

![bond](bond.jpg)

# Bond

Bond your application output to external commands. Using regular expression matching, execute commands based on given matches.

A few use cases:
* Metrics gathering
* Alerting
* Queue Worker

# Quick Setup

Create a `bond.yml` file with key value pairs. The keys are `regular expressions`, values are `bash commands` to execute. `:match` is the string that matched, so you can pass it along to your script

```yaml
"referrer: http://().baddomain.com$": echo -n "baddomain.com:60|c|:match" | nc -4u -w0 127.0.0.1 8125
"^#id: (\d+): |
    echo "Processing id: :match"
    curl -XPOST http://crush.kcmerrill.com/id/:match
```

```shell
$ tail -f myapplication.log | bond
```

## Binaries || Installation

[![MacOSX](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/apple_logo.png "Mac OSX")](http://go-dist.kcmerrill.com/kcmerrill/bond/mac/amd64) [![Linux](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/linux_logo.png "Linux")](http://go-dist.kcmerrill.com/kcmerrill/bond/linux/amd64)

via golang:

`$ go get -u github.com/kcmerrill/bond`
