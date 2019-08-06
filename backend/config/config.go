package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Info ...
type Info struct {
	Database struct {
		Client     string `json:"Client"`
		Connection struct {
			User         string `json:"User"`
			Password     string `json:"Password"`
			DatabaseName string `json:"DatabaseName"`
			Filename     string `json:"filename"`
		} `json:"Connection"`
	} `json:"Database"`
	CookieSecret string `json:"CookieSecret"`
	Port         string `json:"BackendPort"`
}

// ReadConfigInfo Read config information from a JSON file.
func ReadConfigInfo(configFile string) Info {
	jsonFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("File '" + configFile + "' not found!\n")
	}
	dbInfo := Info{}
	_ = json.Unmarshal([]byte(jsonFile), &dbInfo)
	return dbInfo
}
