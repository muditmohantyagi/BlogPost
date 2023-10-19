package models

import (
	"errors"
	"time"

	"blog.com/config"
	"gorm.io/gorm"
)

type Post struct {
	ID      uint
	Title   string `gorm:"type:varchar(250)"`
	Post    string
	Created time.Time `gorm:"autoCreateTime"`
	Updated time.Time `gorm:"autoUpdateTime"`
}

var db = config.GoConnect()

func CreatePost(post_data Post) (bool, error) {
	if result := db.Create(&post_data); result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func ListPost() ([]Post, error) {
	var post []Post
	if result := db.Find(&post); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return nil, nil
		}
		return nil, result.Error
	}
	return post, nil
}
func UpdatePost(post Post) (bool, error) {

	if result := db.Model(&Post{}).Where("id=?", post.ID).Updates(post); result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
func DeletePost(id uint) (bool, error) {
	if result := db.Delete(&Post{}, id); result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func ViewPost(id uint) (*Post, error) {
	var post = new(Post)
	if result := db.Where("id=?", id).Take(&post); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			return nil, nil
		}
		return nil, result.Error
	}
	return post, nil
}
