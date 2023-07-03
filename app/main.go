package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"log"
	_UserHttp "spektr-account-api/user/delivery/http"
	_UserRepo "spektr-account-api/user/repository/rest"
	_UserUcase "spektr-account-api/user/usecase"
	"time"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {

	g := gin.Default()
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	baseURL := viper.GetString("rest.baseURL")

	userRepo := _UserRepo.NewDeliverRepository(baseURL)
	UserUcase := _UserUcase.NewDeliverUsecase(userRepo, timeoutContext)
	_UserHttp.NewUserHandler(g, UserUcase)

	log.Fatal(g.Run(viper.GetString("server.address"))) //nolint
}
