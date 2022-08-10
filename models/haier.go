package models

import "time"

type HaierInfo struct {
	ID        uint   `gorm:"primarykey"`
	Phone     string `json:"phone"`
	ExpiresIn int64 `json:"expiresIn"`
	AccountToken  string  `json:"accountToken"`
	RawData  string  `json:"rawData"`
	CreaetTime int64 `json:"creaetTime"`
}


func AddHaierInfo(haierInfo HaierInfo) error {
	haierInfo.CreaetTime = time.Now().Unix()
	if err := db.Create(&haierInfo).Error; err != nil {
		return err
	}
	return nil
}
