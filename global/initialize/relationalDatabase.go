package initialize

import (
	"github.com/hail-pas/GinStartKit/global"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GormDB() {
	relationalDatabaseConfig := global.Configuration.RelationalDatabase
	if relationalDatabaseConfig.DatabaseName == "" {
		panic("Empty DatabaseName")
	}

	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	var db *gorm.DB
	var err error

	if relationalDatabaseConfig.Type == "mysql" {
		mysqlConfig := mysql.Config{
			DSN:                       relationalDatabaseConfig.Dsn(), // DSN data source name
			DefaultStringSize:         256,                            // string 类型字段的默认长度
			SkipInitializeWithVersion: false,                          // 根据版本自动配置
		}
		if db, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
			panic(err)
		}
	} else {
		pgsqlConfig := postgres.Config{
			DSN:                  relationalDatabaseConfig.Dsn(), // DSN data source name
			PreferSimpleProtocol: true,
		}
		if db, err = gorm.Open(postgres.New(pgsqlConfig), gormConfig); err != nil {
			panic(err)
		}
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(relationalDatabaseConfig.GORM.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(relationalDatabaseConfig.GORM.MaxOpenConnections)
	if err = sqlDB.Ping(); err != nil {
		panic(err)
	}
	global.RelationalDatabase = db
}
