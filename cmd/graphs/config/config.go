package config

import (
	"github.com/spf13/viper"
)

var Config appConfig

type appConfig struct {
	RedisHost string `mapstructure:"redis_host"`
}

func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("server")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("blueprint")
	v.AutomaticEnv()

	//Config.DSN = v.Get("DSN").(string)
	//Config.ApiKey = v.Get("API_KEY").(string)
	v.SetDefault("redis_host", "redis:6379")

	//for _, path := range configPaths {
	//	v.AddConfigPath(path)
	//}
	//if err := v.ReadInConfig(); err != nil {
	//	return fmt.Errorf("failed to read the configuration file: %s", err)
	//}
	return v.Unmarshal(&Config)
}
