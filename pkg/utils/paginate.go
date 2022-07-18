package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/constant"
	"strconv"
)

func GetNumSize(c *gin.Context) (int64, int64) {
	pageNum, err := strconv.ParseInt(c.Query("pageNum"), 10, 64)
	if err != nil {
		response.BadRequest(c, constant.PageNumError)
		c.Abort()
	}
	if pageNum <= 0 {
		response.BadRequest(c, constant.PageNumError)
		c.Abort()
	}
	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 64)

	if err != nil {
		response.BadRequest(c, constant.PageSizeError)
		c.Abort()
	}
	if pageSize <= 0 {
		response.BadRequest(c, constant.PageSizeError)
		c.Abort()
	}
	return pageNum, pageSize

}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Read the page no. query parameter
		pageNum, pageSize := GetNumSize(c)
		// calculate the offset
		offset := (pageNum - 1) * pageSize
		// Return the database object with Offset and Limit
		return db.Offset(offset).Limit(pageSize)
	}
}
