// nolint
package injection

import (
	_ "time/tzdata"

	_ "github.com/joho/godotenv/autoload"

	"github.com/andys920605/hr-system/pkg/conf"
	"github.com/andys920605/hr-system/pkg/logging"
)

type Injection struct {
	Config *conf.Config
	Logger *logging.Logging
}

func New() *Injection {
	config := initConfig()
	logger := initLogger(config)

	return &Injection{
		Config: config,
		Logger: logger,
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

// func initRedisClusterClient(config *conf.Config, logger *logging.Logging) *redis.ClusterClient {
// 	client := redis.NewClusterClient(&redis.ClusterOptions{
// 		Addrs: []string{config.Google.Redis.Cluster.Addr},
// 	})
// 	if _, err := client.Ping(context.Background()).Result(); err != nil {
// 		logger.Emergencyf("failed to initialize redis cluster client: %v", err)
// 	}
// 	logger.Infof("redis cluster client initialized")
// 	return client
// }
