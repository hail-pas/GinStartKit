package account

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/api/service/account"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/common/utils"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
)

func List(c *gin.Context) {
	var total int64
	var users []account.UserResponseModel
	pageNum, pageSize := utils.GetNumSize(c)
	global.RelationalDatabase.Model(&model.User{}).Order("id DESC").Scopes(utils.Paginate(c)).Find(&users)
	global.RelationalDatabase.Model(&model.User{}).Count(&total)
	response.OkWithPageData(c, users, pageSize, pageNum, total)
}
