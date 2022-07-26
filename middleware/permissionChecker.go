package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/constant"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
)

func hasPerm(handlerNAme string, userId int64) bool {
	var permission model.Permission
	global.RelationalDatabase.Select("code").Joins("").Where("").First(&permission)
	if permission.Code == handlerNAme {
		return true
	}
	return false
}

func PermissionChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		userTemp, ok := c.Get(IdentityKey)

		if !ok {
			response.Response(c, constant.CodeForbidden, nil, constant.MessageForbidden, -1, -1, -1)
			c.Abort()
		}

		user := userTemp.(model.User)

		handlerName := c.HandlerName()
		if !hasPerm(handlerName, user.ID) {
			response.Response(c, constant.CodeForbidden, nil, constant.MessageForbidden, -1, -1, -1)
			c.Abort()
		}
	}
}
