package vhosts

import (
	"fileoperator"
	"fmt"
	"os"
	"github.com/go-yaml/yaml"
)


type VirtualHostManager struct {
	VirtualHosts []VirtualHost
}

func (vhManager *VirtualHostManager) LoadVirtualHosts() {
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
