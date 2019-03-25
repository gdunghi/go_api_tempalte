package route

import (
	"context"
	"github.com/labstack/echo"
	md "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"go_api_tempalte/api/login"
	"net/http"
)

type DBer interface {
	GetAll(collection string, selector interface{}, sort string, v interface{}) error
	GetOne(ctx context.Context, collection string, selector interface{}, v interface{}) error
	Insert(collection string, v interface{}) error
	Upsert(collection string, selector interface{}, v interface{}) error
}

func InitRouter(db DBer) *echo.Echo {
	e := echo.New()
	e.Pre(md.RemoveTrailingSlash())
	e.Logger.SetLevel(log.INFO)
	e.Use(
		md.Recover(),
		md.Secure(),
		md.Logger(),
		md.BodyLimit("2M"),
		md.CORSWithConfig(md.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{
				echo.HeaderOrigin,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
				echo.HeaderContentType,
				echo.HeaderAuthorization,
			},
			AllowMethods: []string{
				echo.GET,
				echo.POST,
				echo.PUT,
				echo.PATCH,
				echo.OPTIONS,
			},
			MaxAge: 3600,
		}),
	)

	a := login.NewHandler(db, login.NewAuthentication(db))

	e.POST("/login", a.Auth)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	return e
}
