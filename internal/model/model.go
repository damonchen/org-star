package model

import (
	"gorm.io/gorm"
)

// User must github user or gitee user or gitlab user
type User struct {
	gorm.Model
}

// GitHubUser git hub user
type GitHubUser struct {
	gorm.Model
	UserID uint   `gorm:"user_id"`
	Email  string `gorm:"email"`
}

// Favorite favorite
type Favorite struct {
	gorm.Model
	UserID  uint   `gorm:"user_id"`
	Alias   string `gorm:"alias"` // 别名，允许为空
	Url     string `gorm:"url"`
	GroupId string `gorm:"group_id"` // 组id
}

type Group struct {
	gorm.Model
	UserID uint   `gorm:"user_id"`
	Name   string `gorm:"name"` // 组名，每个用户构建的时候，有一个default的组名
}
