package response

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

type PageInfo struct {
	TotalPage  int64 `json:"totalPage"`
	TotalCount int64 `json:"totalCount"`
	PageNum    int64 `json:"pageNum"`
	PageSize   int64 `json:"pageSize"`
}

const (
	SUCCESS = 000
	ERROR   = 001
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type WithPageInfo struct {
	Response
	PageInfo PageInfo `json:"pageInfo"`
}

func response(c *gin.Context, code int, data interface{}, message string, pageSize, pageNum, totalCount int64) {
	if pageSize == -1 && pageNum == -1 && totalCount == -1 {
		c.JSON(http.StatusOK, Response{
			Code:    code,
			Data:    data,
			Message: message,
		})
	} else {
		c.JSON(http.StatusOK, WithPageInfo{
			Response: Response{
				Code:    code,
				Data:    data,
				Message: message,
			},
			PageInfo: PageInfo{
				TotalPage:  int64(math.Ceil(float64(totalCount) / float64(pageSize))),
				TotalCount: totalCount,
				PageNum:    pageNum,
				PageSize:   pageSize,
			},
		})
	}

}

func Ok(c *gin.Context) {
	response(c, SUCCESS, nil, "success", -1, -1, -1)
}

func OkWithMessage(c *gin.Context, message string) {
	response(c, SUCCESS, nil, message, -1, -1, -1)
}

func OkWithData(c *gin.Context, data interface{}) {
	response(c, SUCCESS, data, "", -1, -1, -1)
}

func OkWithPageData(c *gin.Context, data interface{}, pageSize, pageNum, totalCount int64) {
	response(c, SUCCESS, data, "", pageSize, pageNum, totalCount)
}

func Fail(c *gin.Context, message string) {
	response(c, ERROR, nil, message, -1, -1, -1)
}
