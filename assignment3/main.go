package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Sensor struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func main() {
	// jalankan service dengan goroutine
	loopService()
}

func loopService() {
	ticker := time.NewTicker(15 * time.Second)
	done := make(chan bool)

	go sendRandomData(ticker, done)

	// 1 menit berhenti
	time.Sleep(1 * time.Minute)
	done <- true
	fmt.Println("stop")
}

func sendRandomData(ticker *time.Ticker, done chan bool) {
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			data := Sensor{
				Water: generateNumber(),
				Wind:  generateNumber(),
			}

			// kirim data ke API
			status, err := sendDataToAPI(data)
			if err != nil {
				fmt.Println("Failed to send data to API:", err)
			} else {
				fmt.Println(status)
			}
		}
	}
}

func sendDataToAPI(data Sensor) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return "", err
	}

	fmt.Println(string(jsonData))

	req, err := http.NewRequest("POST", "http://localhost:8081/sensor", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	// set header
	req.Header.Set("Content-Type", "application/json")

	// buat client http
	client := &http.Client{}

	// kirim request
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	respData := struct {
		WaterStatus string `json:"water_status"`
		WindStatus  string `json:"wind_status"`
	}{}

	err = json.NewDecoder(res.Body).Decode(&respData)
	if err != nil {
		return "", err
	}

	resultStatus := fmt.Sprintf("status water: %s\nstatus wind: %s\n", respData.WaterStatus, respData.WindStatus)
	return resultStatus, nil
}

func generateNumber() int {
	return rand.Intn(100) + 1
}
