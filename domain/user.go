package domain

import "context"

type User struct {
	Login    string
	Password string
	Balance  string
	Contract string
	FullName string

	SessionId string `json:"session_id"`
}

type UserUsecase interface {
	SignIn(ctx context.Context, user User) (string, error)
	GetBalance(ctx context.Context, SessionId string) (float64, error)
	GetUserInfo(ctx context.Context, SessionId string) (User, error)
}
type UserRepository interface {
	SignIn(ctx context.Context, user User) (string, error)
	GetBalance(ctx context.Context, SessionId string) (float64, error)
	GetUserInfo(ctx context.Context, SessionId string) (User, error)
}
