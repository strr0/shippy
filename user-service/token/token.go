package token

import (
	"github.com/golang-jwt/jwt"
	"time"
	pb "user-service/proto/user"
)

var key = []byte("mySuperSecretKeyLol")

func Decode(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func Encode(user *pb.User) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.Id,
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour).Unix(),
	}).SignedString(key)
}