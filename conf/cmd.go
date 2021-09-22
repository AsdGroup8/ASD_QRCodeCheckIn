package conf

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Config
var (
	Env      string
	ServerID int32

	Addr      string
	LogFile   string
	LogLevel  string
	DBConnStr string
	DBName    string
	IsDebug   bool
	Secret    []byte
)

// InitFlags initialize command line args
func InitFlags(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()
	flags.StringVar(&Env, "env", "", "server environment. e.g. dev|local|qas|prd")
	flags.Int32Var(&ServerID, "serverid", 0, "server id")
}

// Init initialize config
func Init(cmd *cobra.Command) error {
	if err := initConfigFile(); err != nil {
		return err
	}
	Secret = []byte(envConf.Secret)
	flags := cmd.PersistentFlags()
	flags.StringVar(&Addr, "addr", envConf.Addr, "http server listen address")
	flags.StringVar(&LogFile, "logfile", envConf.LogFile, "log output file. stdout|stderr|file")
	flags.StringVar(&LogLevel, "loglevel", envConf.LogLevel, "log output level")
	dbConnStr := fmt.Sprintf("%s:%s@tcp(%s)/?parseTime=true&loc=Local", envConf.DBUser,
		envConf.DBPasswd, envConf.DBUrl)
	DBName = fmt.Sprintf("%s_%s_%d", envConf.DBName, Env, ServerID)
	flags.StringVar(&DBConnStr, "db", dbConnStr, "database connection string")
	flags.BoolVar(&IsDebug, "debug", true, "if debug")
	return nil
}
