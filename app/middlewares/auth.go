package middlewares

import (
	"disspace/controllers"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtClaims struct {
	Username string `json:"username" bson:"username"`
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

func (jwtConf *ConfigJwt) GenerateToken(username string) (string, error) {
	claims := JwtClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresAt))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(jwtConf.Secret))

	return token, err
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	tokenString := strings.Replace(tokenStr, "bearer ", "", -1)
	// fmt.Println(tokenString)
	hmacSecretString := "UhYiPkGrOuP10fGd"
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}

func GetUsername(c echo.Context) string {
	user := c.Get("username").(*jwt.Token)
	claims := user.Claims.(*JwtClaims)
	return string(claims.Username)
}
