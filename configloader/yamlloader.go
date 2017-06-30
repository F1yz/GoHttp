package configloader

import (
	"github.com/go-yaml/yaml"
)

type ConfigureLoader interface {
	SetConfigure(key string, setConfigData interface{}) error
	LoadConfigure(configData []byte) (map[string]interface{}, error)
}

type YamlLoader struct {
}

func (yamlLoader *YamlLoader) LoadConfigure(configData []byte) (configMap map[string]interface{}, err error) {
	err = yaml.Unmarshal(configData, &configMap)
	return
}

func (yamlLoader *YamlLoader) SetConfigure(key string, setConfigData interface{}) (err error) {
	return
}