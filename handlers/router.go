package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/temperature/:city", func(ctx *gin.Context) {
		city := ctx.Param("city")
		data, err := GetWeatherData(city)
		if err != nil {
			ctx.JSON(200, nil)
		}
		ctx.JSON(200, data)
	})
}

type OpenWeatherMapData struct {
	CityName string `json:"name"`
	Current  struct {
		Temp     float64 `json:"temp"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
		Humidity float64 `json:"humidity"`
	} `json:"main"`
	Sys struct {
		Country string `json:"country"`
	} `json:"sys"`
}

func GetWeatherData(cityId string) (OpenWeatherMapData, error) {
	data := OpenWeatherMapData{}

	API := "api_ma_ban_lay_tu_openweathermap"
	res, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + cityId + "&units=metric&appid=" + API)
	if err != nil || res.StatusCode != 200 {
		return data, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
