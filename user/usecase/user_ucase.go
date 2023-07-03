package usecase

import (
	"context"
	"spektr-account-api/domain"
	"time"
)

type UserUsecase struct {
	UserRepo       domain.UserRepository
	contextTimeout time.Duration
}

func (u UserUsecase) SignIn(ctx context.Context, user domain.User) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()
	session, err := u.UserRepo.SignIn(ctx, user)
	if err != nil {
		return "", err
	}
	return session, nil
}

func (u UserUsecase) GetBalance(ctx context.Context, SessionId string) (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserUsecase) GetUserInfo(ctx context.Context, SessionId string) (domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewDeliverUsecase(u domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &UserUsecase{
		UserRepo:       u,
		contextTimeout: timeout,
	}
}
