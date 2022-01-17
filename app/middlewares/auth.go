package middlewares

import (
	"disspace/controllers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtClaims struct {
	UserID string
	jwt.StandardClaims
}

type ConfigJwt struct {
	Secret    string
	ExpiresAt int64
}

func (config *ConfigJwt) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtClaims{},
		SigningKey: []byte(config.Secret),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return controllers.NewErrorResponse(c, http.StatusForbidden, e)
		}),
	}
}

func (jwtConf *ConfigJwt) GenerateToken(UserID string) (string, error) {
	claims := JwtClaims{
		UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresAt))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(jwtConf.Secret))

	return token, err
}
