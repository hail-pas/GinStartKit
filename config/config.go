package config

import (
	"github.com/gin-gonic/gin"
	"github.com/hail-pas/GinStartKit/global/constant"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	httpPkg "net/http"
	"path/filepath"
	"strings"
)

type Config struct {
	BaseDir            string
	ConfigFilePath     string
	RelationalDatabase RelationalDatabaseConfig
	Redis              RedisConfig
	System             SystemConfig
	Log                LogConfig
	Hbase              HbaseConfig
	Kafka              KafkaConfig
	Jwt                JwtConfig
}

func SetConfig(configPath string) *Config {
	if configPath == "" {
		configPath = "./config/content/default.yaml"
	}
	configPath, err := filepath.Abs(configPath)
	if err != nil {
		panic(err)
	}
	path := filepath.Dir(configPath)
	fileName := filepath.Base(configPath)
	splitFileName := strings.Split(fileName, ".")
	if len(splitFileName) != 2 {
		panic("Config path error")
	}
	file, ext := splitFileName[0], splitFileName[1]
	viper.SetConfigName(file)
	viper.AddConfigPath(path)
	viper.SetConfigType(ext)
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	config := Config{}
	config.ConfigFilePath = configPath
	setRelationalDatabase(&config)
	setRedis(&config)
	setSystem(&config)
	setLog(&config)
	setHbase(&config)
	setKafka(&config)
	setJwt(&config)
	return &config
}

//validate := validator.New()
func setRelationalDatabase(c *Config) {
	relationalDatabaseConfig := RelationalDatabaseConfig{}
	relationalDatabaseConfig.TcpAddr = RetrieveTcpAddr(viper.GetString("RelationalDatabase.Addr"))
	relationalDatabaseConfig.Username = viper.GetString("RelationalDatabase.Username")
	relationalDatabaseConfig.DatabaseName = viper.GetString("RelationalDatabase.DatabaseName")
	relationalDatabaseConfig.Password = viper.GetString("RelationalDatabase.Password")
	relationalDatabaseConfig.Type = RetrieveRelationalDatabaseType(viper.GetString("RelationalDatabase.Type"))
	if relationalDatabaseConfig.Type == "" {
		relationalDatabaseConfig.Type = "postgres"
	}
	relationalDatabaseConfig.PathParam = viper.GetStringMapStringSlice("RelationalDatabase.PathParam")
	c.RelationalDatabase = relationalDatabaseConfig
}

func setRedis(c *Config) {
	redisConfig := RedisConfig{}
	redisConfig.TcpAddr = RetrieveTcpAddr(viper.GetString("Redis.Addr"))
	redisConfig.DB = viper.GetInt("Redis.DB")
	redisConfig.User = viper.GetString("Redis.Username")
	redisConfig.Password = viper.GetString("Redis.Password")
	c.Redis = redisConfig
}

func setSystem(c *Config) {
	systemConfig := SystemConfig{}
	systemConfig.TcpAddr = RetrieveTcpAddr(viper.GetString("System.Addr"))
	systemConfig.Environment = RetrieveEnvironment(viper.GetString("System.Environment"))
	systemConfig.Debug = viper.GetBool("System.Debug")
	if systemConfig.Debug && systemConfig.Environment == Production {
		panic("Cannot set debug mode when environment is Production")
	}
	if systemConfig.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	systemConfig.RequestSchema = RetrieveRequestSchema(viper.GetString("System.RequestSchema"))
	systemConfig.Label = viper.GetString("System.Label")
	systemConfig.Description = viper.GetString("System.Description")
	systemConfig.WhiteList = RetrieveIPList(viper.GetStringSlice("System.WhiteList"))
	systemConfig.BlackList = RetrieveIPList(viper.GetStringSlice("System.BlackList"))

	// CorsConfig
	corsConfig := CorsConfig{}
	corsConfig.AllowAll = viper.GetBool("System.Cors.AllowAll")
	if !corsConfig.AllowAll {
		AllowOrigins := viper.GetStringSlice("System.Cors.AllowOrigins")
		if AllowOrigins == nil {
			corsConfig.AllowOrigins = []string{"*"}
		} else {
			corsConfig.AllowOrigins = AllowOrigins
		}
		AllowMethods := viper.GetStringSlice("System.Cors.AllowMethods")
		if AllowMethods == nil {
			corsConfig.AllowMethods = []string{
				httpPkg.MethodGet, httpPkg.MethodHead, httpPkg.MethodPost, httpPkg.MethodPut, httpPkg.MethodPatch,
				httpPkg.MethodDelete, httpPkg.MethodConnect, httpPkg.MethodOptions, httpPkg.MethodTrace,
			}
		} else {
			corsConfig.AllowMethods = AllowMethods
		}
		AllowHeaders := viper.GetStringSlice("System.Cors.AllowHeaders")
		if AllowHeaders == nil {
			corsConfig.AllowHeaders = []string{
				constant.HeaderAccept, constant.HeaderAcceptEncoding, constant.HeaderAuthorization, constant.HeaderDate,
				constant.HeaderFrom, constant.HeaderHost, constant.HeaderAcceptCharset, constant.HeaderAcceptLanguage,
				constant.HeaderContentType, constant.HeaderCookie, constant.HeaderLocation, constant.HeaderPragma,
				constant.HeaderReferer, constant.HeaderUserAgent,
			}
		} else {
			corsConfig.AllowHeaders = AllowHeaders
		}
		ExposeHeaders := viper.GetStringSlice("System.Cors.ExposeHeaders")
		if ExposeHeaders == nil {
			corsConfig.ExposeHeaders = append(
				corsConfig.AllowHeaders,
				[]string{
					constant.HeaderContentLength, constant.HeaderAccessControlAllowOrigin,
					constant.HeaderAccessControlAllowHeaders, constant.HeaderAccessControlAllowMethods,
					constant.HeaderAccessControlAllowCredentials, constant.HeaderAccessControlExposeHeaders,
				}...,
			)
		} else {
			corsConfig.ExposeHeaders = ExposeHeaders
		}
		corsConfig.AllowCredentials = viper.GetBool("System.Cors.AllowCredentials")
		corsConfig.MaxAge = viper.GetInt64("System.Cors.MaxAge")
	}

	systemConfig.CorsConfig = corsConfig
	c.System = systemConfig
}

func setLog(c *Config) {
	logConfig := LogConfig{}

	logLevel, err := zerolog.ParseLevel(viper.GetString("Log.Level"))

	if err != nil {
		panic(err)
	}

	zerolog.SetGlobalLevel(logLevel)

	log.Logger = log.With().Caller().Logger()

	c.Log = logConfig
}

func setHbase(c *Config) {
	hbaseConfig := HbaseConfig{}
	hbaseConfig.TcpAddrList = RetrieveTcpAddrList(viper.GetStringSlice("Hbase.AddrList"))
	c.Hbase = hbaseConfig
}

func setKafka(c *Config) {
	kafkaConfig := KafkaConfig{}
	kafkaConfig.TcpAddrList = RetrieveTcpAddrList(viper.GetStringSlice("Kafka.AddrList"))
	c.Kafka = kafkaConfig
}

func setJwt(c *Config) {
	jwtConfig := JwtConfig{}
	jwtConfig.Secret = viper.GetString("Jwt.Secret")
	jwtConfig.AuthHeaderPrefix = RetrieveAuthPrefixHeader(viper.GetString("Jwt.AuthHeaderPrefix"))
	jwtConfig.ExpireMinute = viper.GetInt64("Jwt.ExpireMinute")
	jwtConfig.RefreshMinute = viper.GetInt64("Jwt.RefreshMinute")
	jwtConfig.Issuer = viper.GetString("Jwt.Issuer")
	c.Jwt = jwtConfig
}
