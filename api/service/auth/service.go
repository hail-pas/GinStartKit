package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/common/utils"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
	"github.com/jinzhu/copier"
)

func CreateUser(c *gin.Context, userRegisterIn *UserRegisterIn) error {
	var user model.User
	_ = copier.Copy(&user, userRegisterIn)

	// uuid和初始值
	user.UUID, _ = uuid.NewUUID()
	user.Enabled = true
	user.Password = utils.PasswordHash(user.Password)

	var systemWithUsers []model.SystemWithUser

	// 开启事务
	tx := global.RelationalDatabase.Begin()

	userCreateResult := tx.Create(&user)
	if userCreateResult.Error != nil {
		response.ErrorResp(c, userCreateResult.Error)
		tx.Rollback()
		return global.BreakError
	}

	for _, systemId := range userRegisterIn.SystemIds {
		systemWithUser := model.SystemWithUser{}
		systemWithUser.UserId = user.ID
		systemWithUser.SystemId = systemId
		systemWithUsers = append(systemWithUsers, systemWithUser)
	}

	systemWithUsersCreateResult := tx.Create(systemWithUsers)
	if systemWithUsersCreateResult.Error != nil {
		response.ErrorResp(c, systemWithUsersCreateResult.Error)
		tx.Rollback()
		return global.BreakError
	}
	response.Ok(c)

	return tx.Commit().Error
}
