package main

import (
	"fileoperator"
	"fmt"
	"os"
	"github.com/go-yaml/yaml"
)

type VirtualHost struct {
	WebRoot string `yaml:"web_root"`
	Host string `yaml:"host"`
	Router [] struct{
		Pattern string `yaml:"pattern"`
	}
}

type VirtualHostManager struct {
	VirtualHosts []VirtualHost
}


func (vhManager *VirtualHostManager) SitesAvailable(host string) bool  {
	siteAvailable := false
	for _, vhost := range vhManager.VirtualHosts {
		if vhost.Host == host {
			siteAvailable = true
			break
		}
	}

	return siteAvailable
}


func main() {
	configBytes, err := fileoperator.ReadAll("./vhosts.yaml")

	if err != nil {
		fmt.Println("FUCK")
		os.Exit(-2)
	}

	virtualHostManager := VirtualHostManager{}
	err = yaml.Unmarshal(configBytes, &virtualHostManager)

	if err != nil {
		os.Exit(-3)
	}

	allHosts := virtualHostManager.VirtualHosts


	targetHost := "www.leecode.2o"

	if virtualHostManager.SitesAvailable(targetHost) {
		fmt.Println("YEAH!")
	} else {
		fmt.Println("FUCK")
	}

	for _, host := range allHosts {
		str := fmt.Sprintf("WebRoot: %v, Host: %v, Router: %v", host.WebRoot, host.Host, host.Router)
		fmt.Println(str)
	}

	fmt.Println(virtualHostManager)

}
