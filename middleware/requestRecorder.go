package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
	"github.com/jackc/pgtype"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net"
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

func getQuery(c *gin.Context) []byte {
	rawQuery := c.Request.URL.RawQuery
	rawQuery, _ = url.QueryUnescape(rawQuery)
	split := strings.Split(rawQuery, "&")
	queryMap := make(map[string]string)
	for _, v := range split {
		kv := strings.Split(v, "=")
		if len(kv) == 2 {
			queryMap[kv[0]] = kv[1]
		}
	}
	query, _ := json.Marshal(&queryMap)
	return query
}

func getBody(c *gin.Context) []byte {
	var body []byte = nil
	if c.Request.Method != http.MethodGet {
		var err error
		body, err = ioutil.ReadAll(c.Request.Body)
		if err == nil {
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		}
	}
	return body
}

func requestRecordPgTypePrepare(c *gin.Context) (pgtype.Inet, pgtype.JSON, pgtype.JSON, pgtype.JSON) {
	query := getQuery(c)
	var form map[string][]string
	rawForm, _ := c.MultipartForm()
	if rawForm != nil && rawForm.Value != nil {
		form = rawForm.Value
	}
	//rawForm.File
	body := getBody(c)

	ip := net.ParseIP(c.ClientIP())

	var iNet pgtype.Inet

	_ = iNet.Set(ip)

	var queryJson pgtype.JSON
	var formJson pgtype.JSON
	var bodyJson pgtype.JSON

	_ = queryJson.Set(query)
	_ = formJson.Set(form)
	_ = bodyJson.Set(body)

	if len(queryJson.Bytes) <= 2 {
		queryJson.Bytes = nil
	}

	if len(formJson.Bytes) <= 4 {
		formJson.Bytes = nil
	}
	return iNet, queryJson, formJson, bodyJson
}

func RequestRecorder() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		userTemp, ok := c.Get(IdentityKey)

		if !ok {
			return
		}

		user := userTemp.(model.User)

		iNet, queryJson, formJson, bodyJson := requestRecordPgTypePrepare(c)

		record := model.RequestRecord{
			ClientIp: &iNet,
			Method:   c.Request.Method,
			Path:     c.Request.URL.Path,
			Agent:    c.Request.UserAgent(),
			Query:    &queryJson,
			FormData: &formJson,
			Body:     &bodyJson,
			UserID:   user.ID,
		}

		c.Next()

		latency := time.Since(now)
		interval := &pgtype.Interval{}
		_ = interval.Set(latency)
		record.Latency = interval
		record.StatusCode = c.Writer.Status()

		var respJson pgtype.JSON

		_ = respJson.Set(c.MustGet("simpleResp"))

		record.Resp = &respJson
		global.RelationalDatabase.Create(&record)
		switch {
		case record.StatusCode != 200:
			{
				log.Warn().
					Str("time", time.Now().Format(time.RFC3339)).
					Int("statusCode", record.StatusCode).
					Str("method", record.Method).
					Str("path", record.Path).
					Dur("latency", latency).
					Int64("userId", record.UserID).
					Str("clientIp", c.ClientIP()).
					Str("agent", record.Agent).
					Msg("")
			}
		default:
			log.Info().
				Str("time", time.Now().Format(time.RFC3339)).
				Int("statusCode", record.StatusCode).
				Str("method", record.Method).
				Str("path", record.Path).
				Dur("latency", latency).
				Int64("userId", record.UserID).
				Str("clientIp", c.ClientIP()).
				Str("agent", record.Agent).
				Msg("")
		}
	}
}
