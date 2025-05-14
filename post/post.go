package post

import (
	"gorm.io/gorm"
)

type Post struct {
	ID   string   `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID string   `json:"user_id"`
}

type PostRepository struct {}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Post{})
}

func GetAll(db *gorm.DB) ([]Post, error) {
	var posts []Post
	err := db.Find(&posts).Error
	return posts, err
}

func Add(db *gorm.DB, p *Post) error {
	return db.Create(p).Error
} 
