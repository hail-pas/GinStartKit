package model

import (
	"net"
	"time"
)

type OperationRecord struct {
	BaseModel
	Ip           net.IP        `json:"ip" form:"ip"`                     // 请求ip
	Method       string        `json:"method" form:"method"`             // 请求方法
	Path         string        `json:"path" form:"path"`                 // 请求路径
	Status       int           `json:"status" form:"status"`             // 请求状态
	Latency      time.Duration `json:"latency" form:"latency"`           // 延迟
	Agent        string        `json:"agent" form:"agent"`               // 代理
	ErrorMessage string        `json:"errorMessage" form:"errorMessage"` // 错误信息
	Body         string        `json:"body" form:"body"`                 // 请求Body
	Resp         string        `json:"resp" form:"resp"`                 // 响应Body
	UserID       int64         `json:"userId" form:"userId"`             // 用户id
}
