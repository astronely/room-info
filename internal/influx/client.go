package influx

import (
	"context"
	"time"

	"github.com/astronely/room-info/internal/config"
	"github.com/astronely/room-info/internal/entity"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

var client influxdb2.Client
var writeAPI api.WriteAPIBlocking

func Init(cfg *config.InfluxDBConfig) {
	client = influxdb2.NewClient(
		cfg.Addr,
		cfg.Token,
	)

	writeAPI = client.WriteAPIBlocking(cfg.Org, cfg.Bucket)
}

func WriteSensorData(data entity.SensorData) {
	p := influxdb2.NewPoint(
		"env",
		map[string]string{
			"device": data.DeviceID,
		},
		map[string]any{
			"temperature": data.Temperature,
			"humidity":    data.Humidity,
			"pressure":    data.Pressure,
			"light":       data.Light,
		},
		time.Unix(data.Timestamp, 0),
	)

	writeAPI.WritePoint(context.Background(), p)
}
