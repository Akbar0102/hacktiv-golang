package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	go autohit()
	select {}
}

func autohit() {
	for {
		min := 1
		max := 100
		rand.Seed(time.Now().UnixNano())
		numbrandwater := rand.Intn(max - min)
		numbrandwind := rand.Intn(max - min)

		jsondata := map[string]interface{}{
			"Water": numbrandwater,
			"Wind":  numbrandwind,
		}

		requestJson, err := json.Marshal(jsondata)

		client := &http.Client{}
		if err != nil {
			log.Fatalln(err)
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
		req.Header.Set("Content-type", "application/json")

		if err != nil {
			log.Fatalln(err)
		}

		res, err := client.Do(req)

		if err != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Fatalln(err)
		}

		var responseData map[string]interface{}
		err = json.Unmarshal(body, &responseData)
		if err != nil {
			log.Fatalln(err)
		}

		// hapus field "id" dari map jika ada
		delete(responseData, "id")

		// serialize kembali ke JSON
		responseJson, err := json.Marshal(responseData)
		if err != nil {
			log.Fatalln(err)
		}

		// log.Println(string(body) + "\n")
		fmt.Println(string(responseJson))

		jsonprint, err := json.MarshalIndent(jsondata, "", "    ")
		if err != nil {
			fmt.Println("json print error")
			return
		}

		_ = jsonprint
		switch numbrandwater > 0 {
		case numbrandwater <= 5:
			result := "aman"
			fmt.Printf("status water : %s \n", result)
		case numbrandwater >= 6 && numbrandwater <= 8:
			result := "siaga"
			fmt.Printf("status water : %s \n", result)
		case numbrandwater > 8:
			result := "bahaya"
			fmt.Printf("status water : %s \n", result)
		default:
			result := "water measurable error"
			fmt.Println(result)
		}

		switch numbrandwind > 0 {
		case numbrandwind <= 6:
			result := "aman"
			fmt.Printf("status wind : %s \n", result)
		case numbrandwind >= 7 && numbrandwind <= 15:
			result := "siaga"
			fmt.Printf("status wind : %s \n", result)
		case numbrandwind > 15:
			result := "bahaya"
			fmt.Printf("status wind : %s \n", result)
		default:
			result := "wind measurable error"
			fmt.Println(result)
		}
		fmt.Printf("\n")
		time.Sleep(time.Second * 15)
	}
}
