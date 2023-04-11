package middleware

import (
	"errors"
	"github.com/goccy/go-json"
	"go-fiber-postgres/helper"
	"go-fiber-postgres/model"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DecodeToken(tokenString string) (decodedResult model.JwtPayload, err error) {
	jwtPublicKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(os.Getenv("JWT_ACCESS_PUBLIC_KEY")))

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtPublicKey, nil
	})
	if err != nil {
		return decodedResult, err
	}
	if !token.Valid {
		return decodedResult, errors.New("invalid token")
	}

	jsonBody, err := json.Marshal(claims)
	if err != nil {
		return decodedResult, err
	}

	var obj model.JwtPayload
	if err := json.Unmarshal(jsonBody, &obj); err != nil {
		return decodedResult, err
	}

	return obj, nil
}

func CheckToken(c *fiber.Ctx) error {
	// get token (Bearer tokentokentoken)
	tokenSlice := strings.Split(c.Get("Authorization"), "Bearer ")

	var tokenString string
	if len(tokenSlice) == 2 {
		tokenString = tokenSlice[1]
	}

	// extract data from token
	decodedRes, err := DecodeToken(tokenString)
	if err != nil {
		response := model.Response{
			Code:   401,
			Status: "Unauthorized",
			Error: map[string]interface{}{
				"general": "UNAUTHORIZED",
			},
		}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}

	//check in redis
	check, _ := helper.GetRedis[model.JwtPayload](tokenString)

	if check == false {
		response := model.Response{
			Code:   401,
			Status: "Unauthorized",
			Error: map[string]interface{}{
				"general": "UNAUTHORIZED",
			},
		}
		return c.Status(http.StatusUnauthorized).JSON(response)
	}
	// set to global var
	c.Locals("currentUserID", decodedRes.SalesId)
	c.Locals("currentUserName", decodedRes.Username)
	c.Locals("currentToken", tokenString)
	return c.Next()
}
