package main // import "github.com/y1zhou/love100/backend"

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/y1zhou/love100/backend/config"
	"github.com/y1zhou/love100/backend/db"
	"github.com/y1zhou/love100/backend/router"

	"github.com/jinzhu/gorm"
)

var (
	// Command line arguments
	configFile string
	h          bool
	err        error
)

func main() {
	flag.BoolVar(&h, "h", false, "Display this help information.")
	flag.StringVar(&configFile, "c", "./config.json", "JSON config file.")
	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}

	// Connect to MySQL and auto migrate schema
	configInfo := config.ReadConfigInfo(configFile)
	if configInfo.Database.Client == "mysql" {
		dbLogin := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4,utf8&parseTime=True&loc=Local",
			configInfo.Database.Connection.User, configInfo.Database.Connection.Password, configInfo.Database.Connection.DatabaseName)
		db.DB, err = gorm.Open("mysql", dbLogin)
		if err != nil {
			log.Fatal("Error connecting to database: ", err)
		}
		defer db.DB.Close()
		db.DB = db.DB.Set("gorm:table_options", "CHARSET=utf8mb4")
	} else if configInfo.Database.Client == "sqlite3" {
		db.DB, err = gorm.Open("sqlite3", configInfo.Database.Connection.Filename)
		if err != nil {
			log.Fatal("Error connecting to database: ", err)
		}
	}

	db.DB.AutoMigrate(&db.Contents{}, &db.Users{})

	r := router.SetupRouter(configInfo.CookieSecret)
	r.Run("localhost:" + configInfo.Port)
}
