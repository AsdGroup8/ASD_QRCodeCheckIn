package conf

import "github.com/spf13/cobra"

var (
	Addr      string
	LogFile   string
	LogLevel  string
	DBConnStr string
	IsDebug   bool
)

func Init(cmd *cobra.Command) {
	flags := cmd.PersistentFlags()
	flags.StringVar(&Addr, "addr", "127.0.0.1:8080", "http server listen address")
	flags.StringVar(&LogFile, "logfile", "stdout", "log output file. stdout|stderr|file")
	flags.StringVar(&LogLevel, "loglevel", "debug", "log output level")
	flags.StringVar(&DBConnStr, "db", "localhost:3306", "database connection string")
	flags.BoolVar(&IsDebug, "debug", true, "if debug")
}
