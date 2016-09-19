package rethinkdb

import (
	"github.com/prismatik/jabba"
)

var distro = "trusty"

func Go() {
	jabba.RunOrDie("echo", "'deb http://download.rethinkdb.com/apt `lsb_release -cs` main'", "|", "sudo", "tee", "/etc/apt/sources.list.d/rethinkdb.list")
	distro = jabba.DistroString()
	jabba.WriteFile(rethinkSource)
	jabba.RunOrDie("wget", "-qO-", "https://download.rethinkdb.com/apt/pubkey.gpg", "|", "sudo", "apt-key", "add", "-")
	jabba.RunOrDie("sudo", "apt-get", "update")
	jabba.RunOrDie("sudo", "apt-get", "install", "rethinkdb")
	jabba.WriteFile(rethinkConf)
	jabba.RunOrDie("sudo", "/etc/init.d/rethinkdb", "restart")

var rethinkSource = jabba.File{
	Path: "/etc/apt/sources.list.d/rethinkdb.list",
	Perm: 0644,
	Vars: map[string]string{
		"distro": distro,
	},
	Template: `deb http://download.rethinkdb.com/apt {{.distro}} main`,
}

var rethinkConf = jabba.File{
	Path: "/etc/rethinkdb/instances.d/default.conf",
	Perm: 0644,
	Template: `#
# RethinkDB instance configuration sample
#
# - Give this file the extension .conf and put it in /etc/rethinkdb/instances.d in order to enable it.
# - See http://www.rethinkdb.com/docs/guides/startup/ for the complete documentation
# - Uncomment an option to change its value.
#

###############################
## RethinkDB configuration
###############################

### Process options

## User and group used to run rethinkdb
## Command line default: do not change user or group
## Init script default: rethinkdb user and group
# runuser=rethinkdb
# rungroup=rethinkdb

## Stash the pid in this file when the process is running
## Command line default: none
## Init script default: /var/run/rethinkdb/<name>/pid_file (where <name> is the name of this config file without the extension)
# pid-file=/var/run/rethinkdb/rethinkdb.pid

### File path options

## Directory to store data and metadata
## Command line default: ./rethinkdb_data
## Init script default: /var/lib/rethinkdb/<name>/ (where <name> is the name of this file without the extension)
# directory=/var/lib/rethinkdb/default

## Log file options
## Default: <directory>/log_file
# log-file=/var/log/rethinkdb

### Network options

## Address of local interfaces to listen on when accepting connections
## May be 'all' or an IP address, loopback addresses are enabled by default
## Default: all local addresses
bind=all

## Address that other rethinkdb instances will use to connect to this server.
## It can be specified multiple times
# canonical-address=

## The port for rethinkdb protocol for client drivers
## Default: 28015 + port-offset
# driver-port=28015

## The port for receiving connections from other nodes
## Default: 29015 + port-offset
# cluster-port=29015

## The host:port of a node that rethinkdb will connect to
## This option can be specified multiple times.
## Default: none
# join=example.com:29015
join=rethinkmaster.prismatik.com.au:29015

## All ports used locally will have this value added
## Default: 0
# port-offset=0

## r.http(...) queries will use the given server as a web proxy
## Default: no proxy
# reql-http-proxy=socks5://example.com:1080

### Web options

## Port for the http admin console
## Default: 8080 + port-offset
# http-port=8080

## Disable web administration console
# no-http-admin

### CPU options

## The number of cores to use
## Default: total number of cores of the CPU
# cores=2

### Memory options

## Size of the cache in MB
## Default: Half of the available RAM on startup
# cache-size=1024

### Disk

## How many simultaneous I/O operations can happen at the same time
# io-threads=64

## Enable direct I/O
# direct-io

### Meta

## The name for this server (as will appear in the metadata).
## If not specified, it will be randomly chosen from a short list of names.
# server-name=server1
`,
}
