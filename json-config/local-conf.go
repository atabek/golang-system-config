package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
)

type Config struct {
	Domain   string `json:"domain"`
	Hostname string `json:"hostname"`
	Ip       string `json:"ip"`

	Database struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
}

func main() {
	conf := LoadConfig("./config.json")

	fmt.Printf("Domain: %s\n", conf.Domain)
	fmt.Printf("Hostname: %s\n", conf.Hostname)
	fmt.Printf("IP: %s\n", conf.Ip)
	fmt.Printf("DB username: %s\n", conf.Database.Username)
	fmt.Printf("DB password: %s\n", conf.Database.Password)
}

func LoadConfig(file string) Config {
	var config Config
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json.Unmarshal(configFile, &config)

	return config
}
