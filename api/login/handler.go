package login

import "C"
import (
	"context"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authenticationer interface {
	CheckLogin(auth Auth) (bool, error)
}

//Handler ...
type Handler struct {
	client *mongo.Client
	ctx    context.Context
	auth   Authenticationer
}

func NewHandler(ctx context.Context, client *mongo.Client, auth Authenticationer) *Handler {
	return &Handler{
		client: client,
		ctx:    ctx,
		auth:   auth,
	}
}

func (h *Handler) Auth(c echo.Context) error {
	l := new(Login)
	if err := c.Bind(l); err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	r, err := h.auth.CheckLogin(Auth{l.Username, l.Password})

	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	if !r {
		c.String(http.StatusUnauthorized, "")
	}

	return c.String(http.StatusOK, "")
}
