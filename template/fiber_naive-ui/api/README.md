# 说明

## 开发说明

### 生成 swagger 文档

方法一：手动执行

```bash
go get -u github.com/swaggo/swag/cmd/swag
# 1.16 or newer
go install github.com/swaggo/swag/cmd/swag@latest

cd {where is the dir of go.mod}
swag init --generalInfo cmd/api/main.go --outputTypes json  --output ./docs --propertyStrategy camelcase
```

方法二：用“go generate”自动执行

```bash
cd ./cmd/api/
go generate
```

### 权限验证

添加：

1. captcha 验证码

在 struct 字段中添加“`validate:"captcha" label:"验证码"`”(例子：dto/auth.go 中的 Captcha 类型)，即可给字段添加验证码，类型中还需要有 CaptchaID 字段。使用方法：参考 web/auth.go 文件的 Register 方法 和 dto/auth.go 文件的 Register 类型

2. dbUnique 数据库唯一验证

在 struct 字段中添加如“`validate:"dbUnique={tableName}:{tableField}&{tableField}->{structField}," label:"{字段说明}"`”(例子：dto/auth.go 中的 Register 类型)，即可给字段添加数据库唯一验证。其中，

{tableName} 需要做唯一验证的数据表名称

{tableField} 需要做唯一验证的数据表字段

{structField} 作为验证值的数据来源，来源于本 struct 对应的字段名称，没有“->{structField}”的话，默认用当前字段的值作为数据来源

{字段说明} 报错时，字段的显示说明

可以用`&`符号来连接多个查询条件

使用方法：参考 web/auth.go 文件的 Register 方法 和 dto/auth.go 文件的 Register 类型。比如：`validate:"dbUnique=user:phone&tenant_id->TenantID,min=11,max=16" label:"手机号"`，用 user 表中的 phone 和 tenant_id 做为查询条件进行唯一值索引查询，其中

## 服务

### api 入口

cmd/api/main.go

### 定时任务

cmd/crontab/main.go
