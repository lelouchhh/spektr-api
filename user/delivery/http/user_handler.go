package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
	"net/http"
	"spektr-account-api/domain"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type UserHandler struct {
	UUsecase domain.UserUsecase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewUserHandler(e *gin.Engine, us domain.UserUsecase) {
	handler := &UserHandler{
		UUsecase: us,
	}
	//e.GET("/deliver", handler.GetByID)
	g := e.Group("/auth")

	g.POST("/sign_in", handler.SignIn)
	g.GET("/balance", handler.GetBalance)
	g.GET("/user_info", handler.GetUserInfo)
}

func (a *UserHandler) SignIn(c *gin.Context) {
	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrBadParamInput.Error())
		return
	}
	if user.Password == "" || user.Login == "" {
		c.JSON(http.StatusBadRequest, ResponseError{Message: domain.ErrBadParamInput.Error()})
		return
	}
	ctx := c.Request.Context()
	session, err := a.UUsecase.SignIn(ctx, user)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
	return
}
func (a *UserHandler) GetBalance(c *gin.Context) {
	return
}

func (a *UserHandler) GetUserInfo(c *gin.Context) {
	return
}

func isRequestValid(m *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
