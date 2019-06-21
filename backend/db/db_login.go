package db

import (
	"encoding/json"
	"io/ioutil"
)

// MysqlInfo ...
type MysqlInfo struct {
	Database string `json:"MysqlDatabase"`
	User     string `json:"MysqlUser"`
	Password string `json:"MysqlPassword"`
}

// ReadMysqlInfo Read login credentials for MySQL from a JSON file.
func ReadMysqlInfo(configFile string) MysqlInfo {
	jsonFile, _ := ioutil.ReadFile(configFile)
	dbInfo := MysqlInfo{}
	_ = json.Unmarshal([]byte(jsonFile), &dbInfo)
	return dbInfo
}
