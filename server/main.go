package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"practice/weather-server/dto"
	"strings"
)

type Response struct {
	Cod     int       `json:"cod"`
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
	Wind    Wind      `json:"wind"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Humidity int     `json:"humidity"`
}

type Wind struct {
	Speed float32 `json:"speed"`
}

const BASE_URL = "https://api.openweathermap.org/data/2.5/weather"

func handler(w http.ResponseWriter, r *http.Request) {

	apiKey := os.Getenv("API_KEY")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != "GET" {
		fmt.Fprintf(w, "Method not support")
		return
	}

	var apiRes dto.Response[*dto.WeatherDto]

	response, err := http.Get(BASE_URL + "?q=" + strings.ReplaceAll(string(r.URL.Query().Get("city")), " ", "%20") + "&appid=" + apiKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		apiRes.StatusCode = "FAIL"
		b, err := json.Marshal(&apiRes)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%s", string(b))
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		apiRes.StatusCode = "FAIL"
		b, err := json.Marshal(&apiRes)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "%s", string(b))
		log.Fatal(err)
	}

	var res Response
	err = json.Unmarshal(responseData, &res)
	if err != nil {
		fmt.Println(err)
	}
	if res.Cod != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		apiRes.StatusCode = "FAIL"
		apiRes.Data = nil
		b, err := json.Marshal(&apiRes)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, "%s", string(b))
		return
	}

	fmt.Println("===================")
	fmt.Println(res.Cod)
	fmt.Println("===================")

	var dto dto.WeatherDto

	dto.Tempotary = res.Main.Temp
	dto.Description = res.Weather[0].Description
	dto.Humidity = res.Main.Humidity
	dto.Main = res.Weather[0].Main
	dto.WindSpeed = res.Wind.Speed

	apiRes.StatusCode = "OK"
	apiRes.Data = &dto

	b, err := json.Marshal(&apiRes)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(w, "%s", string(b))
}

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API_KEY environment variable not set.")
		os.Exit(1)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
