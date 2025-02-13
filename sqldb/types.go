package sqldb

import (
	"errors"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var (
	ErrThisAssetTypeDoesNotExist = errors.New("this asset type does not exists")
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

type AudienceWithFavour struct {
	Audience
	IsFavour bool `gorm:"column:is_favourite"`
}

type ChartWithFavour struct {
	Chart
	IsFavour bool `gorm:"column:is_favourite"`
}

type InsightWithFavour struct {
	Insight
	IsFavour bool `gorm:"column:is_favourite"`
}

type Chart struct {
	gorm.Model
	Title       string         `gorm:"column:title"`
	XTitle      string         `gorm:"column:x_title"`
	YTitle      string         `gorm:"column:y_title"`
	Description string         `gorm:"column:description"`
	Data        datatypes.JSON `gorm:"column:data"`
}

type FavouriteInsight struct {
	gorm.Model
	InsightID uint `gorm:"column:insight_id"`
	UserID    uint `gorm:"column:user_id"`
}

type FavouriteChart struct {
	gorm.Model
	ChartID uint `gorm:"column:chart_id"`
	UserID  uint `gorm:"column:user_id"`
}

type FavouriteAudience struct {
	gorm.Model
	AudienceID uint `gorm:"column:audience_id"`
	UserID     uint `gorm:"column:user_id"`
}

type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(200)"`
	Password string `gorm:"column:password;type:varchar(200)"`
	IsAdmin  bool   `gorm:"is_admin"`
}
