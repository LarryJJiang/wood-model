package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"woods/pkg/e"
	"woods/pkg/gredis"
	"woods/pkg/logging"
	"woods/pkg/setting"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Uuid     string `json:"uuid"`
	Uid      int    `json:"uid"`
	Account  string `json:"account"`
	Identity int    `json:"identity"`
	State    int    `json:"state"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(username string, uid int, account string, identity int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(setting.AppSetting.JwtTokenExpireTime) * time.Second)
	var state int
	userTokenState := fmt.Sprintf(e.UserLoginTokenState, account)
	reply, err := gredis.Get(userTokenState)
	fmt.Printf("userTokenState = %v reply = %v, err = %v\n", userTokenState, reply, err)
	if err != nil {
		state = 1
	} else {
		reply, _ := strconv.Atoi(string(reply))
		state = reply + 1
	}
	if err := gredis.Set(userTokenState, state, setting.AppSetting.JwtTokenExpireTime); err != nil {
		logging.Error(err)
	}
	uuid := GetUUID()
	claims := Claims{
		username,
		uuid,
		uid,
		account,
		identity,
		state,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    setting.AppSetting.JwtIssuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			reply, err := gredis.Get(fmt.Sprintf(e.UserLoginTokenState, claims.Account))
			if err == nil {
				reply, _ := strconv.Atoi(string(reply))
				if reply != claims.State {
					return nil, jwt.NewValidationError("token not use", jwt.ValidationErrorMalformed)
				}
				return claims, nil
			}
			return nil, jwt.NewValidationError("token not use", jwt.ValidationErrorMalformed)
		}
	}

	return nil, err
}

// 刷新用户token
func RefreshToken(c *gin.Context) (string, error) {
	claims, err := ParseToken(c.GetHeader("Authorization"))
	if err != nil {
		return "", err
	}
	uuid := claims.Uuid
	uid := claims.Uid
	username := claims.Username
	c.Set(e.UserUuidKey, uuid)
	c.Set(e.UserIdKey, uid)
	c.Set(e.UserNameKey, username)
	c.Set(e.UserIdentityKey, claims.Identity)
	token, err := GenerateToken(claims.Username, claims.Uid, claims.Account, claims.Identity)
	return token, err
}
