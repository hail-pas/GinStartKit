package middleware

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/global/common/utils"
	"github.com/hail-pas/GinStartKit/global/constant"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
	"time"
)

const (
	IdentityKey = "user"
)

var authMiddleware *jwt.GinJWTMiddleware

func GetAuthJwtMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(model.User); ok {
				return jwt.MapClaims{
					IdentityKey: v.UUID,
					"username":  v.Username,
					"phone":     v.Phone,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			uuid := claims[IdentityKey]
			var user model.User
			global.RelationalDatabase.Where("uuid = ?", uuid).First(&user)
			if user.ID == 0 {
				return nil
			}
			return user
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginData model.UserLoginWithPhone
			if err := c.ShouldBindJSON(&loginData); err != nil {
				return "", errors.New("帐号和密码字段必填")
			}

			var user model.User

			global.RelationalDatabase.Where("phone = ?", loginData.Phone).First(&user)

			if utils.VerifyHashAndPassword(user.Password, loginData.Password) {
				c.Set("user", user)
				return user, nil
			}

			return nil, errors.New("帐号或密码错误")
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			var exists model.User
			user, ok := data.(model.User)
			if !ok {
				return false
			}
			err := global.RelationalDatabase.Model(model.User{}).Select("id").Where(
				"uuid = ?",
				user.UUID.String(),
			).Find(&exists).Error
			if err != nil || exists.ID == 0 {
				return false
			}
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			response.Response(c, code, nil, message, -1, -1, -1)
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: global.Configuration.Jwt.AuthHeaderPrefix,

		TimeFunc: time.Now,
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			user, _ := c.Get("user")
			data := map[string]interface{}{
				"token":    token,
				"expireAt": expire,
				"user":     user,
			}
			response.OkWithData(c, data)
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			data := map[string]interface{}{
				"token":    token,
				"expireAt": expire,
			}
			response.OkWithData(c, data)
		},
	})

	err := authMiddleware.MiddlewareInit()

	if err != nil {
		panic(err)
	}

	return authMiddleware
}

func unauthorized(c *gin.Context, code int, message string) {
	c.Header("WWW-Authenticate", "JWT realm="+authMiddleware.Realm)
	if !authMiddleware.DisabledAbort {
		c.Abort()
	}

	authMiddleware.Unauthorized(c, code, message)
}

func AuthMiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := authMiddleware.GetClaimsFromJWT(c)
		if err != nil {
			unauthorized(c, constant.CodeUnauthorized, authMiddleware.HTTPStatusMessageFunc(err, c))
			return
		}

		if claims["exp"] == nil {
			unauthorized(c, constant.CodeBadRequest, authMiddleware.HTTPStatusMessageFunc(errors.New(constant.MessageTokenParseFailed), c))
			return
		}

		if _, ok := claims["exp"].(float64); !ok {
			unauthorized(c, constant.CodeBadRequest, authMiddleware.HTTPStatusMessageFunc(errors.New(constant.MessageTokenParseFailed), c))
			return
		}

		if int64(claims["exp"].(float64)) < authMiddleware.TimeFunc().Unix() {
			unauthorized(c, constant.CodeUnauthorized, authMiddleware.HTTPStatusMessageFunc(errors.New(constant.MessageTokenExpire), c))
			return
		}

		c.Set("JWT_PAYLOAD", claims)
		identity := authMiddleware.IdentityHandler(c)

		if identity != nil {
			c.Set(authMiddleware.IdentityKey, identity)
		}

		if !authMiddleware.Authorizator(identity, c) {
			unauthorized(c, constant.CodeForbidden, authMiddleware.HTTPStatusMessageFunc(errors.New(constant.MessageForbidden), c))
			return
		}
		c.Next()
	}
}
