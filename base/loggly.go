package base

import (
	"github.com/prismatik/jabba"
	"github.com/prismatik/secrets"
)

var logglySecrets = map[string]string{
	"apiKey": secrets.Get("loggly", "apiKey"),
}

var logglyFile = jabba.File{
	Path: "/etc/rsyslog.d/22-loggly.conf",
	Perm: 0644,
	Vars: logglySecrets,
	Template: `# Setup disk assisted queues
$WorkDirectory /var/spool/rsyslog # where to place spool files
$ActionQueueFileName fwdRule1     # unique name prefix for spool files
$ActionQueueMaxDiskSpace 1g       # 1gb space limit (use as much as possible)
$ActionQueueSaveOnShutdown on     # save messages to disk on shutdown
$ActionQueueType LinkedList       # run asynchronously
$ActionResumeRetryCount -1        # infinite retries if host is down

template(name="LogglyFormat" type="string"
 string="<%pri%>%protocol-version% %timestamp:::date-rfc3339% %HOSTNAME% %app-name% %procid% %msgid% [{{.apiKey}} tag=\"syslog\"] %msg%\n")

# Send messages to Loggly over TCP using the template.
action(type="omfwd" protocol="tcp" target="logs-01.loggly.com" port="514" template="LogglyFormat")`,
}
