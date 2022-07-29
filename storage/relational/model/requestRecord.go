package model

import (
	"github.com/jackc/pgtype"
)

type RequestRecord struct {
	BaseModel
	ClientIp   *pgtype.Inet     `json:"clientIp"`   // 请求ip
	Method     string           `json:"method"`     // 请求方法
	Path       string           `json:"path"`       // 请求路径
	StatusCode int              `json:"statusCode"` // 请求状态
	Latency    *pgtype.Interval `json:"latency"`    // 延迟
	Agent      string           `json:"agent"`      // 代理
	Query      *pgtype.JSON     `json:"param"`
	FormData   *pgtype.JSON     `json:"formData"`
	Body       *pgtype.JSON     `json:"body"`   // 请求Body
	Resp       *pgtype.JSON     `json:"resp"`   // 响应Body
	UserID     int64            `json:"userId"` // 用户id
}
