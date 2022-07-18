package config

import (
	"fmt"
	"github.com/rs/zerolog"
	"net"
	"net/url"
	"strings"
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

func RetrieveAuthPrefixHeader(in string) string {
	switch in {
	case "Jwt", "Bearer":
		return in
	default:
		panic(fmt.Sprintf("Wrong value of authHeaderPrefix: %s", in))

	}
}

func RetrieveRelationalDatabaseType(in string) string {
	in = strings.ToLower(in)
	switch in {
	case "mysql":
		return in
	case "postgres":
		return in
	default:
		panic(fmt.Sprintf("Unsupported database type: %s", in))
	}
}

type GORMConfig struct {
	MaxIdleConnections int `json:"maxIdleConnections" yaml:"maxIdleConnections"`
	MaxOpenConnections int `json:"maxOpenConnections" yaml:"maxOpenConnections"`
}

type RelationalDatabaseConfig struct {
	TCPAddrField
	Username     string              `json:"user" yaml:"user"`
	Password     string              `json:"password" yaml:"password"`
	DatabaseName string              `json:"databaseName" yaml:"databaseName"`
	Type         string              `json:"type" yaml:"type"`
	PathParam    map[string][]string `json:"pathParam" yaml:"pathParam"`
	GORM         GORMConfig          `json:"gorm" yaml:"gorm"`
}

func (r RelationalDatabaseConfig) Dsn() string {
	var pathValues url.Values
	pathValues = r.PathParam
	dsn := ""
	switch r.Type {
	case "mysql":
		dsn = (&url.URL{
			User:     url.UserPassword(r.Username, r.Password),
			Scheme:   r.Type,
			Host:     r.TcpAddr.String(),
			Path:     r.DatabaseName,
			RawQuery: (&pathValues).Encode(),
		}).String()
		//dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", r.Username, r.Password, r.TcpAddr.String(), r.DatabaseName, r.PathParam)
	case "postgres":
		dsn = (&url.URL{
			User:     url.UserPassword(r.Username, r.Password),
			Scheme:   r.Type,
			Host:     r.TcpAddr.String(),
			Path:     r.DatabaseName,
			RawQuery: (&pathValues).Encode(),
		}).String()
		//dsn = fmt.Sprintf("postgresql://%s:%s@%s/%s?%s", r.Username, r.Password, r.TcpAddr.String(), r.DatabaseName, r.PathParam)
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
	Environment   string     `json:"environment" yaml:"environment"`
	Debug         bool       `json:"debug" yaml:"debug"`
	RequestSchema string     `json:"requestSchema" yaml:"requestSchema"`
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
