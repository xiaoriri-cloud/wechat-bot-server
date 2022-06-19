package models

import (
	"math/rand"
	"time"
)

type Wallpaper struct {
	ID       uint   `gorm:"primarykey"`
	ImgUrl   string `json:"url"`
	Source   string `json:"source"`
	Category int64  `json:"category"`
}

func PageWallpaper(page, step int) ([]Wallpaper, error) {

	var Wallpapers []Wallpaper

	var count int64 = 7533 //暂时写死
	//db.Find(&Wallpapers).Count(&count)

	rand.Seed(time.Now().UnixNano())
	id := rand.Int63n(count)

	db.Limit(20).Where("id >= ?", id).Find(&Wallpapers)

	return Wallpapers, nil
}
