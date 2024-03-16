package v1

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strings"
)

const (
	statusOK    = "OK"
	statusError = "Error"
)

var errBadRequest = errors.New("bad request")

type Auth interface {
	RegisterNewUser(
		ctx context.Context,
		email string,
		login string,
		password string,
	) (userID int64, err error)

	Login(
		ctx context.Context,
		email string,
		password string,
	) (token string, err error)
}

type authRoutes struct {
	auth Auth
	log  *slog.Logger
}

type signUpRequest struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type signUpResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func newAuthRoutes(g *echo.Group) {
	r := authRoutes{}
	g.POST("/sign-in", r.signIn)
	g.POST("/sign-up", r.signUp)
}

func (r *authRoutes) signIn(c echo.Context) error {
	return c.JSON(http.StatusOK, "Test")
}

// TODO: Сделать нормальную валидацию входных данных
func (r *authRoutes) signUp(c echo.Context) error {
	r.log.Info("registration attempt")
	var req = signUpRequest{}

	if err := c.Bind(&req); err != nil {
		r.log.Error("registration attempt failed", slog.String("error", err.Error()))
		return c.JSON(http.StatusBadRequest, signUpResponse{
			Status:  statusError,
			Message: "failed to register",
		})
	}

	errMsg, err := validateSignUp(req)
	if err != nil {
		r.log.Info("registration attempt failed", slog.String("error", err.Error()))
		return c.JSON(http.StatusBadRequest, signUpResponse{
			Status:  statusError,
			Message: errMsg,
		})
	}

	//TODO: обработать возмодные ошибки (пользователь уже существует)
	userID, err := r.auth.RegisterNewUser(c.Request().Context(), req.Email, req.Login, req.Password)
	if err != nil {
		r.log.Error("registration attempt failed", slog.String("error", err.Error()))
		return c.JSON(http.StatusInternalServerError, signUpResponse{
			Status:  statusError,
			Message: "internal error",
		})
	}

	return c.JSON(http.StatusCreated, signUpResponse{
		Status:  statusOK,
		Message: fmt.Sprintf("User created: ID = %d", userID),
	})
}

func validateSignUp(req signUpRequest) (string, error) {
	var errMsgs []string
	if req.Email == "" {
		errMsgs = append(errMsgs, "email is required field")
	}

	if req.Login == "" {
		errMsgs = append(errMsgs, "login is required field")
	}

	if req.Password == "" {
		errMsgs = append(errMsgs, "password is required field")
	}

	if len(errMsgs) == 0 {
		return "", nil
	}

	return strings.Join(errMsgs, ","), errBadRequest
}
