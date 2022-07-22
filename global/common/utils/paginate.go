package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func GetNumSize(c *gin.Context) (int64, int64, error) {
	pageNum, err := strconv.ParseInt(c.Query("pageNum"), 10, 64)
	if err != nil {
		return 0, 0, err
	}
	if pageNum <= 0 {
		return 0, 0, err
	}
	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 64)

	if err != nil {
		return 0, 0, err
	}
	if pageSize <= 0 {
		return 0, 0, err
	}
	return pageNum, pageSize, nil

}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	pageNum, pageSize, _ := GetNumSize(c)
	return func(db *gorm.DB) *gorm.DB {
		// calculate the offset
		offset := (pageNum - 1) * pageSize
		// Return the database object with Offset and Limit
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
