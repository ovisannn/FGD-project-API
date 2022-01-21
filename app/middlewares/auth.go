package middlewares

import (
	"disspace/controllers"
	"net/http"
	"time"

	"errors"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtClaims struct {
	UserID string `json:"UserID"`
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

func Auth(c echo.Context) error {
	tokenString := c.Request().Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, errors.New("Unexpected signing method: HS256")
		}

		return []byte("UhYiPkGrOuP10fGd"), nil
	})

	if token != nil && err == nil {
		return controllers.NewSuccessResponse(c, "token authorized")
	} else {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errors.New("token unauthorized"))
	}
}
