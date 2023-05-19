package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/dentonliu/go-clean-starter/internal/util"
)

type JWTOptions struct {
	Realm         string
	SigningMethod string
	TokenHandler  func(*gin.Context, *jwt.Token) error
}

func JWTAuth(verificationKey string, options ...JWTOptions) gin.HandlerFunc {
	var opt JWTOptions
	if len(options) > 0 {
		opt = options[0]
	}
	if opt.Realm == "" {
		opt.Realm = "API"
	}
	if opt.SigningMethod == "" {
		opt.SigningMethod = "HS256"
	}
	if opt.TokenHandler == nil {
		opt.TokenHandler = DefaultJWTTokenHandler
	}
	parser := &jwt.Parser{
		ValidMethods: []string{opt.SigningMethod},
	}

	return func(c *gin.Context) {
		message := ""
		userToken := ""

		// 支持在请求头或者query参数中传值
		header := c.GetHeader("Authorization")
		if strings.HasPrefix(header, "Bearer ") {
			userToken = header[7:]
		} else {
			userToken = c.Query("token")
		}

		token, err := parser.Parse(userToken, func(t *jwt.Token) (interface{}, error) {
			return []byte(verificationKey), nil
		})
		if err == nil && token.Valid {
			err = opt.TokenHandler(c, token)
		}
		if err == nil {
			return
		}

		message = err.Error()

		c.Header("WWW-Authenticate", `Bearer realm="`+opt.Realm+`"`)

		if message != "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New(message))
			return
		}

		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func DefaultJWTTokenHandler(c *gin.Context, token *jwt.Token) error {
	jwtMap := token.Claims.(jwt.MapClaims)

	if float64(time.Now().Unix()) > jwtMap["exp"].(float64) {
		return errors.New("Token is expired.")
	}

	user := util.User{}

	user.ID = jwtMap["id"].(string)
	user.IP = c.ClientIP()

	c.Set(util.CKUser, &user)
	return nil
}
