package main

import (
	"flag"
	"fmt"
	"github.com/prismatik/config/base"
	"github.com/prismatik/config/buildkite"
	"github.com/prismatik/config/codescreen"
	"github.com/prismatik/config/docker"
	"github.com/prismatik/config/elastic"
	"github.com/prismatik/config/influxdb"
	"github.com/prismatik/config/postgres"
	"github.com/prismatik/config/rethinkdb"
	"github.com/prismatik/config/ufw"
	"strings"
)

func main() {
	roles := flag.String("role", "base", "A comma separated list of roles to execute on the machine")

	flag.Parse()

	parsedRoles := strings.Split(*roles, ",")

	for _, role := range parsedRoles {
		fmt.Println("Executing role", role)
		switch role {
		case "base":
			base.Go()
		case "influxdb":
			influxdb.Go()
		case "postgres":
			postgres.Go()
		case "rethinkdb":
			rethinkdb.Go()
		case "ufw":
			ufw.Go()
		case "codescreen":
			codescreen.Go()
		case "buildkite":
			buildkite.Go()
		case "docker":
			docker.Go()
		case "elastic":
			elastic.Go()
		}
	}
}
