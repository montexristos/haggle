package models

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type SiteConfig struct {
	Id      string            `json:"id" yaml:"id"`
	BaseUrl string            `json:"baseUrl" yaml:"baseUrl"`
	Urls    map[string]string `json:"urls"`
	Parser  string            `json:"parser" yaml:"parser"`
}

func ParseSiteConfig(website string) SiteConfig {
	yamlFile, err := ioutil.ReadFile(fmt.Sprintf("config/%s.yaml", website))
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	y := SiteConfig{}
	err = yaml.Unmarshal(yamlFile, &y)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(y)
	return y
}
