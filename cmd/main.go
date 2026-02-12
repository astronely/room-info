package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/astronely/room-info/internal/api"
	"github.com/astronely/room-info/internal/config"
	"github.com/astronely/room-info/internal/influx"
)

func main() {
	config.Load()

	httpConfig := config.NewHttpServerConfig()
	influxConfig := config.NewInfluxConfig()

	influx.Init(influxConfig)

	http.HandleFunc("/metrics", api.SensorHandler)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		log.Printf("Сервер успешно запущен по адресу: %v!\n", httpConfig.Addr())
		if err := http.ListenAndServe(httpConfig.Addr(), nil); err != nil {
			log.Fatalf("Server was closed with error: %v", err)
		}
	}()

	wg.Wait()
}
