package main // import "github.com/y1zhou/love100/backend"

import (
	"flag"
	"fmt"
	"os"

	"github.com/y1zhou/love100/backend/db"
	"github.com/y1zhou/love100/backend/router"

	"github.com/jinzhu/gorm"
)

var (
	// Command line arguments
	configFile string
	h          bool
	port       string
)

func main() {
	flag.BoolVar(&h, "h", false, "Display this help information.")
	flag.StringVar(&configFile, "c", "./config.json", "JSON config file.")
	flag.StringVar(&port, "p", "9423", "The port to listen to.")
	flag.Parse()
	if h {
		flag.Usage()
		os.Exit(0)
	}

	// Connect to MySQL and auto migrate schema
	dbInfo := db.ReadMysqlInfo(configFile)
	dbLogin := fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbInfo.User, dbInfo.Password, dbInfo.Database)
	db.DB, _ = gorm.Open("mysql", dbLogin)
	defer db.DB.Close()
	db.DB.AutoMigrate(&db.Contents{}, &db.Users{})

	r := router.SetupRouter()
	r.Run("localhost:" + port)
}
