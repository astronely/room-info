package config

import (
	"net"

	"github.com/spf13/viper"
)

func Load() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

type InfluxDBConfig struct {
	Addr   string
	Token  string
	Org    string
	Bucket string
}

func NewInfluxConfig() *InfluxDBConfig {
	return &InfluxDBConfig{
		Addr:   viper.GetString("INFLUX_ADDR"),
		Token:  viper.GetString("INFLUX_TOKEN"),
		Org:    viper.GetString("INFLUX_ORG"),
		Bucket: viper.GetString("INFLUX_BUCKET"),
	}
}

type HttpServerConfig struct {
	Host string
	Port string
}

func NewHttpServerConfig() *HttpServerConfig {
	return &HttpServerConfig{
		Host: viper.GetString("HTTP_HOST"),
		Port: viper.GetString("HTTP_PORT"),
	}
}

func (c *HttpServerConfig) Addr() string {
	return net.JoinHostPort(c.Host, c.Port)
}
