package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/astronely/room-info/internal/entity"
	"github.com/astronely/room-info/internal/influx"
)

func SensorHandler(w http.ResponseWriter, r *http.Request) {
	var data entity.SensorData
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if data.Timestamp == 0 {
		data.Timestamp = time.Now().Unix()
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Data received: %+v\n", data)

	influx.WriteSensorData(data)

	w.WriteHeader(http.StatusAccepted)
}
