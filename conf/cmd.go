package conf

import (
	"fmt"
	"log"

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
	IsDebug   bool
)

// Init initialize config
func Init(cmd *cobra.Command) {

	flags := cmd.PersistentFlags()
	flags.StringVar(&Env, "env", "", "environment")
	if Env == "" {
		log.Fatal("--env must be specified. dev|local|qas|prd")
	}
	flags.Int32Var(&ServerID, "serverid", 0, "server id")
	if ServerID == 0 {
		log.Fatal("--serverid must be > 0")
	}
	if err := InitConfigFile(); err != nil {
		log.Panic(err)
	}
	flags.StringVar(&Addr, "addr", envConf.Addr, "http server listen address")
	flags.StringVar(&LogFile, "logfile", envConf.LogFile, "log output file. stdout|stderr|file")
	flags.StringVar(&LogLevel, "loglevel", envConf.LogLevel, "log output level")
	dbConnStr := fmt.Sprintf("%s:%s@%s/%s_%d?parseTime=true&loc=Local", envConf.DBUser,
		envConf.DBPasswd, envConf.DBUrl, envConf.DBName, ServerID)
	flags.StringVar(&DBConnStr, "db", dbConnStr, "database connection string")
	flags.BoolVar(&IsDebug, "debug", true, "if debug")
}
