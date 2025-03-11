package database

import (
	"fmt"
	"gomod/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(cfg *config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &DB{db}, nil
}
