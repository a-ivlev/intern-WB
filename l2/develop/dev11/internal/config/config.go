package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v3"
)

var (
	ErrLoadConf  = errors.New("load config error")
	ErrParseConf = errors.New("parse config error")
)

type Configuration struct {
	SrvHost string `yaml:"srv-host"`
	SrvPort string `yaml:"srv-port"`
}

func LoadConfig(source string) (*Configuration, error) {
	conf := Configuration{
		"localhost",
		"9000",
	}
	f, err := ioutil.ReadFile(source)
	if err != nil {
		return nil, fmt.Errorf("%s: %s\n", ErrLoadConf, err.Error())
	}

	switch path.Ext(source) {
	case ".yml":
		fallthrough
	case ".yaml":
		err = yaml.Unmarshal(f, &conf)
		if err != nil {
			return nil, fmt.Errorf("%s: %s\n", ErrParseConf, err.Error())
		}
	}

	return &conf, nil
}
