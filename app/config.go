package app

type NetatmoConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
}

type InfluxConfig struct {
	Address      string `yaml:"address"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	Database     string `yaml:"database"`
	Precision    string `yaml:"precision"`
	MetricPrefix string `yaml:"metric_prefix"`
}

type AppConfig struct {
	Schedule string        `yaml:"schedule"`
	Netatmo  NetatmoConfig `yaml:"netatmo"`
	Influx   InfluxConfig  `yaml:"influx"`
}

var Config *AppConfig
var NetatmoCh = make(chan string, 10)
var InfluxCh = make(chan NetatmoValues, 10)
