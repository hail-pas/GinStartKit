package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var respPool sync.Pool

func init() {
	respPool.New = func() interface{} {
		return make([]byte, 1024)
	}
}

func OperationRecord() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId int
		if c.Request.Method != http.MethodGet {
			var err error
			body, err = ioutil.ReadAll(c.Request.Body)
			if err != nil {
				log.Error().Err(err)
			} else {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		// TODO: UserId
		userId = 1
		record := model.OperationRecord{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   string(body),
			UserID: userId,
		}

		// 上传文件时候 中间件日志进行裁断操作
		if strings.Index(c.GetHeader("Content-Type"), "multipart/form-data") > -1 {
			if len(record.Body) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, record.Body)
				record.Body = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if strings.Index(c.Writer.Header().Get("Pragma"), "public") > -1 ||
			strings.Index(c.Writer.Header().Get("Expires"), "0") > -1 ||
			strings.Index(c.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Type"), "application/force-download") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Type"), "application/octet-stream") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Type"), "application/download") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Disposition"), "attachment") > -1 ||
			strings.Index(c.Writer.Header().Get("Content-Transfer-Encoding"), "binary") > -1 {
			if len(record.Resp) > 1024 {
				// 截断
				newBody := respPool.Get().([]byte)
				copy(newBody, record.Resp)
				record.Body = string(newBody)
				defer respPool.Put(newBody[:0])
			}
		}

		global.RelationalDatabase.Create(&record)
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
