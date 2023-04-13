package core

type Request struct {
	Page   int                    `json:"page"`   // 页码从1开始
	Size   int                    `json:"size"`   // 每页项目数
	Sort   string                 `json:"sort"`   // 排序
	With   string                 `json:"with"`   // 需要额外加载的字段
	UserID int64                  `json:"userID"` // 用户ID
	Params map[string]interface{} `json:"params"` //所有参数
}

// Offset 获取分页偏移量
func (request *Request) Offset() int {
	return (request.Page - 1) * request.Size
}
