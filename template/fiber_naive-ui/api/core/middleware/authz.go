package middleware

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"{{$.moduleName}}/core"
	"{{$.moduleName}}/initialize"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/log"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var Enforcer *casbin.Enforcer

// Auth 权限验证保护
func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tid, sub, _, err := GetByContextKey(c)
		if err != nil {
			return core.Resp(c, core.CodeErrorToken, err.Error(), nil)
		}
		if ok, err := Enforcer.Enforce(strconv.FormatInt(sub, 10), strconv.FormatInt(tid, 10), c.Route().Path, c.Method()); err != nil {
			return core.ErrorSystem(c, err.Error())
		} else if !ok {
			return core.Resp(c, core.CodeErrorForbidden, "无此权限，拒绝访问", nil)
		}
		return c.Next()
	}
}

func OnInitAuthz(db *gorm.DB) {
	type CasbinRule struct {
		ID    uint   `gorm:"primaryKey;autoIncrement"`
		Ptype string `gorm:"size:100;uniqueIndex:unique_index"`
		V0    string `gorm:"size:100;uniqueIndex:unique_index"`
		V1    string `gorm:"size:100;uniqueIndex:unique_index"`
		V2    string `gorm:"size:100;uniqueIndex:unique_index"`
		V3    string `gorm:"size:100;uniqueIndex:unique_index"`
		V4    string `gorm:"size:100;uniqueIndex:unique_index"`
		V5    string `gorm:"size:100;uniqueIndex:unique_index"`
	}
	adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRule{}, viper.GetString("authz.casbinRuleTable"))

	m, _ := model.NewModelFromString(viper.GetString("authz.model"))
	Enforcer, _ = casbin.NewEnforcer(m, adapter)

	Enforcer.EnableAutoSave(true)

	logger := &log.DefaultLogger{}
	logger.EnableLog(true)
	Enforcer.SetLogger(logger)

	initialize.InitPolicy(Enforcer)
}
