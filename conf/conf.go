package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

var envConf EnvConfig

// EnvConfig ...
type EnvConfig struct {
	Addr     string `mapstructure:"addr"`
	LogFile  string `mapstructure:"logfile"`
	LogLevel string `mapstructure:"loglevel"`
	DBUrl    string `mapstructure:"db_url"`
	DBUser   string `mapstructure:"db_user"`
	DBPasswd string `mapstructure:"db_password"`
	DBName   string `mapstructure:"db_name"`
	Secret   string `mapstructure:"secret"`
}

// initConfigFile ...
func initConfigFile() error {
	if Env == "" {
		return fmt.Errorf("invalid server environment. %s", Env)
	}
	if ServerID == 0 {
		return fmt.Errorf("server id must > 0. but receive %d", ServerID)
	}
	viper.AddConfigPath("./conf")
	viper.SetConfigName("conf_" + Env)
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("fail to read in config. %v", err)
	}
	if err := viper.Unmarshal(&envConf); err != nil {
		return fmt.Errorf("fail to unmarshal config file. %v", err)
	}
	return nil
}
