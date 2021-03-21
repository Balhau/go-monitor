package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"os"
	"reflect"
	"time"

	"text/template"

	"git.balhau.net/monitor/pkg/dns"
	"github.com/go-co-op/gocron"
	"gopkg.in/yaml.v2"
)

const (
	DNS_UPDATER_FREQ = 2
)

var domains, erra = parseDomains()
var globalips, errb = buildGlobalIPsMap(*domains)

func main() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(DNS_UPDATER_FREQ).Second().Do(updateDNS)
	s.StartBlocking()
}

func testCron() {
	log.Println("Task executed")
	log.Println(globalips)
	var updatedIps, _ = updatedIps(globalips)
	log.Println(updatedIps)
}

func updatedIps(globalIps map[string][]string) (map[string][]string, error) {
	var updatedIps = make(map[string][]string)
	for name, _ := range globalIps {
		addrs, err := net.LookupHost(name)
		if err != nil {
			return nil, err
		}
		updatedIps[name] = addrs
	}
	return updatedIps, nil
}

func buildGlobalIPsMap(domains dns.DNSBindConfig) (map[string][]string, error) {
	log.Println(domains)
	if &domains == nil {
		return nil, errors.New("Invalid dns.DNSBindConfig")
	}

	var globalips = make(map[string][]string)

	for _, domain := range domains.Domains {
		globalips[domain.Name] = domain.GlobalAddress
	}

	return globalips, nil
}

func parseDomains() (*dns.DNSBindConfig, error) {
	yamlDNSFilePath := os.Getenv(dns.DNS_CONFIG_PATH_ENV)

	yamlDNSFile, err := ioutil.ReadFile(yamlDNSFilePath)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var dnsConfigs dns.DNSBindConfig
	err = yaml.Unmarshal(yamlDNSFile, &dnsConfigs)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &dnsConfigs, nil
}

func updateDNS() {

	var updatedIps, _ = updatedIps(globalips)

	if !reflect.DeepEqual(globalips, updatedIps) {
		log.Println("Detected change of ips, updating dns entries", updatedIps)
		for _, domain := range domains.Domains {
			domainTemplate, err := ioutil.ReadFile(domain.TemplatePath)
			if err != nil {
				log.Fatal(err)
				return
			}

			t, err := template.New(domain.Name).Parse(string(domainTemplate))

			if err != nil {
				log.Fatal(err)
				return
			}

			buff := bytes.NewBufferString("")

			domain.GlobalAddress = updatedIps[domain.Name]

			t.Execute(buff, domain)

			f, err := os.Create(domain.OutputPath)
			if err != nil {
				log.Fatal(err)
				return
			}

			f.Write(buff.Bytes())
			f.Close()
		}
		globalips = updatedIps
	} else {
		log.Println("No dns changes, skipping")
	}
}
