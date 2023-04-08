package service

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"{{$.moduleName}}/core"
	"{{$.moduleName}}/core/middleware"
	"{{$.moduleName}}/do"
	"{{$.moduleName}}/dto"
)

type authService struct{}

var AuthService = &authService{}

func (authService *authService) Register(register *dto.Register) (*dto.Token, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.MinCost+10)
	if err != nil {
		return nil, errors.New("密码格式错误")
	}

	auth := &do.Auth{}
	auth.TenantID = register.TenantID
	auth.LoginName = register.LoginName
	auth.Phone = register.Phone
	auth.PasswordHash = string(bytes)

	result := core.DB.Create(&auth)
	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry") {
		return nil, errors.New("用户已经存在")
	}

	if auth.ID > 0 {
		token := &dto.Token{}
		token.AccessToken, token.RefreshToken, err = middleware.CreateToken(auth.TenantID, auth.ID, auth.Nickname, true)
		return token, err
	} else {
		return nil, result.Error
	}
}

// Login 通过登录名查找用户详情
func (authService *authService) Login(login *dto.Login) (*dto.Token, error) {
	auth := &do.Auth{}
	result := core.DB.Where("tenant_id =? and login_name = ?", login.TenantID, login.LoginName).First(&auth)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("账号或密码错误")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(auth.PasswordHash), []byte(login.Password)); err != nil {
		return nil, errors.New("账号或密码错误")
	}

	var err error
	token := &dto.Token{}
	token.AccessToken, token.RefreshToken, err = middleware.CreateToken(auth.TenantID, auth.ID, auth.Nickname, true)
	return token, err
}

func (authService *authService) RefreshAccessToken(token *dto.Token) (*dto.Token, error) {
	jwtToken := new(jwt.Token)
	jwtToken, err := jwt.ParseWithClaims(token.RefreshToken, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return viper.GetString("jwt.key"), nil
	})
	if err != nil {
		return nil, err
	}
	if err == nil && jwtToken.Valid {
		return nil, errors.New("RefreshToken过期")
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	if !claims.VerifyIssuer("rt", true) {
		return nil, errors.New("请传入合法的RefreshToken")
	}

	token.AccessToken, token.RefreshToken, err = middleware.CreateToken(int64(claims["TID"].(float64)), int64(claims["UID"].(float64)), claims["name"].(string), true) //RefreshToken本身不能续期
	return token, err
}
