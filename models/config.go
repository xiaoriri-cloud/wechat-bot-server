package models

import "gorm.io/gorm"

type Config struct {
	Name   string `json:"name"`
	Config string `json:"config"`
}

func ExistConfigByName(name string) (bool, error) {
	var config Config
	err := db.Select("name").Where("name = ? ", name).First(&config).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if config.Name != "" {
		return true, nil
	}

	return false, nil
}

func GetConfig(key string) (*Config, error) {
	var config Config
	err := db.Where("name = ?", key).First(&config).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &config, nil
}

func SetConfig(config Config) error {
	isExist, _ := ExistConfigByName(config.Name)
	if isExist {
		if err := db.Model(&config).Where("name = ?", config.Name).Update("config", config.Config).Error; err != nil {
			return err
		}
	}else {
		if err := db.Create(config).Error; err != nil {
			return err
		}
	}

	return nil
}
