package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func GetNumSize(c *gin.Context) (int64, int64) {
	pageNum, err := strconv.ParseInt(c.Query("pageNum"), 10, 64)
	if err != nil {
		pageNum = 1
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 64)

	if err != nil {
		pageSize = 10
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return pageNum, pageSize

}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	pageNum, pageSize := GetNumSize(c)
	return func(db *gorm.DB) *gorm.DB {
		// calculate the offset
		offset := (pageNum - 1) * pageSize
		// Return the database object with Offset and Limit
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
