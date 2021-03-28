package dns

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"reflect"
	"time"

	"text/template"

	"github.com/go-co-op/gocron"
	"gopkg.in/yaml.v2"
)

const (
	DNS_CONFIG_PATH_ENV string = "DNS_CONFIG_PATH_ENV"
)

type Domain struct {
	TemplatePath   string            `yaml:"templatePath,omitempty"`
	Serial         int               `yaml:"serial,omitempty"`
	Refresh        int               `yaml:"refresh,omitempty"`
	Retry          int               `yaml:"retry,omitempty"`
	Expire         int               `yaml:"expire,omitempty"`
	Minimum        int               `yaml:"minimum,omitempty"`
	Cname          map[string]string `yaml:"cname,omitempty"`
	Address        map[string]string `yaml:"address,omitempty"`
	NameResolver   string            `yaml:"nameresolver,omitempty"`
	Soa            []string          `yaml:"soa,omitempty"`
	Nameservers    []string          `yaml:"nameservers,omitempty"`
	GlobalAddress  []string          `yaml:"globalAddress,omitempty"`
	Mailservers    []string          `yaml:"mailservers,omitempty"`
	Name           string            `yaml:"name,omitempty"`
	OutputPath     string            `yaml:"outputPath"`
	RestartCommand []string          `yaml:"restartCommand,omitempty"`
}

//Dns Bind Template configurations.
type DNSBindConfig struct {
	Domains []Domain `yaml:"domains"`
}

type DnsBindSpy struct {
	configYamlPath string
	spyFrequency   int
	dnsBindConfig  DNSBindConfig
	globalIps      map[string][]string
}

func NewDnsBindSpy(configPath string, spyFreq int) DnsBindSpy {
	return DnsBindSpy{
		configYamlPath: configPath,
		spyFrequency:   spyFreq,
	}
}

func (spy *DnsBindSpy) parseDomains() error {

	yamlDNSFile, err := ioutil.ReadFile(spy.configYamlPath)

	if err != nil {
		log.Fatal(err)
		return err
	}

	var dnsConfigs DNSBindConfig

	err = yaml.Unmarshal(yamlDNSFile, &dnsConfigs)

	if err != nil {
		log.Fatal(err)
		return err
	}

	spy.dnsBindConfig = dnsConfigs
	return nil
}

func (spy *DnsBindSpy) buildIpMap() error {
	if &spy.dnsBindConfig == nil {
		return errors.New("Invalid dns.DNSBindConfig")
	}

	spy.globalIps = make(map[string][]string)

	for _, domain := range spy.dnsBindConfig.Domains {
		log.Println("GlobalIpAddress: ", domain.GlobalAddress)
		spy.globalIps[domain.NameResolver] = domain.GlobalAddress
	}

	return nil
}

func (spy *DnsBindSpy) updatedIps(globalIps map[string][]string) (map[string][]string, error) {
	var updatedIps = make(map[string][]string)
	for name, value := range spy.globalIps {
		log.Println("Resolving DNS: ", name)
		updatedIps[name] = value
		addrs, err := net.LookupHost(name)
		if err != nil {
			log.Println("Error fetching dns: ", name)
		} else {
			updatedIps[name] = addrs
		}
	}
	return updatedIps, nil
}

func (spy *DnsBindSpy) InitSpy() error {
	var err = spy.parseDomains()
	if err != nil {
		return err
	}
	err = spy.buildIpMap()
	if err != nil {
		return err
	}
	return nil
}

func StartBlockingSpy(spy *DnsBindSpy) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(spy.spyFrequency).Second().Do(spy.updateDNS)
	s.StartBlocking()
}

func (spy *DnsBindSpy) updateDNS() {

	var updatedIps, _ = spy.updatedIps(spy.globalIps)

	if !reflect.DeepEqual(spy.globalIps, updatedIps) {
		log.Println("Detected change of ips, updating dns entries")
		log.Println("GlobalIPs: ", spy.globalIps)
		log.Println("UpdatedIPs: ", updatedIps)

		for _, domain := range spy.dnsBindConfig.Domains {
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

			domain.GlobalAddress = updatedIps[domain.NameResolver]

			t.Execute(buff, domain)

			f, err := os.Create(domain.OutputPath)
			if err != nil {
				log.Fatal(err)
				return
			}

			f.Write(buff.Bytes())
			f.Close()

			cmd := exec.Command(domain.RestartCommand[0], domain.RestartCommand[1:]...)

			err = cmd.Run()

			if err != nil {
				log.Fatal(err)
			}
		}
		spy.globalIps = updatedIps
	} else {
		log.Println("No dns changes, skipping")
	}
}
