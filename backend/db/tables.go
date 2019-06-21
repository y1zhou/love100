package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	// Model gorm.Model definition
	Model struct {
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}

	// Contents ...
	Contents struct {
		CID    int64  `gorm:"PRIMARY_KEY;Type:serial;NOT NULL;"`
		Title  string `gorm:"Type:varchar(200);NOT NULL"`
		Status bool   `gorm:"Type:BOOLEAN;DEFAULT:0"`
		gorm.Model
	}

	//Users ...
	Users struct {
		UID      int64  `gorm:"PRIMARY_KEY;Type:serial;NOT NULL"`
		Username string `gorm:"Type:varchar(32);DEFAULT:NULL;UNIQUE;UNIQUE_INDEX:users_name"`
		Password string `gorm:"Type:varchar(64);NOT NULL"`
		Email    string `gorm:"Type:varchar(200);UNIQUE;DEFAULT:NULL"`
		gorm.Model
	}
)
