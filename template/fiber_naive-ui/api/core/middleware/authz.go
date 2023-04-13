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
			return core.ErrorUnauthorized(c, err.Error(), nil)
		}
		if ok, err := Enforcer.Enforce(strconv.FormatInt(sub, 10), strconv.FormatInt(tid, 10), c.Route().Path, c.Method()); err != nil {
			return core.ErrorSystem(c, err.Error())
		} else if !ok {
			return core.ErrorForbidden(c, "无此权限，拒绝访问", nil)
		}
		return c.Next()
	}
}

func OnInitAuthz(db *gorm.DB) {
	type CasbinRule struct {
		ID    int64  `json:"ID" gorm:"column:id;type:BIGINT UNSIGNED;primaryKey;autoIncrement;not null;comment:权限ID"`          // 权限ID
		Ptype string `json:"ptype" gorm:"column:ptype;type:VARCHAR(100);index:ptype,unique;not null;default:'';comment:PType"` // PType
		V0    string `json:"v0" gorm:"column:v0;type:VARCHAR(100);index:ptype,unique;not null;default:'';comment:v0"`          // v0
		V1    string `json:"v1" gorm:"column:v1;type:VARCHAR(100);index:ptype,unique;not null;default:'';comment:v1"`          // v1
		V2    string `json:"v2" gorm:"column:v2;type:VARCHAR(100);index:ptype,unique;not null;default:'';comment:v2"`          // v2
		V3    string `json:"v3" gorm:"column:v3;type:VARCHAR(100);index:ptype,unique;not null;default:'';comment:v3"`          // v3
		V4    string `json:"v4" gorm:"column:v4;type:VARCHAR(100);index:ptype,unique;not null;default:'';comment:v4"`          // v4
		V5    string `json:"v5" gorm:"column:v5;type:VARCHAR(100);index:ptype,unique;not null;default:'';comment:v5"`          // v5
	}

	adapter, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRule{}, viper.GetString("authz.casbinRuleTable"))

	m, _ := model.NewModelFromString(viper.GetString("authz.model"))
	var err error
	Enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		panic(err)
	}

	Enforcer.EnableAutoSave(true)

	logger := &log.DefaultLogger{}
	logger.EnableLog(true)
	Enforcer.SetLogger(logger)

	initialize.InitPolicy(Enforcer)
}
