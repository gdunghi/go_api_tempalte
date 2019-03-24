package route

import (
	"context"
	"github.com/labstack/echo"
	md "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go_api_tempalte/api/login"
	"net/http"
)

func InitRouter(ctx context.Context, mc *mongo.Client) *echo.Echo {
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

	a := login.NewHandler(ctx, mc, login.NewAuthentication(ctx, mc))

	e.POST("/login", a.Auth)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	return e
}
