package main

import (
	"fileoperator"
	"fmt"
	"os"
	"github.com/go-yaml/yaml"
	"vhosts"
)



func main() {
	configBytes, err := fileoperator.ReadAll("./vhosts.yaml")

	if err != nil {
		fmt.Println("FUCK")
		os.Exit(-2)
	}

	virtualHostManager := vhosts.VirtualHostManager{}
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

		routerStr := ""
		for _, route := range host.Router {
			routerStr  += route.String()
		}
		str := fmt.Sprintf("WebRoot: %v, Host: %v, Router: %v", host.WebRoot, host.Host, routerStr)
		fmt.Println(str)
	}

	fmt.Println(virtualHostManager)

}
