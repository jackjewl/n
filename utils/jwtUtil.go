// Package utils 工具包
package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Rule 角色，角色分为用户和管理员
type Rule struct {

	// 用户类型
	RuleType int

	// 账号，账号这里指邮箱
	Account string

	jwt.StandardClaims
}

//用户类型常量
const (
	// Admin 管理员
	Admin = iota + 1
	// User 用户
	User
	// Tourist 未登录的游客
	Tourist
)

// SigningKey 加密私钥
const (
	SigningKey = "the key"
)

// GenerateToken 根据角色生成token
func GenerateToken(ruleType int, account string) (rsult string, err error) {

	rule := Rule{
		RuleType: ruleType,
		Account:  account,
		StandardClaims: jwt.StandardClaims{

			//开始生效时间
			NotBefore: time.Now().Unix(),
			//过期时间
			ExpiresAt: time.Now().Unix() + 60*10,

			//签发者
			Issuer: "zhj",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, rule)
	token, err := t.SignedString([]byte(SigningKey))

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return token, nil
}

//	ParseFromToken	解析token,返回角色

func ParseFromToken(token string) (rule *Rule, err error) {

	t, err := jwt.ParseWithClaims(token, &Rule{}, func(token *jwt.Token) (interface{}, error) {

		return SigningKey, nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	rule = t.Claims.(*Rule)

	return rule, nil

}
