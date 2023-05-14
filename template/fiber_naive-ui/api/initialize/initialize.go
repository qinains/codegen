package initialize

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
	"{[{$.moduleName}]}/do"
)

func InitPolicy(enforcer *casbin.Enforcer) {
	enforcer.AddPermissionForUser("1", "1", "/user/create", "POST")
	enforcer.AddPermissionForUser("1", "1", "/menu", "POST")
	enforcer.AddPermissionForUser("1", "1", "/menu/:ID", "GET")
	enforcer.AddRoleForUser("1", "root", "1")
	enforcer.AddRoleForUser("1", "admin", "1")
	enforcer.AddPolicy("1", "1", "/user/info", "GET")
	fmt.Println(enforcer.GetAllRoles())
}

func InitAutoMigrate(DB *gorm.DB) error {
	return DB.AutoMigrate(&do.Menu{}, &do.Role{}, &do.UserRole{}, &do.Tenant{}, &do.User{}, &do.Dict{}, &do.DictItem{}, &do.Department{}, &do.Job{}, &do.Config{}, &do.File{}, &do.LogLogin{}, &do.LogOperation{})
}
