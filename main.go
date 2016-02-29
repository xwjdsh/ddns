package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
)

var optConf = flag.String("c", "./config.json", "config file")

func main() {
	initConfig(*optConf)
}

func initConfig(configPath string) *Config {
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
	if config.Check() {
		panic("config not complete!")
	}
	return &config
}
