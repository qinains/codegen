# 说明

## 开发说明

### 生成swagger文档

方法一：手动执行

```bash
go get -u github.com/swaggo/swag/cmd/swag
# 1.16 or newer
go install github.com/swaggo/swag/cmd/swag@latest

cd {where is the dir of go.mod}
swag init --generalInfo cmd/api/main.go --outputTypes json  --output ./docs --propertyStrategy pascalcase
```

方法二：用“go generate”自动执行

```bash
cd ./cmd/api/
go generate
```

## 服务

### api入口

cmd/api/main.go

### 定时任务

cmd/crontab/main.go
