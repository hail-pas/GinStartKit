package account

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/api/service/account"
	"github.com/hail-pas/GinStartKit/global"
	_ "github.com/hail-pas/GinStartKit/global/common/request"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/common/utils"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
)

// List 用户列表
// @Tags Account
// @Summary 用户列表
// @ID get-account-list
// @Security Jwt
// @accept application/x-www-form-urlencoded
// @Produce json
// @Param "pageParam" query request.PageRequestIn true "翻页参数"
// @Success 200 {object} response.WithPageInfo[[]account.UserResponseModel] 成功
// @Router /account [get]
func List(c *gin.Context) {
	// []account.UserResponseModel
	var total int64
	var users []account.UserResponseModel
	pageNum, pageSize := utils.GetNumSize(c)
	global.RelationalDatabase.
		Model(&model.User{}).
		Preload("Systems").
		Order("id DESC").
		Scopes(utils.Paginate(c)).
		Find(&users)
	global.RelationalDatabase.Model(&model.User{}).Count(&total)
	response.OkWithPageData(c, users, pageSize, pageNum, total)
}

//func Retrieve(c *gin.Context) {
//	c.ShouldBindQuery()
//}
