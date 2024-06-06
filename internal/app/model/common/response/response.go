package response

import (
	"net/http"

	"github.com/championlong/go-quick-start/internal/pkg/constants"
	"github.com/gin-gonic/gin"
)

var (
	MetaOk = Meta{Code: SUCCESS, Message: "OK"}
)

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	meta := &Meta{
		Code:    code,
		Message: msg,
	}
	response := MakeResponse(data, meta, nil)
	// 开始时间
	c.JSON(http.StatusOK, response)
}

func MakeResponse(object interface{}, meta *Meta, err *constants.Error) Response {
	var response Response
	if err != nil {
		response.Meta = Meta{
			Code:    err.Code,
			Message: err.Message,
		}
		return response
	}
	if meta != nil {
		response.Meta = *meta
	}
	if object != nil {
		response.Data = object
	}
	return response
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
