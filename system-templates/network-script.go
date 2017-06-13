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
	t := time.Now()
	datetime := strings.Replace(strings.Split(t.String(), ".")[0], " ", "-", -1)
	
	//conf contains all the data contained in config.json file
	conf := LoadConfig("./config.json")

	// A new template is created and parsed before being used
	tpl := template.Must(template.New("network.tpl").ParseFiles("network.tpl"))
	
	// filenames contains [etc network interfaces] after splitting /etc/network/interfaces
	filenames := strings.Split(conf.Filepaths.Netinfo, "/")
	// filename grabs last element from filenames, appends datetime and bak
	filename  := []string{filenames[len(filenames) - 1], datetime, "bak"}
	
	bakDir  := []string{conf.Filepaths.Netinfo, "bak"}
	//bakFile is a new file name for the backup file with dots between elements
	bakFile := strings.Join(filename, ".")
	
	bakDirAbsPath  := strings.Join(bakDir, ".")
	bakFileJoin := []string{bakDirAbsPath, bakFile}
	bakFileAbsPath := strings.Join(bakFileJoin, "/")
	
	fmt.Println("bakFile:", bakFile)
	fmt.Println("bakDirAbsPath:", bakDirAbsPath)

	// create directory if it does not exits
	if _, err := os.Stat(bakDirAbsPath); os.IsNotExist(err) {
		fmt.Println("Directory does not exist")
		err := os.Mkdir(bakDirAbsPath, os.ModePerm)
		if err != nil {
			fmt.Println("Could not create directory:", err)
			return
		}
	}
	fmt.Printf("%s created\n", bakDirAbsPath)
	// Move original file to a bak directory
	MoveFile(conf.Filepaths.Netinfo, bakFileAbsPath)
	
	// create an empty configuration file
	f, err := os.Create(conf.Filepaths.Netinfo)
	defer f.Close()
	// if there will be an error revert back the old config file
	if err != nil {
		fmt.Println("Could not create file:", err)
		MoveFile(bakFileAbsPath, conf.Filepaths.Netinfo)
	}

	// execute the template with the json data
	err = tpl.Execute(f, conf)
	if err != nil {
		fmt.Println("Could not execute template:", err)
	}
}

func MoveFile(src, dest string) {
	err := os.Rename(src, dest)
	if err != nil {
		fmt.Println("Could not rename file: ", err)
		return
	}
	fmt.Printf("Moved %s to %s\n", src, dest)
}

func LoadConfig(file string) Config {
	var config Config
	configFile, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Could not load config.json:", err.Error())
		os.Exit(1)
	}

	json.Unmarshal(configFile, &config)
	fmt.Println("json file loaded")

	return config
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
