package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secretkey")

func GenerateToken(username string) (map[string]string, error) {

	fmt.Println(time.Now().Unix())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(50 * time.Hour).Unix(),
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	})

	fmt.Println(time.Now().Add(5 * time.Minute).Unix())
	tokenString, _ := token.SignedString(secretKey)
	refreshTokenString, err := refreshToken.SignedString(secretKey)

	if err != nil {
		fmt.Println("error")
		return nil, err
	}

	return map[string]string{
		"token":        tokenString,
		"refreshToken": refreshTokenString,
	}, nil
}

func VerifyToken(ctx *gin.Context) error {
	token := ExtractToken(ctx)
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	fmt.Println(claims["exp"])

	if !ok || !parsedToken.Valid {
		return fmt.Errorf("validate: invalid token")
	}
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ValidateRefreshToken(ctx *gin.Context, token string) (map[string]string, error) {

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	username := claims["username"].(string)

	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	if err != nil {
		return nil, err
	}

	tokenstring, err := GenerateToken(username)

	if err != nil {
		return nil, err
	}
	return tokenstring, nil

}

func ExtractUserNameFromToken(ctx *gin.Context) (string, error) {
	token := ExtractToken(ctx)
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return "", fmt.Errorf("can not parse token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if ok {
		return claims["username"].(string), nil
	}
	return "", fmt.Errorf("validate: invalid token")
}
