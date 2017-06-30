package httpparse

import "github.com/go-yaml/yaml"

type VHostLoader struct {
	
}

func (loader *VHostLoader) Load(configData []byte) []VHostItem {
	vHosts := make([]VHostItem, 1)
	yaml.Unmarshal(configData, &vHosts)
	return vHosts
}
