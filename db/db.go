package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"not null"`
	LtmidV2       string `gorm:"not null"`
	LtuidV2       string `gorm:"not null"`
	LtokenV2      string `gorm:"not null"`
	Genshin       bool
	Starrail      bool
	ZZZ           bool
	LastCheckinAt *time.Time
}

type DB struct {
	db *gorm.DB
}

func InitDB(dbPath string) (*DB, error) {
	gormdb, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := gormdb.AutoMigrate(&User{}); err != nil {
		return nil, err
	}
	return &DB{db: gormdb}, nil
}

func (d *DB) GetAllUsers() ([]User, error) {
	var users []User
	result := d.db.Find(&users)
	return users, result.Error
}

func (d *DB) InsertUser(user *User) error {
	return d.db.Create(user).Error
}

// UpdateLastCheckinAt 將指定 user 的 last_checkin_at 欄位更新為 now
func (d *DB) UpdateLastCheckinAt(userID uint, now time.Time) error {
	return d.db.Model(&User{}).Where("id = ?", userID).Update("last_checkin_at", now).Error
}
