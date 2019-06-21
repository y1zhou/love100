package db

import (
	"github.com/jinzhu/gorm"
)

type (
	// Contents ...
	Contents struct {
		gorm.Model
		Title  string `gorm:"Type:varchar(200);NOT NULL"`
		Status bool   `gorm:"Type:BOOLEAN;DEFAULT:0"`
	}

	//Users ...
	Users struct {
		gorm.Model
		Username string `gorm:"Type:varchar(32);DEFAULT:NULL;UNIQUE;UNIQUE_INDEX:users_name"`
		Password string `gorm:"Type:varchar(64);NOT NULL"`
		Email    string `gorm:"Type:varchar(200);UNIQUE;DEFAULT:NULL"`
	}
)
