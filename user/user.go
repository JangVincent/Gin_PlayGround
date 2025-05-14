package user

import (
	"fmt"
	"go-server/post"

	"gorm.io/gorm"
)

type User struct {
	ID    string `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email string `json:"email" gorm:"uniqueIndex"`
	Name  string `json:"name" gorm:"uniqueIndex"`
	Posts []post.Post `json:"posts" gorm:"foreignKey:UserID"`
}

type UserRepository interface {
	AutoMigrate(db *gorm.DB) error
	GetAll(db *gorm.DB) ([]User, error)
	Add(db *gorm.DB, u *User) error
	GetUser(db *gorm.DB, name string) (User, error)
	Delete(db *gorm.DB, id uint) error
	// (필요시 GetUser, Update, Delete 등 추가 가능)
}

type UserRepositoryImpl struct {}

func (repo *UserRepositoryImpl) AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

func (repo *UserRepositoryImpl) GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	fmt.Println(users)
	err := db.Find(&users).Error
	fmt.Println(users)
	return users, err
}

func (repo *UserRepositoryImpl) GetUser(db *gorm.DB, name string) (User, error) {
	var user User
	err := db.Find(&user, "name = ?", name).Error

	return user, err
}

func (repo *UserRepositoryImpl) Add(db *gorm.DB, u *User) error {
	return db.Create(u).Error
}

func (repo *UserRepositoryImpl) Delete(db *gorm.DB, id uint) error {
	return db.Delete(&User{}, id).Error
} 
