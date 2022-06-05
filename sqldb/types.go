package sqldb

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Insight struct {
	gorm.Model
	Text        string `gorm:"column:text;type:varchar(200)"`
	Description string `gorm:"column:description;type:varchar(200)"`
}

type Audience struct {
	gorm.Model
	AgeMax            int    `gorm:"column:age_max"`
	AgeMin            int    `gorm:"column:age_min"`
	Gender            string `gorm:"column:gender"`
	Country           string `gorm:"column:country"`
	HoursSpent        int    `gorm:"column:hours_spent"`
	NumberOfPurchases int    `gorm:"column:purchases"`
	Description       string `gorm:"column:description"`
}

type Chart struct {
	gorm.Model
	Title       string         `gorm:"column:title"`
	XTitle      string         `gorm:"column:x_title"`
	YTitle      string         `gorm:"column:y_title"`
	Description string         `gorm:"column:title"`
	Data        datatypes.JSON `gorm:"column:data"`
}

type FavouriteAsset struct {
	gorm.Model
	AssetType string `gorm:"column:asset_type"`
	AssetID   uint   `gorm:"column:asset_id"`
	UserID    uint   `gorm:"column:user_id"`
}

type User struct {
	gorm.Model
	Username   string `gorm:"column:username"`
	PassHashed string `gorm:"column:pass_hashed"`
	IsAdmin    bool   `gorm:"is_admin"`
}
