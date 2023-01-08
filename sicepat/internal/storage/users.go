package storage

import (
	"context"
	"sicepat/internal/constant"
	"sicepat/internal/dto"
	"sicepat/internal/ierr"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userModel struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	Email       string
	DateOfBirth *time.Time
	CreatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (userModel) TableName() string {
	return "users"
}

func (s *Storage) GetUsers(ctx context.Context) (*[]dto.User, error) {
	conn := s.pool(ctx)

	response := make([]dto.User, 0)
	data := make([]userModel, 0)

	err := conn.WithContext(ctx).Find(&data).Error
	if err != nil {
		logrus.WithField("error", err.Error()).Error("failed to get users")

		return nil, err
	}

	for _, d := range data {
		dob := ""
		if d.DateOfBirth != nil {
			dob = d.DateOfBirth.Format(constant.DateReturnPayload)
		}

		response = append(response, dto.User{
			ID:          int64(d.ID),
			Name:        d.Name,
			Email:       d.Email,
			DateOfBirth: dob,
			CreatedAt:   d.CreatedAt.Format(constant.DateReturnPayload),
		})
	}

	return &response, nil
}

func (s *Storage) DeleteUser(ctx context.Context, userID int) error {
	conn := s.pool(ctx)

	err := conn.WithContext(ctx).Where("id = ?", userID).Delete(&userModel{}).Error
	if err != nil {
		logrus.WithField("error", err.Error()).Error("failed to delete user")

		return err
	}

	return nil
}

func (s *Storage) CreateOrUpdateUser(ctx context.Context, req *dto.UserRequest) error {
	conn := s.pool(ctx)

	var dob *time.Time

	if req.DateOfBirth != "" {
		date, _ := time.Parse(dto.RequestDateLayout, req.DateOfBirth)
		dob = &date
	}

	user := userModel{
		ID:          req.ID,
		Name:        req.Name,
		Email:       req.Email,
		DateOfBirth: dob,
	}

	res := conn.WithContext(ctx).Omit("created_at").Save(&user)

	err := res.Error
	if err != nil {
		logrus.WithField("error", err.Error()).Error("failed to create or update user")
	}

	if res.RowsAffected == 0 {
		return ierr.NoRowsAffected
	}

	return nil
}
