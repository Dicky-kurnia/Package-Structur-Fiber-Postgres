package helper

import (
	"fmt"
	"go-fiber-postgres/config"
	"go-fiber-postgres/model"

	"github.com/goccy/go-json"

	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func InsertRedis(payload model.SetDataRedis) {
	var redis = config.RedisConnection()
	jsonResult, err := json.Marshal(payload.Data)
	IsShouldPanic(err)
	err = redis.Set(payload.Key, jsonResult, payload.Exp).Err()
	IsShouldPanic(err)
}

func GetRedis[T any](key string) (cek bool, raw T) {
	var redis = config.RedisConnection()

	value, _ := redis.Get(key).Result()
	if value == "" {
		return false, raw
	}

	_ = json.Unmarshal([]byte(value), &raw)

	return true, raw
}

func DelRedis(key string) {
	var redis = config.RedisConnection()
	redis.Del(key)
}

func DelRedisByPattern(pattern string) {
	var redis = config.RedisConnection()
	redis.Eval("for i, name in ipairs(redis.call('KEYS', '"+pattern+"')) do redis.call('expire', name, 0); end", []string{"*"})
}

func CreateToken(request model.JwtPayload) *model.TokenDetails {
	accessExpired, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_MINUTE"))

	td := &model.TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * time.Duration(accessExpired)).Unix()

	keyAccess, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("JWT_ACCESS_PRIVATE_KEY")))
	IsShouldPanic(err)

	now := time.Now().UTC()

	atClaims := jwt.MapClaims{}
	atClaims["sales_id"] = request.SalesId
	atClaims["username"] = request.Username
	atClaims["role"] = request.Role
	atClaims["exp"] = td.AtExpires
	atClaims["iat"] = now.Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	at.Header["topindopay"] = "jwt"
	td.AccessToken, err = at.SignedString(keyAccess)

	if err != nil {
		IsShouldPanic(errors.New(model.AUTHENTICATION_FAILURE_ERR_TYPE))
	}

	return td

}

func CreateAuth(request model.JwtPayload, td *model.TokenDetails) {
	at := time.Unix(td.AtExpires, 0)
	now := time.Now()

	InsertRedis(model.SetDataRedis{
		Key:  td.AccessToken,
		Data: request,
		Exp:  at.Sub(now),
	})
}

func TimeDiffInMinutes(startTime, endTime time.Time) (minutesDiff float64) {
	t1 := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour(), startTime.Minute(), 0, 0, time.UTC)
	t2 := time.Date(endTime.Year(), endTime.Month(), endTime.Day(), endTime.Hour(), endTime.Minute(), 0, 0, time.UTC)
	fmt.Printf("The minutes difference is: %f", t2.Sub(t1).Minutes())
	return t2.Sub(t1).Minutes()
}
