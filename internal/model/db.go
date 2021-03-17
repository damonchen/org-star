package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/damonchen/org-star/internal/config"
)

var (
	db *gorm.DB
)

// DBInit db init
func DBInit(cfg *config.DB) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.Charset)

	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//panic("failed to connect database")
		return nil, err
	}

	db = _db
	return _db, nil
}
