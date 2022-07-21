package utils

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/api/service"
	"github.com/hail-pas/GinStartKit/global"
	"github.com/hail-pas/GinStartKit/global/common/response"
	"github.com/hail-pas/GinStartKit/storage/relational/model"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	IdentityKey = "uuid"
)

func GetAuthJwtMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*service.UserProxy); ok {
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
			identity := claims[IdentityKey]
			return identity
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginData model.UserLoginWithPhone
			if err := c.ShouldBind(&loginData); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			var user model.User

			global.RelationalDatabase.Where("phone = ?", loginData.Phone).First(&user)

			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))

			if err == nil {
				return user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			var exists bool

			v, ok := data.(*service.UserProxy)
			if !ok {
				return false
			}
			err := global.RelationalDatabase.Model(model.User{}).Select("id").Where("id = ?", v.ID).Find(&exists).Error
			if err != nil {
				log.Info().Msgf("User doesnt exists with id %v", v.ID)
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
			userProxy, _ := c.Get("user")
			userProxy = userProxy.(service.UserProxy)
			data := map[string]interface{}{
				"token":    token,
				"expireAt": expire,
				"user":     userProxy,
			}
			response.OkWithData(c, data)
		},
	})

	if err != nil {
		panic(err)
	}

	err = authMiddleware.MiddlewareInit()

	if err != nil {
		panic(err)
	}

	return authMiddleware
}
