package redis

import "fmt"

type CacheKey struct {
	Key         string
	Description string
}

func (c CacheKey) format(v interface{}) string {
	return fmt.Sprintf(c.Key, v)
}

var (
	UserInfoKey = CacheKey{Key: "CacheInfo:%s", Description: "用户信息"}
)
