package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
	"time"
	"strings"
)

type Config struct {
	Domain   string `json:"domain"`
	Hostname string `json:"hostname"`

	Netinfo struct {
		Eth0      string `json:"eth0"`
		Eth1      string `json:"eth1"`
		IpAddr    string `json:"ipAddr"`
		Network   string `json:"network"`
		Netmask   string `json:"netmask"`
		Broadcast string `json:"broadcast"`
		Gateway   string `json:"gateway"`
		Dns       string `json:"dns"`
	}

	Database struct {
		Type     string `json:"type"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	
	Filepaths struct {
		Netinfo string `json:"netinfo"`
	}

} 

func main() {
	//const shortForm = "2006-Jan-02"

	conf := LoadConfig("./config.json")
	t := template.Must(template.New("network.tpl").ParseFiles("network.tpl"))
	
	bakFile := []string{conf.Filepaths.Netinfo, "bak"}
	bakFileAbsPath := strings.Join(bakFile, ".")
	MoveFile(conf.Filepaths.Netinfo, bakFileAbsPath)
	fmt.Println(bakFileAbsPath)
	
	f, err := os.Create(conf.Filepaths.Netinfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	err = t.Execute(f, conf)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func MoveFile(src, dest string) {
	err := os.Rename(src, dest)
	if err != nil {
		fmt.Println("Could not move file: ", err)
		return
	}

	fmt.Println(time.Now())
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
