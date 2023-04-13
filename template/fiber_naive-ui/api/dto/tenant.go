package dto

import "{{$.moduleName}}/core"

// TenantReq租户请求:可根据需要增删改字段
type TenantReq struct {
	core.Request
	ID          int64   `json:"ID"`          // 租户ID
	IDList      []int64 `json:"IDList"`      // 租户ID列表
	Code        string  `json:"code"`        // 编码
	Name        string  `json:"name"`        // 名称
	Description string  `json:"description"` // 描述
	StartTime   []int64 `json:"startTime"`   // 开始时间，毫秒时间戳
	EndTime     []int64 `json:"endTime"`     // 结束时间，毫秒时间戳
	Status      int8    `json:"status"`      // 状态:10启用、20禁用
	StatusList  []int8  `json:"statusList"`  // 状态列表
	CreateTime  []int64 `json:"createTime"`  // 创建时间，毫秒时间戳
	UpdateTime  []int64 `json:"updateTime"`  // 更新时间，毫秒时间戳
}

// TenantResp租户返回:可根据需要增删改字段
type TenantResp struct {
	ID          int64  `json:"ID" gorm:"column:id;"`                   // 租户ID
	Code        string `json:"code" gorm:"column:code;"`               // 编码
	Name        string `json:"name" gorm:"column:name;"`               // 名称
	Description string `json:"description" gorm:"column:description;"` // 描述
	StartTime   int64  `json:"startTime" gorm:"column:start_time;"`    // 开始时间，毫秒时间戳
	EndTime     int64  `json:"endTime" gorm:"column:end_time;"`        // 结束时间，毫秒时间戳
	Status      int8   `json:"status" gorm:"column:status;"`           // 状态:10启用、20禁用
	CreateTime  int64  `json:"createTime" gorm:"column:create_time;"`  // 创建时间，毫秒时间戳
	UpdateTime  int64  `json:"updateTime" gorm:"column:update_time;"`  // 更新时间，毫秒时间戳

}

func (TenantResp) TableName() string {
	return "tenant"
}

// 可添加其他“请求”或“返回”的实体，建议以Req、Resp结尾
