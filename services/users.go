package services

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/major28/monitoring/utils/logger"
)

type contextKey string

var ContextKeyUsername = contextKey("username")

type UsersIService interface {
	CheckAuth(ctx context.Context, username string, password string) (bool, error)
	AddRequestInHistory(ctx context.Context, request string) error
	GetRequestsFromHistory(ctx context.Context) ([]usersHistoryRequest, error)
}

type usersService struct {
	db     *gorm.DB
	logger logger.ILogger
}

type user struct {
	gorm.Model
	Username string
	Password string
}

type usersHistoryRequest struct {
	gorm.Model
	FkUser  uint
	Request string
}

func NewUsersService(db *gorm.DB, logger logger.ILogger) UsersIService {
	return &usersService{
		db:     db,
		logger: logger,
	}
}

func (s *usersService) CheckAuth(ctx context.Context, name string, password string) (bool, error) {

	resp := &user{}
	err := s.db.Where("username=?", name).Where("password=?", password).First(resp).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (s *usersService) AddRequestInHistory(ctx context.Context, request string) error {

	usernameCtx, ok := ctx.Value(ContextKeyUsername).([]string)
	if !ok {
		return fmt.Errorf("can't type assert for username")
	}
	username := usernameCtx[0]

	user := &user{}
	err := s.db.Where("username=?", username).First(user).Error
	if err != nil {
		return fmt.Errorf("can't get user (%s)", err.Error())
	}

	err = s.db.Save(&usersHistoryRequest{
		FkUser:  user.ID,
		Request: request,
	}).Error
	if err != nil {
		return fmt.Errorf("can't save user request in history (%s)", err.Error())
	}

	return nil
}

func (s *usersService) GetRequestsFromHistory(ctx context.Context) ([]usersHistoryRequest, error) {
	usernameCtx, ok := ctx.Value(ContextKeyUsername).([]string)
	if !ok {
		return nil, fmt.Errorf("can't type assert for username")
	}
	username := usernameCtx[0]

	user := &user{}
	err := s.db.Where("username=?", username).First(user).Error
	if err != nil {
		return nil, fmt.Errorf("can't get user (%s)", err.Error())
	}

	var usersHistoryRequests []usersHistoryRequest
	err = s.db.Where("fk_user=?", user.ID).Order("id DESC").Find(&usersHistoryRequests).Error
	if err != nil {
		return nil, fmt.Errorf("can't get user (%s)", err.Error())
	}

	return usersHistoryRequests, nil
}
