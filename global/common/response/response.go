package response

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/utils"
	"github.com/hail-pas/GinStartKit/global/constant"
	"math"
	"net/http"
)

type PageInfo struct {
	TotalPage  int64 `json:"totalPage"`
	TotalCount int64 `json:"totalCount"`
	PageNum    int64 `json:"pageNum"`
	PageSize   int64 `json:"pageSize"`
}

type response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type WithPageInfo struct {
	response
	PageInfo PageInfo `json:"pageInfo"`
}

func Response(c *gin.Context, code int, data interface{}, message string, pageSize, pageNum, totalCount int64) {
	if pageSize == -1 && pageNum == -1 && totalCount == -1 {
		c.JSON(http.StatusOK, response{
			Code:    code,
			Data:    data,
			Message: message,
		})
	} else {
		c.JSON(http.StatusOK, WithPageInfo{
			response: response{
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
	Response(c, constant.CodeSuccess, nil, constant.MessageSuccess, -1, -1, -1)
}

func OkWithMessage(c *gin.Context, message string) {
	Response(c, constant.CodeSuccess, nil, message, -1, -1, -1)
}

func OkWithData(c *gin.Context, data interface{}) {
	Response(c, constant.CodeSuccess, data, constant.MessageSuccess, -1, -1, -1)
}

func OkWithPageData(c *gin.Context, data interface{}, pageSize, pageNum, totalCount int64) {
	Response(c, constant.CodeSuccess, data, constant.MessageSuccess, pageSize, pageNum, totalCount)
}

func Fail(c *gin.Context) {
	Response(c, constant.CodeError, nil, constant.MessageError, -1, -1, -1)
}

func FailWithMessage(c *gin.Context, message string) {
	Response(c, constant.CodeError, nil, message, -1, -1, -1)
}
func BadRequest(c *gin.Context, message string) {
	Response(c, constant.CodeBadRequest, nil, message, -1, -1, -1)
}

func ErrorResp(c *gin.Context, err error) {
	if validateError, ok := err.(validator.ValidationErrors); !ok {
		Response(c, constant.CodeBadRequest, nil, err.Error(), -1, -1, -1)
	} else {
		Response(
			c,
			constant.CodeBadRequest,
			nil,
			utils.ObtainFirstValueOfValidationErrorsTranslations(validateError.Translate(global.Translator)),
			-1, -1, -1,
		)
	}
}
