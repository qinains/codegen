package service

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"{[{$.moduleName}]}/core"
	"{[{$.moduleName}]}/core/middleware"
	"{[{$.moduleName}]}/do"
	"{[{$.moduleName}]}/dto"
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
	result := core.DB.Where(&dto.AuthResp{TenantID: login.TenantID, LoginName: login.LoginName}).Or(&dto.AuthResp{TenantID: login.TenantID, Phone: login.LoginName}).First(&auth)

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
		return []byte(viper.GetString("jwt.key")), nil
	})
	if err != nil {
		return nil, err
	}
	if !jwtToken.Valid {
		return nil, errors.New("RefreshToken过期")
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	if !claims.VerifyIssuer("rt", true) {
		return nil, errors.New("请传入合法的RefreshToken")
	}

	token.AccessToken, _, err = middleware.CreateToken(int64(claims["TID"].(float64)), int64(claims["UID"].(float64)), claims["name"].(string), false) //RefreshToken本身不续期
	return token, err
}

// GetByPhoneOrLoginName 根据手机号或登录名获取用户详情
func (authService *authService) GetByPhoneOrLoginName(tenantID int64, phone, loginName string) (*dto.AuthResp, error) {
	auth := &dto.AuthResp{}
	result := core.DB.Where(&dto.AuthResp{TenantID: tenantID, LoginName: loginName}).Or(&dto.AuthResp{TenantID: tenantID, Phone: phone}).First(&auth)
	return auth, result.Error
}

// Get 获取用户详情
func (authService *authService) Get(ID int64) (*dto.AuthInfoResp, error) {
	var auth *dto.AuthInfoResp
	result := core.DB.First(&auth, ID)
	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	auth.Permissions = &[]dto.Permission{{Label: "主控台", Value: "dashboard_console"}, {Label: "监控页", Value: "dashboard_monitor"}, {Label: "工作台", Value: "dashboard_workplace"}, {Label: "基础列表", Value: "basic_list"}, {Label: "基础列表删除", Value: "basic_list_delete"}}
	if auth.Nickname == "" {
		auth.Nickname = auth.LoginName
	}
	return auth, result.Error
}
