package configloader

import (
	"github.com/go-yaml/yaml"
)

type ConfigureLoader interface {
	LoadConfigure(configData []byte) (map[interface{}]interface{}, error)
	SetConfigure(key interface{}, setConfigData interface{}) error
}

type YamlLoader struct {
}

func (yamlLoader *YamlLoader) LoadConfigure(configData []byte) (configMap map[interface{}]interface{}, err error) {
	err = yaml.Unmarshal(configData, &configMap)
	return
}

func (yamlLoader *YamlLoader) SetConfigure(key interface{}, setConfigData interface{}) (err error) {
	return
}