package request

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/api/schema/common/response"
	"github.com/hail-pas/GinStartKit/global/constant"
	"strconv"
)

type PageRequestIn struct {
	PageNum  int64 `json:"pageNum" form:"PageNum" validate:"required,gt=0"`   // 页码
	PageSize int64 `json:"pageSize" form:"pageSize" validate:"required,gt=0"` // 页长
}

func (p PageRequestIn) GetNumSize(c *gin.Context) (int64, int64) {
	pageNum, err := strconv.ParseInt(c.Query("pageNum"), 10, 64)
	if err != nil {
		response.BadRequest(c, constant.PageNumError)
		c.Abort()
	}
	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 64)

	if err != nil {
		response.BadRequest(c, constant.PageSizeError)
		c.Abort()
	}

	return pageNum, pageSize

}

// IdRequestIn Find by id structure
type IdRequestIn struct {
	ID int `json:"id" form:"id" validate:"required,gt=0"`
}

func (r *IdRequestIn) Uint() uint {
	return uint(r.ID)
}

type IdsRequestIn struct {
	Ids []int `json:"ids" form:"ids"`
}

type Empty struct{}
