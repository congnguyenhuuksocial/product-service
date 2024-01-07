package config

import (
	"go.uber.org/fx"
	"product-service/pkg"
	"product-service/pkg/constants"
)

type Config struct {
	Logger   *LoggerConfig
	Database *DatabaseConfig
	GRPC     *GRPCConfig
	Server   *ServerConfig
	Search   *SearchConfig
	Business *BusinessConfig
	GEO      *GeoConfig
	Redis    *Redis
	Kafka    *KafkaConfig
}

type Redis struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	URI  string `mapstructure:"uri"`
	TTL  int    `mapstructure:"ttl"`
}

type LoggerConfig struct {
	Level    string `mapstructure:"level"`
	Format   string `mapstructure:"format"`
	Prefix   string `mapstructure:"prefix"`
	LogLevel string `mapstructure:"log_level"`
	DevMode  bool   `mapstructure:"dev_mode"`
	Encoder  string `mapstructure:"encoder"`
}

type DatabaseConfig struct {
	Host                  string `mapstructure:"host"`
	Port                  int    `mapstructure:"port"`
	Username              string `mapstructure:"username"`
	Password              string `mapstructure:"password"`
	Database              string `mapstructure:"database"`
	ConnMaxLifetimeSecond int    `mapstructure:"conn_max_lifetime_second"`
	MaxOpenConns          int    `mapstructure:"max_open_conns"`
	MaxIdleConns          int    `mapstructure:"max_idle_conns"`
	InitErr               *InitErrConfig
	Driver                string `mapstructure:"driver"`
}

type InitErrConfig struct {
	MaxRetryNumber      int `mapstructure:"max_retry_number"`
	RetryIntervalSecond int `mapstructure:"retry_interval_second"`
}

type GRPCConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Env      string `mapstructure:"env"`
	UseRedis bool   `mapstructure:"use_redis"`
	Cdn      string `mapstructure:"cdn"`
}

type KafkaConfig struct {
	URI     string `mapstructure:"uri"`
	Topic   string `mapstructure:"topic"`
	GroupID string `mapstructure:"group"`
}
type BusinessConfig struct{}

type SearchConfig struct {
	Url          string `mapstructure:"url"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	MaxRetries   int    `mapstructure:"max_retries"`
	ProductIndex string `mapstructure:"product_index"`
}

type GeoConfig struct{}

func NewConfig() *Config {
	return &Config{
		Logger: &LoggerConfig{
			Level:   pkg.GetEnv(constants.EnvLogLevel, "debug"),
			Format:  pkg.GetEnv(constants.EnvLogFormat, "json"),
			Prefix:  pkg.GetEnv(constants.EnvLogPrefix, "product-service"),
			DevMode: pkg.GetEnvBool(constants.EnvLogDevMode, true),
			Encoder: pkg.GetEnv(constants.EnvLogEncoder, "console"),
		},
		Database: &DatabaseConfig{
			Host:                  pkg.GetEnv(constants.EnvDbHost, "localhost"),
			Port:                  pkg.GetEnvInt(constants.EnvDbPort, 3306),
			Username:              pkg.GetEnv(constants.EnvDbUsername, "product"),
			Password:              pkg.GetEnv(constants.EnvDbPassword, "product"),
			Database:              pkg.GetEnv(constants.EnvDbName, "product"),
			ConnMaxLifetimeSecond: pkg.GetEnvInt(constants.EnvDbConnMaxLifetimeSecond, 60),
			MaxOpenConns:          pkg.GetEnvInt(constants.EnvMaxOpenConns, 10),
			MaxIdleConns:          pkg.GetEnvInt(constants.EnvMaxIdleConns, 5),
			InitErr: &InitErrConfig{
				MaxRetryNumber:      pkg.GetEnvInt(constants.EnvDbInitMaxRetryNumber, 5),
				RetryIntervalSecond: pkg.GetEnvInt(constants.EnvDbInitRetryIntervalSecond, 5),
			},
			Driver: pkg.GetEnv(constants.EnvDbDriver, "mysql"),
		},
		GRPC: &GRPCConfig{
			Host: pkg.GetEnv(constants.EnvGrpcHost, "localhost"),
			Port: pkg.GetEnvInt(constants.EnvGrpcPort, 8092),
		},
		Server: &ServerConfig{
			Host:     pkg.GetEnv(constants.EnvServerHost, "localhost"),
			Port:     pkg.GetEnvInt(constants.EnvServerPort, 8080),
			Env:      pkg.GetEnv(constants.EnvServerEnv, "local"),
			UseRedis: false,
			Cdn:      pkg.GetEnv(constants.EnvServerCdn, "https://d363fblnfweia.cloudfront.net"),
		},
		Search: &SearchConfig{
			Url:          pkg.GetEnv(constants.EnvSearchUrl, "localhost:9200"),
			Username:     pkg.GetEnv(constants.EnvSearchUsername, "search"),
			Password:     pkg.GetEnv(constants.EnvSearchPassword, "search"),
			MaxRetries:   pkg.GetEnvInt(constants.EnvSearchMaxRetries, 5),
			ProductIndex: pkg.GetEnv(constants.EnvSearchFlavorIndex, "product"),
		},
		Redis: &Redis{
			Host: pkg.GetEnv(constants.EnvRedisHost, "localhost"),
			Port: pkg.GetEnvInt(constants.EnvRedisPort, 6379),
			URI:  pkg.GetEnv(constants.EnvRedisUri, "localhost:6379"),
			TTL:  pkg.GetEnvInt(constants.EnvRedisTTL, 5),
		},
		Kafka: &KafkaConfig{
			URI:     pkg.GetEnv(constants.EnvKafkaUri, "localhost:9092"),
			Topic:   pkg.GetEnv(constants.EnvKafkaTopic, "product.created"),
			GroupID: pkg.GetEnv(constants.EnvKafkaGroup, "product"),
		},
	}
}

var Module = fx.Provide(NewConfig)
