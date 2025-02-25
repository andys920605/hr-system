// nolint
package injection

import (
	"context"
	_ "time/tzdata"

	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"

	domain_service "github.com/andys920605/hr-system/internal/domain/service"
	"github.com/andys920605/hr-system/internal/north/local/appservice"
	employee_dao "github.com/andys920605/hr-system/internal/south/adapter/repository/dao/employee/mysql"
	employee_rep "github.com/andys920605/hr-system/internal/south/adapter/repository/employee"
	"github.com/andys920605/hr-system/pkg/conf"
	"github.com/andys920605/hr-system/pkg/logging"
	"github.com/andys920605/hr-system/pkg/migration"
	"github.com/andys920605/hr-system/pkg/mysqlx"
	"github.com/andys920605/hr-system/pkg/snowflake"
)

type Injection struct {
	Config             *conf.Config
	Logger             *logging.Logging
	EmployeeAppService *appservice.EmployeeAppService
}

func New() *Injection {
	config := initConfig()
	logger := initLogger(config)

	snowflake.Init(logger)

	mysqlxClient := initMysqlClient(config, logger)
	migrate(mysqlxClient, config, logger)

	redisClient := initRedisClusterClient(config, logger)

	employeeDao := employee_dao.NewEmployeeDao(mysqlxClient)
	employeeRep := employee_rep.NewEmployeeRepository(employeeDao, redisClient)
	employeeDomainSvc := domain_service.NewEmployeeDomainService(logger, employeeRep)
	employeeAppSvc := appservice.NewEmployeeAppService(logger, employeeDomainSvc)

	return &Injection{
		Config:             config,
		Logger:             logger,
		EmployeeAppService: employeeAppSvc,
	}
}

func initLogger(config *conf.Config) *logging.Logging {
	loggingLevel, err := logging.ParserLevel(config.Log.Level)
	if err != nil {
		panic(err)
	}

	logger := logging.New(
		logging.WithServiceName(config.Server.Name),
		logging.WithLevel(loggingLevel),
		logging.WithShowCaller(),
	)
	return logger
}

func initConfig() *conf.Config {
	config, err := conf.NewConfig()
	if err != nil {
		panic(err)
	}
	return config
}

func initMysqlClient(config *conf.Config, logger *logging.Logging) *mysqlx.Client {
	client, err := mysqlx.NewClient(config)
	if err != nil {
		logger.Emergencyf("failed to initialize mysql client: %v", err)
	}
	logger.Infof("mysql client initialized")
	return client
}

func initRedisClusterClient(config *conf.Config, logger *logging.Logging) *redis.ClusterClient {
	logger.Info(config.Redis.Cluster.Addr + ":" + config.Redis.Cluster.Port)
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{config.Redis.Cluster.Addr + ":" + config.Redis.Cluster.Port},
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		logger.Emergencyf("failed to initialize redis cluster client: %v", err)
	}
	logger.Infof("redis cluster client initialized")
	return client
}

func migrate(client *mysqlx.Client, config *conf.Config, logger *logging.Logging) {
	if err := migration.AutoMigrate(client.DB); err != nil {
		logger.Emergencyf("failed to auto migrate")
	}
	if config.Seed.Data.Enabled {
		migration.SeedData(client.DB, logger)
	}

	return
}
