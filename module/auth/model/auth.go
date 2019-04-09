package model

import (
	"errors"
	"regexp"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

var (
	// JWTkey key
	JWTkey string = "asjdkasjdkajskdjksajd"
)

// JWTPayload struct
type JWTPayload struct {
	*jwt.StandardClaims
	UserID       int    `json:"user_id"`
	UserFullName string `json:"user_fullname"`
}

var jwtSecretKey = beego.AppConfig.String("jwtsecretkey")
var tokenValidTime = beego.AppConfig.String("tokenexpireminutes")

// JWTClaims Token claims model
type JWTClaims struct {
	Username string
	jwt.StandardClaims
}

// Token JWT Token Model
type Token struct {
	TokenString string
	Key         string
}

// TokenGenerator token generator model
type TokenGenerator struct {
	Key string
	Alg jwt.SigningMethod
}

// ParseJWTClaim func
func ParseJWTClaim(tokenString string) (jwt.MapClaims, error) {

	re := regexp.MustCompile("(?i)bearer\\s+")
	refinedToken := re.ReplaceAllString(tokenString, "")

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(refinedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTkey), nil
	})

	if err == nil {
		return claims, nil
	}

	return nil, err
}

func parseToken(t *Token) (token *jwt.Token, err error) {
	token, err = jwt.Parse(t.TokenString, func(jt *jwt.Token) (interface{}, error) {
		// Untuk mencegah JWT Signing method NONE attack
		// Maka pastikan untuk memvalidasi juga Algoritma signing nya
		if _, ok := jt.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(t.Key), nil
		}
		return nil, errors.New("Unknown Signing Method")
	})
	return
}

// GenerateToken generate token
func (gen *TokenGenerator) GenerateToken(claims jwt.Claims) (token string, err error) {
	t := jwt.NewWithClaims(gen.Alg, claims)
	token, err = t.SignedString([]byte(gen.Key))
	return
}

// GetPayload func
func (t *Token) GetPayload() jwt.MapClaims {

	if token, err := parseToken(t); err == nil && token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			return claims
		}
	}

	return nil
}

// IsValidToken validate JWT Token
func (t *Token) IsValidToken() bool {
	if token, err := parseToken(t); err == nil {
		return token.Valid
	}
	return false
}

// JwtDecodeHelper untuk membantu decode payload data jwt
func JwtDecodeHelper(tokenString string) (jwt.Claims, error) {

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("jwtsecretkey"), nil
	})

	if err != nil {
		return token.Claims, nil
	}

	return nil, err
}

func initTokenGenerator() *TokenGenerator {
	t := &TokenGenerator{
		Key: jwtSecretKey,
		Alg: jwt.SigningMethodHS256,
	}
	return t
}
