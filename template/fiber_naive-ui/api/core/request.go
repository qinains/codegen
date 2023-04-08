package core

type Request struct {
	Page   int                    // 页码从1开始
	Size   int                    // 每页项目数
	Sort   string                 // 排序
	With   string                 // 需要额外加载的字段
	UserID int64                  // 用户ID
	Params map[string]interface{} //所有参数
}

// Offset 获取分页偏移量
func (request *Request) Offset() int {
	return (request.Page - 1) * request.Size
}
