package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"text/template"

	"git.balhau.net/monitor/pkg/dns"
	"gopkg.in/yaml.v2"
)

func main() {
	log.Println("DNS_CONFIG_PATH: ", os.Getenv(dns.DNS_CONFIG_PATH_ENV))

	yamlDNSFilePath := os.Getenv(dns.DNS_CONFIG_PATH_ENV)

	yamlDNSFile, err := ioutil.ReadFile(yamlDNSFilePath)

	if err != nil {
		log.Fatal(err)
		return
	}

	var dnsConfigs dns.DNSBindConfig
	err = yaml.Unmarshal(yamlDNSFile, &dnsConfigs)

	if err != nil {
		log.Fatal(err)
		return
	}

	for _, domain := range dnsConfigs.Domains {
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

		t.Execute(buff, domain)
		log.Println(buff)

	}

}
