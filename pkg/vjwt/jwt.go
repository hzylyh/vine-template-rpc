/**
 * @Description: jwt生成token
 * @Author: Ethan Howard
 * @Github: https://github.com/hzylyh
 * @Date: 2023/10/26 18:35
 * @LastEditors: Ethan Howard
 * @LastEditTime: 2023/10/26 18:35
 */

package vjwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secureKey = "2wer43@WER$#"

type Claims struct {
	UserId   uint   `json:"userId,omitempty"`
	Username string `json:"username,omitempty"`
	Nonce    string `json:"nonce,omitempty"`
	Tag      string `json:"tag,omitempty"`
	Scope    string `json:"scope,omitempty"`
	jwt.RegisteredClaims
}

func GenerateJwtToken(id uint, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(1) * time.Hour)
	//expireTime := nowTime.Add(time.Duration(10) * time.Second)

	//claim := &jwt.MapClaims{
	//	"userId":    userId,
	//	"ExpiresAt": jwt.NewNumericDate(expireTime),
	//}

	claim := Claims{
		UserId: id,
		//Nonce: nonce,
		// FIXME: A workaround for custom claim by reusing `tag` in user info
		//Tag:   user.Tag,
		//Scope: scope,
		RegisteredClaims: jwt.RegisteredClaims{
			//Issuer:    originBackend,
			Subject: username,
			//Audience:  []string{application.ClientId},
			ExpiresAt: jwt.NewNumericDate(expireTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			//ID:        jti,
		},
	}

	var token *jwt.Token
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	key := []byte(secureKey)
	tokenString, err := token.SignedString(key)

	return tokenString, err
}

func ParseJwtToken(tokenStr string) (claims *Claims, err error) {
	//var hmacSampleSecret []byte
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secureKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("token is invalid")
	}

	return claims, err
}
