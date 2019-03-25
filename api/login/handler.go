package login

import "C"
import (
	"context"
	"github.com/labstack/echo"
	"net/http"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DBer interface {
	GetAll(collection string, selector interface{}, sort string, v interface{}) error
	GetOne(ctx context.Context, collection string, selector interface{}, v interface{}) error
	Insert(collection string, v interface{}) error
	Upsert(collection string, selector interface{}, v interface{}) error
}

type Authenticationer interface {
	CheckLogin(ctx context.Context, auth Auth) (bool, error)
}

//Handler ...
type Handler struct {
	db   DBer
	auth Authenticationer
}

func NewHandler(db DBer, auth Authenticationer) *Handler {
	return &Handler{
		db:   db,
		auth: auth,
	}
}

func (h *Handler) Auth(c echo.Context) error {
	l := new(Login)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	if err := c.Bind(l); err != nil {
		return c.String(http.StatusBadRequest, "")
	}

	r, err := h.auth.CheckLogin(ctx, Auth{l.Username, l.Password})

	if err != nil {
		return c.String(http.StatusBadRequest, "")
	}
	if !r {
		c.String(http.StatusUnauthorized, "")
	}

	return c.String(http.StatusOK, "")
}
