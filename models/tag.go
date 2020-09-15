package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	Model

	Name     string    `json:"name"`
	State    int       `json:"state"`
	Articles []Article `json:"articles"`
}

// AddTag Add a Tag
func AddTag(name string, state int) (Tag, error) {
	tag := Tag{
		Name:  name,
		State: state,
	}
	err := db.Create(&tag).Error
	return tag, err
}

// GetTags gets a list of tags based on paging and constraints
func GetTags(pageNum int, pageSize int) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)
	if pageSize > 0 && pageNum >= 0 {
		err = db.Offset(pageNum).Limit(pageSize).Find(&tags).Error
	} else {
		err = db.Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, err
}

func GetTag(id string) (Tag, error) {
	var tag Tag
	err := db.Preload("Articles").First(&tag, id).Error

	return tag, err
}

func EditTag(id string, data map[string]interface{}) (Tag, error) {
	var tag Tag
	err := db.Where("id = ?", id).First(&tag).Updates(data).Error

	return tag, err
}

func DeleteTag(id string) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}

	return nil
}
