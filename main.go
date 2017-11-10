package main

import (
	"./app"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
)

func init() {
	configFile := flag.String("config", "config/config.yml", "configuration file")
	flag.Parse()
	file, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(err.Error())
	}
	yaml.Unmarshal(file, &app.Config)
}

func main() {
	go app.Scheduler()
	go func() {
		for range app.NetatmoCh {
			values, err := app.GetNetatmoValues()
			if err != nil {
				log.Println(err.Error())
			} else {
				app.InfluxCh <- values
			}
		}
	}()
	go func() {
		for points := range app.InfluxCh {
			err := app.SendToInflux(points)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
