package dns

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
