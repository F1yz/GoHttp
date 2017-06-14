package configloader

import (
	"github.com/go-yaml/yaml"
)

type YamlLoader struct {
}

func (yamlLoader *YamlLoader) LoadConfigure(configData []byte) (configMap map[interface{}]interface{}, err error) {
	err = yaml.Unmarshal(configData, &configMap)
	return
}

func (yamlLoader *YamlLoader) SetConfigure(key interface{}, setConfigData interface{}) (err error) {
	return
}