package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL driver from gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB Global database connection
var DB *gorm.DB
