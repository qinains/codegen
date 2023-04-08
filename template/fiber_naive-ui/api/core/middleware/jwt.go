package middleware

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"{{$.moduleName}}/core"
)

// JWT 验证路由保护
func JWT() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(viper.GetString("jwt.key")),
		ContextKey: viper.GetString("jwt.contextKey"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return core.ErrorSystem(c, "验证信息丢失")
			} else {
				return core.ErrorSystem(c, "验证失败或过期")
			}
		},
	})
}

// CreateToken 创建token
func CreateToken(TID, UID int64, name string, isCreateRefreshToken bool) (accessToken, refreshToken string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["TID"] = TID
	claims["UID"] = UID
	claims["name"] = name

	claims["exp"] = time.Now().Add(time.Second * time.Duration(viper.GetInt64("jwt.accessTokenExpireIn"))).Unix()
	accessToken, err = token.SignedString([]byte(viper.GetString("jwt.key")))
	if err != nil {
		return
	}

	if isCreateRefreshToken {
		claims["exp"] = time.Now().Add(time.Second * time.Duration(viper.GetInt64("jwt.refreshTokenExpireIn"))).Unix()
		claims["iss"] = "rt"
		refreshToken, err = token.SignedString([]byte(viper.GetString("jwt.key")))
	}
	return
}

func GetByContextKey(c *fiber.Ctx) (TID, UID int64, name string, err error) {
	token := c.Locals(viper.GetString("jwt.contextKey"))
	if token == nil {
		return 0, 0, "", errors.New("Token错误")
	}
	claims := token.(*jwt.Token).Claims.(jwt.MapClaims)
	TID = int64(claims["TID"].(float64))
	UID = int64(claims["UID"].(float64))
	name = claims["name"].(string)
	return
}
