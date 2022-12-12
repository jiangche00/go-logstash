package main

import (
	"encoding/json"
	"fmt"
	"go-logstash/config"
	"go-logstash/pkg/logstashtcp"
	pg "go-logstash/pkg/postgresql"
	"log"
	"time"
)

func main() {
	config.InitConfig()
	ticker := time.NewTicker(time.Duration(config.Conf.TickerTime) * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				l := logstashtcp.New(config.Conf.LogstashHost, config.Conf.LogstashPort, config.Conf.LogstashTimeout)
				_, err := l.Connect()
				if err != nil {
					fmt.Println(err)
				}
				bs, _ := json.Marshal(pg.GetSysInfo())
				log.Println(string(bs))
				err = l.Writeln(string(bs))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}()
	select {}
}
