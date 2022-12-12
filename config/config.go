package config

import (
	flag "github.com/spf13/pflag"
)

type Config struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	SSLMode         string
	LogstashHost    string
	LogstashPort    int
	LogstashTimeout int
	TickerTime      int
}

var (
	Conf    *Config
	VERSION string
)

func InitConfig() {
	Conf = new(Config)
	flag.StringVar(&Conf.Host, "host", "127.0.0.1", "postgresql host")
	flag.IntVar(&Conf.Port, "port", 5432, "postgresql port")
	flag.StringVar(&Conf.User, "user", "postgres", "postgresql user")
	flag.StringVar(&Conf.Password, "password", "1qaz@WSX", "postgresql password")
	flag.StringVar(&Conf.DBName, "dbname", "postgres", "postgresql dbname")
	flag.StringVar(&Conf.SSLMode, "sslmode", "disable", "postgresql sslmode")
	flag.StringVar(&Conf.LogstashHost, "logstash-host", "127.0.0.1", "logstash host")
	flag.IntVar(&Conf.LogstashPort, "logstash-port", 2514, "logstash port")
	flag.IntVar(&Conf.LogstashTimeout, "logstash-timeout", 60, "logstash timeout")
	flag.IntVar(&Conf.TickerTime, "ticker-time", 15, "ticker time seconds")
	flag.Parse()
}
