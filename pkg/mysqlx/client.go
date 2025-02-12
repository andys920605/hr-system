package mysqlx

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/andys920605/hr-system/pkg/conf"
	"github.com/andys920605/hr-system/pkg/errors"
	"github.com/andys920605/hr-system/pkg/migration"
)

type Client struct {
	*gorm.DB
}

func NewClient(config *conf.Config) (*Client, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&readTimeout=30s&writeTimeout=30s",
		config.MySQL.Username, config.MySQL.Password, config.MySQL.Host, config.MySQL.Port, config.MySQL.Database)

	gormDB, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	if config.MySQL.MaxIdle == 0 || config.MySQL.MaxOpen == 0 {
		return nil, errors.New("missing maxIdle or maxOpen")
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get database instance")
	}

	sqlDB.SetMaxIdleConns(config.MySQL.MaxIdle)
	sqlDB.SetMaxOpenConns(config.MySQL.MaxOpen)

	if err = sqlDB.Ping(); err != nil {
		return nil, errors.Wrapf(err, "error pinging database")
	}

	if err := migration.AutoMigrate(gormDB); err != nil {
		return nil, errors.Wrap(err, "auto migration")
	}

	return &Client{gormDB}, nil
}
