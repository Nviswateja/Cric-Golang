package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	None = iota
	SeriesList
	CountryList
	CurrentMatches
	PlayersSearch
)

func GetData(input int, playerName string) {
	switch input {
	case SeriesList:
		GetSeries()
	case CountryList:
		GetCountryList()
	case CurrentMatches:
		GetCurrentMatches()
	case PlayersSearch:
		SearchYourPlayer(playerName)
	}
}

func SearchYourPlayer(playerName string) {
	res := fmt.Sprintf("https://api.cricapi.com/v1/players?apikey=27f64851-0fb6-48cd-969f-9736508a1fcc&offset=0&search=%s", playerName)
	resp1, err := http.Get(res)
	if err != nil {
		panic(err)
	}
	defer resp1.Body.Close()
	body1 := ReadData(resp1)

	s1 := string(body1)

	data := PlayerInfo{}
	json.Unmarshal([]byte(s1), &data)
	url2 := fmt.Sprintf("https://api.cricapi.com/v1/players_info?apikey=27f64851-0fb6-48cd-969f-9736508a1fcc&offset=0&id=%s", data.Data[0].ID)
	resp2, err := http.Get(url2)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()
	body2 := ReadData(resp2)

	s2 := string(body2)
	data2 := PlayerDetailInfo{}
	json.Unmarshal([]byte(s2), &data2)
	fmt.Printf("Name = %s\n", data2.Data.Name)
	fmt.Printf("Batting style = %s\n", data2.Data.BattingStyle)
	fmt.Printf("DOB = %s\n", data2.Data.DateOfBirth)
}

func GetCurrentMatches() {
	resp, err := http.Get("https://api.cricapi.com/v1/currentMatches?apikey=27f64851-0fb6-48cd-969f-9736508a1fcc&offset=0")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body := ReadData(resp)

	s := string(body)
	data := CurrentMatchList{}
	json.Unmarshal([]byte(s), &data)
	for _, service := range data.Data {
		fmt.Printf("%#v\n", service.Name)
	}
}

func GetCountryList() {
	resp, err := http.Get("https://api.cricapi.com/v1/countries?apikey=27f64851-0fb6-48cd-969f-9736508a1fcc&offset=0")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body := ReadData(resp)

	s := string(body)
	if s == "" {
		fmt.Println("error while fetching data")
	}
	data := SeriesData{}
	json.Unmarshal([]byte(s), &data)
	for _, service := range data.Data {
		fmt.Printf("%#v\n", service.Name)
	}
}

func GetSeries() {
	resp, err := http.Get("https://api.cricapi.com/v1/series?apikey=27f64851-0fb6-48cd-969f-9736508a1fcc&offset=0")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body := ReadData(resp)

	s := string(body)
	if s == "" {
		fmt.Println("error while fetching data")
	}
	data := SeriesData{}
	json.Unmarshal([]byte(s), &data)
	for _, service := range data.Data {
		fmt.Printf("%#v\n", service.Name)
	}
}

func ReadData(resp *http.Response) []byte {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}
