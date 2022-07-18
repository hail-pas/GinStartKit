package service

import (
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
	"time"
)

type UserProxy struct {
	model.User
	TokenExpiredAt time.Time `json:"tokenExpiredAt"`
}

func (userProxy UserProxy) create() model.User {
	global.RelationalDatabase.Select("id").Where("userName = ?", "phoenix").First(&model.User{})
	return global.RelationalDatabase.Create(&userProxy.User).Statement.Dest.(model.User)
}
