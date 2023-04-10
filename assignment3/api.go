package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Data struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type ApiResponse struct {
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
}

func main() {
	r := gin.Default()

	// konfigurasi koneksi ke database PostgreSQL
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=pascal99 dbname=h8assignment sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close()

	// endpoint untuk menerima data melalui POST request
	r.POST("/sensor", func(c *gin.Context) {
		var data Data
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		waterStatus, windStatus := data.checkStatus()

		// simpan data ke dalam database
		sqlStatement := `INSERT INTO tb_m_sensor (water, wind) VALUES ($1, $2)`
		_, err = db.Exec(sqlStatement, data.Water, data.Wind)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		response := ApiResponse{
			WaterStatus: waterStatus,
			WindStatus:  windStatus,
		}

		c.JSON(http.StatusOK, response)
	})

	// jalankan server
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func (s Data) checkStatus() (water string, wind string) {
	if s.Water <= 5 {
		water = "aman"
	} else if s.Water >= 6 && s.Water <= 8 {
		water = "siaga"
	} else {
		water = "bahaya"
	}

	if s.Wind <= 6 {
		wind = "aman"
	} else if s.Wind >= 7 && s.Wind <= 15 {
		wind = "siaga"
	} else {
		wind = "bahaya"
	}

	return
}
