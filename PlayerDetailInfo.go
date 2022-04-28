package main

type PlayerDetailInfo struct {
	Apikey string `json:"apikey"`
	Data   struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		DateOfBirth  string `json:"dateOfBirth"`
		Role         string `json:"role"`
		BattingStyle string `json:"battingStyle"`
		BowlingStyle string `json:"bowlingStyle"`
		PlaceOfBirth string `json:"placeOfBirth"`
		Country      string `json:"country"`
		PlayerImg    string `json:"playerImg"`
		Stats        []struct {
			Fn        string `json:"fn"`
			Matchtype string `json:"matchtype"`
			Stat      string `json:"stat"`
			Value     string `json:"value"`
		} `json:"stats"`
	} `json:"data"`
	Status string `json:"status"`
	Info   struct {
		HitsToday int     `json:"hitsToday"`
		HitsLimit int     `json:"hitsLimit"`
		Credits   int     `json:"credits"`
		Server    int     `json:"server"`
		QueryTime float64 `json:"queryTime"`
		S         int     `json:"s"`
	} `json:"info"`
}