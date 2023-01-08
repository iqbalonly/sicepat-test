package storage

import (
	"context"
	"sicepat/internal/dto"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	TX = "TX"
)

type Storage struct {
	conn *gorm.DB
}

type Config interface {
	DatabaseServerName() string
}

func NewStorage(cfg Config) (StorageDAO, error) {
	db, err := gorm.Open(mysql.Open(cfg.DatabaseServerName()), &gorm.Config{})
	if err != nil {
		logrus.WithField("error", err.Error()).Fatal("failed to initiate db")
		return nil, err
	}

	return &Storage{
		conn: db,
	}, nil
}

type StorageDAO interface {
	GetUsers(ctx context.Context) (*[]dto.User, error)
	DeleteUser(ctx context.Context, userID int) error
	CreateOrUpdateUser(ctx context.Context, req *dto.UserRequest) error
}

func (s *Storage) pool(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(TX).(*gorm.DB)
	if !ok {
		db = s.conn
	}

	return db
}
