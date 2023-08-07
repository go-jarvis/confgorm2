package pgsql

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	Host     string `env:""`
	Port     uint16 `env:""`
	Username string `env:""`
	Password string `env:""`
	DBName   string `env:""`
}

func (s *Server) DB() *gorm.DB {
	dsn := s.dsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// 链接测试
	sqldb, err := db.DB()
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err = sqldb.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	return db
}

func (s *Server) dsn() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		s.Username, s.Password,
		s.Host, s.Port,
		s.DBName,
	)
}
