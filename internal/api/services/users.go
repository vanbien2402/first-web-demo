package services

import (
	"context"
	"errors"
	"github.com/vanbien2402/first-web-demo/internal/api/domain"
	"github.com/vanbien2402/first-web-demo/internal/api/models"
	"github.com/vanbien2402/first-web-demo/internal/pkg/identifier"
	"github.com/vanbien2402/first-web-demo/internal/pkg/rds"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type userService struct {
	userRepo domain.UserRepository
}

func (s *userService) Create(ctx context.Context, req domain.CreateParams) (user *models.User, err error) {
	user = &models.User{
		Model:    rds.NewModel(identifier.Generate()),
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
	}
	var password []byte
	password, err = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return
	}
	user.Password = string(password)
	if s.userRepo.CheckExist(ctx, req.UserName) {
		err = errors.New("user name has existed")
		return
	}
	if err = s.userRepo.Create(ctx, user); err != nil {
		return
	}
	return
}

func (s *userService) Validate(ctx context.Context, userName, password string) (res *domain.UserInfo, err error) {
	user, err := s.userRepo.FindByUserName(ctx, userName)
	if err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return
	}
	res = &domain.UserInfo{
		Model:    user.Model,
		UserName: user.UserName,
		Email:    user.Email,
	}
	return
}

func (s *userService) Get(ctx context.Context, userName string) (res *domain.UserInfo, err error) {
	user, err := s.userRepo.FindByUserName(ctx, userName)
	if err != nil {
		return
	}
	res = &domain.UserInfo{
		Model:    user.Model,
		UserName: user.UserName,
		Email:    user.Email,
	}
	return
}

func (s *userService) Update(ctx context.Context, req domain.UpdateParams) (user *models.User, err error) {
	user, err = s.userRepo.FindByID(ctx, req.ID)
	if err != nil {
		log.Println("FindByUserName error")
		return
	}
	if user.Version != req.Version {
		err = errors.New("version not match")
		log.Println(err)
		return
	}
	if len(req.UserName) > 0 && req.UserName != user.UserName {
		if s.userRepo.CheckExist(ctx, req.UserName) {
			err = errors.New("user name has existed")
			return
		}
		user.UserName = req.UserName
	}
	if len(req.Password) > 0 {
		var pw []byte
		pw, err = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
		if err != nil {
			log.Println("encrypt password error")
			return
		}
		user.Password = string(pw)
	}
	if len(req.Email) > 0 && req.Email != user.Email {
		user.Email = req.Email
	}
	if err = s.userRepo.Update(ctx, user); err != nil {
		return
	}
	return
}

func (s *userService) Delete(ctx context.Context, id string) (err error) {
	if err = s.userRepo.Delete(ctx, id); err != nil {
		return
	}
	return
}

//NewUserService init user service
func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{
		userRepo: userRepo,
	}
}
