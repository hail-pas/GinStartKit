package response

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/utils"
	"github.com/hail-pas/GinStartKit/global/constant"
	"math"
	"net/http"
	"time"
)

type PageInfo struct {
	TotalPage  int64 `json:"totalPage"`
	TotalCount int64 `json:"totalCount"`
	PageNum    int64 `json:"pageNum"`
	PageSize   int64 `json:"pageSize"`
}

type Resp struct {
	Code    int       `json:"code"`
	Data    any       `json:"data" swaggertype:"object"`
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
} //@name Response

type WithPageInfo struct {
	Resp
	PageInfo PageInfo `json:"pageInfo"`
} //@name ResponseWithPageInfo

func Response(c *gin.Context, code int, data any, message string, pageSize, pageNum, totalCount int64) {
	resp := Resp{
		Code:    code,
		Message: message,
		Time:    time.Now(),
	}
	// 用于日志记录
	simpleResp, _ := json.Marshal(resp)
	c.Set("simpleResp", simpleResp)
	resp.Data = data
	if pageSize == -1 && pageNum == -1 && totalCount == -1 {
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusOK, WithPageInfo{
			Resp: resp,
			PageInfo: PageInfo{
				TotalPage:  int64(math.Ceil(float64(totalCount) / float64(pageSize))),
				TotalCount: totalCount,
				PageNum:    pageNum,
				PageSize:   pageSize,
			},
		})
	}
}

func WithoutPageInfo(c *gin.Context, code int, data any, message string) {
	Response(c, code, data, message, -1, -1, -1)
}

func Ok(c *gin.Context) {
	WithoutPageInfo(c, constant.CodeSuccess, nil, constant.MessageSuccess)
}

func OkWithMessage(c *gin.Context, message string) {
	WithoutPageInfo(c, constant.CodeSuccess, nil, message)
}

func OkWithData(c *gin.Context, data any) {
	WithoutPageInfo(c, constant.CodeSuccess, data, constant.MessageSuccess)
}

func OkWithPageData(c *gin.Context, data any, pageSize, pageNum, totalCount int64) {
	Response(c, constant.CodeSuccess, data, constant.MessageSuccess, pageSize, pageNum, totalCount)
}

func Fail(c *gin.Context) {
	WithoutPageInfo(c, constant.CodeError, nil, constant.MessageError)
}

func FailWithMessage(c *gin.Context, message string) {
	WithoutPageInfo(c, constant.CodeError, nil, message)
}
func BadRequest(c *gin.Context, message string) {
	WithoutPageInfo(c, constant.CodeBadRequest, nil, message)
}

func ErrorResp(c *gin.Context, err error) {
	if validateError, ok := err.(validator.ValidationErrors); !ok {
		Response(c, constant.CodeBadRequest, nil, err.Error(), -1, -1, -1)
	} else {
		WithoutPageInfo(
			c,
			constant.CodeBadRequest,
			nil,
			utils.ObtainFirstValueOfValidationErrorsTranslations(validateError.Translate(global.Translator)),
		)
	}
}
