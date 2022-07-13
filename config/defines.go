package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"net"
)

const (
	Development string = "Development"
	Test        string = "Test"
	Production  string = "Production"
)

const (
	http  string = "http"
	https string = "https"
)

type TCPAddrField struct {
	TcpAddr net.TCPAddr `json:"tcpAddr" yaml:"tcpAddr"` // 服务器地址:端口
}

func RetrieveTcpAddr(tcpAddrString string) net.TCPAddr {
	addr, err := net.ResolveTCPAddr("tcp", tcpAddrString)
	if err != nil {
		panic(err)
	}
	return *addr
}

func RetrieveEnvironment(in string) string {
	switch in {
	case Development, Test, Production:
		return in
	case "D", "d", "Dev", "dev":
		return Development
	case "T", "t", "test":
		return Test
	case "P", "p", "Pro", "pro":
		return Production

	default:
		panic("Wrong value of environment, available: Development | Test | Production")
	}
}

func RetrieveRequestSchema(in string) string {
	switch in {
	case http:
		return http
	case https:
		return https
	default:
		panic("Wrong value of request schema, available: http | https")
	}
}

func RetrieveIP(in string) net.IP {
	ip := net.ParseIP(in)
	if ip == nil {
		panic(fmt.Sprintf("Cannot parse ip from string %s", in))
	}
	return ip
}

func RetrieveIPList(in []string) []net.IP {
	var ipList []net.IP
	for _, ipString := range in {
		ipList = append(ipList, RetrieveIP(ipString))
	}
	return ipList
}

func RetrieveTcpAddrList(in []string) []net.TCPAddr {
	var tcpAddrList []net.TCPAddr
	for _, tcpAddrString := range in {
		tcpAddrList = append(tcpAddrList, RetrieveTcpAddr(tcpAddrString))
	}
	return tcpAddrList
}

type GORMConfig struct {
	MaxIdleConnections int `json:"maxIdleConnections" yaml:"maxIdleConnections"`
	MaxOpenConnections int `json:"maxOpenConnections" yaml:"maxOpenConnections"`
}

type RelationalDatabaseConfig struct {
	TCPAddrField
	Username     string     `json:"user" yaml:"user"`
	Password     string     `json:"password" yaml:"password"`
	DatabaseName string     `json:"databaseName" yaml:"databaseName"`
	Type         string     `json:"type" yaml:"type"`
	PathParam    string     `json:"pathParam" yaml:"pathParam"`
	GORM         GORMConfig `json:"gorm" yaml:"gorm"`
}

func (r RelationalDatabaseConfig) Dsn() string {
	dsn := ""
	switch r.Type {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", r.Username, r.Password, r.TcpAddr.String(), r.DatabaseName, r.PathParam)
	case "Postgresql":
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d %s",
			r.TcpAddr.String(), r.Username, r.Password, r.DatabaseName, r.TcpAddr.Port, r.PathParam,
		)
	}
	return dsn
}

type RedisConfig struct {
	TCPAddrField
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`
}

type CorsConfig struct {
	AllowAll         bool     `json:"allowAll" yaml:"allowAll"`
	AllowMethods     []string `json:"allowMethods" yaml:"allowMethods"`
	AllowOrigins     []string `json:"allowOrigins" yaml:"allowOrigins"`
	AllowHeaders     []string `json:"allowHeaders" yaml:"allowHeaders"`
	ExposeHeaders    []string `json:"exposeHeaders" yaml:"exposeHeaders"`
	AllowCredentials bool     `json:"allowCredentials" yaml:"allowCredentials"`
	MaxAge           int64    `json:"maxAge" yaml:"maxAge"`
}

type SystemConfig struct {
	TCPAddrField             // 服务器地址:端口
	Environment   string     `json:"environment" yaml:"environment" validate:"oneof=Development Test Production"`
	Debug         bool       `json:"debug" yaml:"debug"`
	RequestSchema string     `json:"requestSchema" yaml:"requestSchema" validate:"oneof=http https"`
	Label         string     `json:"label" yaml:"label"`
	Description   string     `json:"description" yaml:"description"`
	WhiteList     []net.IP   `json:"whiteList" yaml:"whiteList"`
	BlackList     []net.IP   `json:"blackList" yaml:"blackList"`
	CorsConfig    CorsConfig `json:"corsConfig" yaml:"corsConfig"`
}

type LogConfig struct {
	Level zerolog.Level `json:"Level" yaml:"Level"`
}

type HbaseConfig struct {
	TcpAddrList []net.TCPAddr `json:"tcpAddrList" yaml:"tcpAddrList"`
}

type KafkaConfig struct {
	TcpAddrList []net.TCPAddr `json:"tcpAddrList" yaml:"tcpAddrList"`
}

type JwtConfig struct {
	AuthHeaderPrefix string `json:"authHeaderPrefix" yaml:"authHeaderPrefix"`
	Secret           string `json:"secret" yaml:"secret"`               // 密钥
	ExpireMinute     int64  `json:"expireMinute" yaml:"expireMinute"`   // 过期时间
	RefreshMinute    int64  `json:"refreshMinute" yaml:"refreshMinute"` // 缓冲时间
	Issuer           string `json:"issuer" yaml:"issuer"`               // 签发者
}
