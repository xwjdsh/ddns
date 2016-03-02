package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var optConf = flag.String("c", "./config.json", "config file")

func main() {
	ddns, err := newConfig(*optConf).newDDNS()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ddns.initHttp()

	domainId, err := ddns.domainID()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	recordId, recordIp, err := ddns.recordID(domainId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(domainId, recordId)
	ip, err := currentIP()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ip", ip, recordIp)

}

func newConfig(configPath string) *Config {
	if _, err := os.Stat(configPath); err != nil && os.IsNotExist(err) {
		panic("can't find config file!")
	}
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic("read file error!")
	}
	config := Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic("deserialize json error!")
	}
	if config.checked {
		panic("config not complete!")
	}
	return &config
}
