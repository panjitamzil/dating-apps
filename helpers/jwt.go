package helpers

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secretKey = "dating-apps"

func GenerateToken(email, subscription string) string {
	claims := jwt.MapClaims{
		"email":        email,
		"subscription": subscription,
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		return EMPTY
	}

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")
	if !bearer {
		return nil, errors.New(ERR_TOKEN_SIGNIN)
	}

	tokenString := headerToken[len("Bearer "):]
	token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(ERR_TOKEN_SIGNIN)
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errors.New(ERR_TOKEN_SIGNIN)
	}

	// check exp
	expirationTime, _ := token.Claims.(jwt.MapClaims)["exp"].(float64)

	// Convert the expiration time to time.Time
	expiration := time.Unix(int64(expirationTime), 0)
	currentTime := time.Now()
	if currentTime.After(expiration) {
		return nil, errors.New(ERR_TOKEN_EXPIRED)
	}

	return token.Claims.(jwt.MapClaims), nil
}
