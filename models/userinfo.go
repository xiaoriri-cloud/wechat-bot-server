package models

import "gorm.io/gorm"

type UserInfo struct {
	ID        uint   `gorm:"primarykey"`
	Wid       string `json:"wid"`
	Account   string `json:"account"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	Nation    string `json:"nation"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Mobile    string `json:"mobile"`
	Gender    string `json:"gender"`
}

func ExistUserInfoByWid(wid string) (bool, error) {
	var userinfo UserInfo
	err := db.Select("id").Where("wid = ? ", wid).First(&userinfo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if userinfo.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddUserInfo(userinfo UserInfo) error {

	if err := db.Create(&userinfo).Error; err != nil {
		return err
	}

	return nil
}

func GetUserInfo(wid string) (*UserInfo, error) {
	var userinfo UserInfo
	err := db.Where("wid = ?", wid).First(&userinfo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &userinfo, nil
}
