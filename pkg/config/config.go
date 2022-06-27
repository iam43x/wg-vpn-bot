package config

import (
	"flag"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v3"
)

var Config globalConfig

type users struct {
	Allows []string `yaml:"allows"`
	Blocked []string `yaml: "blocked"`
}

type globalConfig struct {
	Users users `yaml:"users"`
    ClientConfig map[string]string `yaml:"clientConfig"`
    ApiToken string `yaml:"apiToken"`
    WgServerConfigPath string `yaml:"wgServerConf"`
    
}

func init() {
    var configPath string
    flag.StringVar(&configPath, "config", "config.yaml", "configuration file")
    flag.Parse()
	yamlFile, err := ioutil.ReadFile(configPath)
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    
    err = yaml.Unmarshal(yamlFile, &Config)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
}