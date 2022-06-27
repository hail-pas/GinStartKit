package config

type AddrField struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"` // 服务器地址:端口
}

type RelationalDatabaseConfig struct {
	AddrField
	User         string `mapstructure:"user" json:"user" yaml:"user"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	DatabaseName string `mapstructure:"databaseName" json:"databaseName" yaml:"databaseName"`
}

type RedisConfig struct {
	AddrField
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

type EnvironmentEnum string

const (
	Development EnvironmentEnum = "Development"
	Test        EnvironmentEnum = "Test"
	Production  EnvironmentEnum = "Production"
)

type SystemConfig struct {
	AddrField                   // 服务器地址:端口
	Environment EnvironmentEnum `mapstructure:"environment" json:"environment" yaml:"environment" validate:"oneof=Development Test Production"`
	Debug       bool            `mapstructure:"debug" json:"debug" yaml:"debug"`
	Label       bool            `mapstructure:"label" json:"label" yaml:"label"`
	Description string          `mapstructure:"description" json:"description" yaml:"description"`
}

type HbaseConfig struct {
	AddrField
}

type KafkaConfig struct {
	AddrField
}

type JwtConfig struct {
	AuthHeaderPrefix string `mapstructure:"authHeaderPrefix" json:"authHeaderPrefix" yaml:"authHeaderPrefix"`
	Secret           string `mapstructure:"secret" json:"secret" yaml:"secret"`                      // 密钥
	ExpireMinute     int64  `mapstructure:"expireMinute" json:"expireMinute" yaml:"expireMinute"`    // 过期时间
	RefreshMinute    int64  `mapstructure:"refreshMinute" json:"refreshMinute" yaml:"refreshMinute"` // 缓冲时间
	Issuer           string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                      // 签发者
}

type Config struct {
	RelationalDatabase RelationalDatabaseConfig
	Redis              RedisConfig
	System             SystemConfig
	Hbase              HbaseConfig
	Kafka              KafkaConfig
	Jwt                JwtConfig
}
