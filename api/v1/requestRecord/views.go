package requestRecord

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/common/utils"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
)

func List(c *gin.Context) {
	var total int64
	var requestRecords []model.RequestRecord
	pageNum, pageSize := utils.GetNumSize(c)
	global.RelationalDatabase.Model(&model.RequestRecord{}).Order("id DESC").Scopes(utils.Paginate(c)).Find(&requestRecords)
	global.RelationalDatabase.Model(&model.RequestRecord{}).Count(&total)
	response.OkWithPageData(c, requestRecords, pageSize, pageNum, total)
}
