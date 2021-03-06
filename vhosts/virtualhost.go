package vhosts

type VirtualHost struct {
	WebRoot string `yaml:"web_root"`
	Host string `yaml:"host"`
	Router [] struct{
		Pattern string `yaml:"pattern"`
	}
}