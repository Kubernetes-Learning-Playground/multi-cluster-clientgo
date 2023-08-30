package config

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"golanglearning/new_project/multi_cluster_client/pkg/model"

	"io/ioutil"
	"log"
)

// TODO: 配置文件

var SysConfig *Config

type Config struct {
	Clusters       []model.Cluster `json:"clusters" yaml:"clusters"`
}

func NewConfig() *Config {
	return &Config{}
}

func loadConfigFile(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	return b
}

func BuildConfig(path string) (*Config, error) {
	config := NewConfig()
	if b := loadConfigFile(path); b != nil {

		err := yaml.Unmarshal(b, config)
		if err != nil {
			return nil, err
		}
		return config, err
	} else {
		return nil, fmt.Errorf("load config file error...")
	}

}