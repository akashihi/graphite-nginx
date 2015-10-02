# graphite-nginx

## What is this?

Nginx status to graphite gateway.

## Building it

1. Install [go](http://golang.org/doc/install)

2. Install "graphite-golang" go get -u github.com/marpaia/graphite-golang

2. Install "go-logging" go get -u github.com/op/go-logging

4. Compile graphite-nginx

        git clone git://github.com/akashihi/graphite-nginx.git
        cd graphite-nginx
        go build .

## Running it

Generally:

    conntrack-logger -url https://localhost/server_status -metrics-host 192.168.1.1 -metrics-port 2003 -metrics-prefix test -period 60

All parameters could be omited. Run with --help to het parameters description

## License 

See LICENSE file.

Copyright 2015 Denis V Chapligin <akashihi@gmail.com>
