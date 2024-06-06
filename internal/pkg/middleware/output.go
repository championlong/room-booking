package middleware

import (
	"github.com/championlong/go-quick-start/internal/app/model/common/response"
	"github.com/championlong/go-quick-start/internal/pkg/constants"
	"github.com/championlong/go-quick-start/pkg/log"
	"github.com/gin-gonic/gin"
)

func Output() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		req, _ := c.Get(constants.ContextKeyRequest)
		resp, _ := c.Get(constants.ContextKeyResponse)
		meta, _ := c.Get(constants.ContextKeyMeta)
		err := c.Errors.Last()
		if err != nil {
			e, ok := err.Err.(*constants.Error)
			if !ok {
				e = constants.ErrInternalServerError
			}
			envelope(c, response.MakeResponse(nil, nil, e))
		} else {
			m, ok := meta.(response.Meta)
			if ok {
				envelope(c, response.MakeResponse(resp, &m, nil))
			} else {
				envelope(c, response.MakeResponse(resp, &response.MetaOk, nil))

			}
		}
		printLog(c, req, resp, err)

	}
}

type Log struct {
	Request  interface{} `json:"request"`
	Response interface{} `json:"response"`
	Error    interface{} `json:"error"`
	Path     string      `json:"path"`
}

func printLog(c *gin.Context, req interface{}, resp interface{}, err interface{}) {
	logInfo := Log{
		Request:  req,
		Response: resp,
		Error:    err,
		Path:     c.Request.URL.Path,
	}
	log.WithContext(c.Request.Context()).Infow("request log", log.Any("log", logInfo))
}

func envelope(c *gin.Context, response response.Response) {
	code := response.Meta.Code
	if code >= 100000 {
		code = code / 1000
	}
	c.JSON(code, response)
}
