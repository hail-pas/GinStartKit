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

type Resp[T any] struct {
	Code    int    `json:"code"`
	Data    T      `json:"data"`
	Message string `json:"message"`
}

type WithPageInfo[T any] struct {
	Resp[T]
	PageInfo PageInfo `json:"pageInfo"`
}

func Response[T any](c *gin.Context, code int, data T, message string, pageSize, pageNum, totalCount int64) {
	if pageSize == -1 && pageNum == -1 && totalCount == -1 {
		c.JSON(http.StatusOK, Resp[T]{
			Code:    code,
			Data:    data,
			Message: message,
		})
	} else {
		c.JSON(http.StatusOK, WithPageInfo[T]{
			Resp: Resp[T]{
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

func WithoutPageInfo[T any](c *gin.Context, code int, data T, message string) {
	Response[T](c, code, data, message, -1, -1, -1)
}

func Ok(c *gin.Context) {
	WithoutPageInfo[any](c, constant.CodeSuccess, nil, constant.MessageSuccess)
}

func OkWithMessage(c *gin.Context, message string) {
	WithoutPageInfo[any](c, constant.CodeSuccess, nil, message)
}

func OkWithData[T any](c *gin.Context, data T) {
	WithoutPageInfo[T](c, constant.CodeSuccess, data, constant.MessageSuccess)
}

func OkWithPageData[T any](c *gin.Context, data T, pageSize, pageNum, totalCount int64) {
	Response[T](c, constant.CodeSuccess, data, constant.MessageSuccess, pageSize, pageNum, totalCount)
}

func Fail(c *gin.Context) {
	WithoutPageInfo[any](c, constant.CodeError, nil, constant.MessageError)
}

func FailWithMessage(c *gin.Context, message string) {
	WithoutPageInfo[any](c, constant.CodeError, nil, message)
}
func BadRequest(c *gin.Context, message string) {
	WithoutPageInfo[any](c, constant.CodeBadRequest, nil, message)
}

func ErrorResp(c *gin.Context, err error) {
	if validateError, ok := err.(validator.ValidationErrors); !ok {
		Response[any](c, constant.CodeBadRequest, nil, err.Error(), -1, -1, -1)
	} else {
		WithoutPageInfo[any](
			c,
			constant.CodeBadRequest,
			nil,
			utils.ObtainFirstValueOfValidationErrorsTranslations(validateError.Translate(global.Translator)),
		)
	}
}
