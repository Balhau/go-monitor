package main

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
