package main

import (
	"encoding/json"

	"io/ioutil"
	"log"
)

type Config struct {
	LocalHost  string
	LocalPort  string
	RemoteHost string
	RemotePort string

	Socks5Host string
	Socks5Port string

	Key string
}

func getConfig(configPath string, config *Config) {

	// 读取配置文件
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	// 解析配置文件

	err = json.Unmarshal(data, config)
	if err != nil {
		log.Fatal(err)
	}

}
