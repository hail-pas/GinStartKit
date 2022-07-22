package account

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/common/utils"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
)

func List(c *gin.Context) {
	var users []model.User
	global.RelationalDatabase.Find(&users).Scopes(utils.Paginate(c))
	response.OkWithPageData(c, users, 12, 10, 129)
}
