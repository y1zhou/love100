package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ConfigInfo ...
type ConfigInfo struct {
	Database     string `json:"MysqlDatabase"`
	User         string `json:"MysqlUser"`
	Password     string `json:"MysqlPassword"`
	CookieSecret string `json:"CookieSecret"`
	Port         string `json:"BackendPort"`
}

// ReadConfigInfo Read config information from a JSON file.
func ReadConfigInfo(configFile string) ConfigInfo {
	jsonFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("File '" + configFile + "' not found!\n")
	}
	dbInfo := ConfigInfo{}
	_ = json.Unmarshal([]byte(jsonFile), &dbInfo)
	return dbInfo
}
