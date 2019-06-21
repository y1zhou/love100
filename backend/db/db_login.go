package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// MysqlInfo ...
type MysqlInfo struct {
	Database string `json:"MysqlDatabase"`
	User     string `json:"MysqlUser"`
	Password string `json:"MysqlPassword"`
}

// ReadMysqlInfo Read login credentials for MySQL from a JSON file.
func ReadMysqlInfo(configFile string) MysqlInfo {
	jsonFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("File '" + configFile + "' not found!\n")
	}
	dbInfo := MysqlInfo{}
	_ = json.Unmarshal([]byte(jsonFile), &dbInfo)
	return dbInfo
}
