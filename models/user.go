package models

import (
	"test/database"
	"time"
)

type User struct {
	Id                 int64 `gorm:"primaryKey"`
	FirstName          string
	LastName           string
	FullName           string
	Email              string
	Gender             string
	IpAddress          string
	Username           string
	Birthdate          time.Time
	Location           string
	Bio                string
	ProfilePic         string
	FollowersCount     string
	FollowingCount     string
	PostCount          string
	LastLogin          time.Time
	VerifiedAccount    string
	Interests          string
	Website            string
	PhoneNumber        string
	RelationshipStatus string
	Education          string
	Workplace          string
}

func init() {
	if err := database.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
}

func (u *User) Insert() error {
	tx := database.DB.Save(u)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *User) Update() error {
	tx := database.DB.Save(u)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *User) Delete() error {
	tx := database.DB.Delete(u)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
