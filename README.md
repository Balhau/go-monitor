# go-monitor

## What?

Monitoring collector tools golang. 

## Running

This project uses a `makefile` to automate some of the build steps needed.

To build the binaries you need to run

```shell
make build
```

To clean compilation artifacts run

```shell
make clean
```

If you invoke `make` the default associated target steps are the equivalent of

```shell
make format build
```

## Components

### DnsSpy

This is a cron based dns updater for bind dns daemon. It consumes a [bind.template](resources/dns/templates/bind.template) file, and a [yaml domain file](resources/dns/domains.yml) and creates `bind` dns configuration files. The strategy will issue a bind reboot command.

As example

```go
import (
	"log"
	"os"
	"git.balhau.net/monitor/pkg/dns"
)

const (
	DNS_UPDATER_FREQ = 20
)

func main() {
	var dnsSpy = dns.NewDnsBindSpy(os.Getenv(dns.DNS_CONFIG_PATH_ENV), DNS_UPDATER_FREQ)
	var err = dnsSpy.InitSpy()
	if err != nil {
		log.Fatal(err)
	}
	dnsSpy.StartBlockingSpy()
}
``` 