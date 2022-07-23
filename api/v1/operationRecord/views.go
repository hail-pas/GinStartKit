package operationRecord

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/common/utils"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
)

func List(c *gin.Context) {
	var total int64
	var operationRecords []model.OperationRecord
	pageNum, pageSize := utils.GetNumSize(c)
	global.RelationalDatabase.Model(&model.OperationRecord{}).Scopes(utils.Paginate(c)).Find(&operationRecords)
	global.RelationalDatabase.Model(&model.OperationRecord{}).Count(&total)
	response.OkWithPageData(c, operationRecords, pageSize, pageNum, total)
}
