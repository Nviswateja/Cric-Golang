package main

type SeriesData struct {
	Apikey string `json:"apikey"`
	Data   []struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		StartDate string `json:"startDate,omitempty"`
		EndDate   string `json:"endDate"`
		Odi       int    `json:"odi"`
		T20       int    `json:"t20"`
		Test      int    `json:"test"`
		Squads    int    `json:"squads"`
		Matches   int    `json:"matches"`
	} `json:"data"`
	Status string `json:"status"`
	Info   struct {
		HitsToday  int     `json:"hitsToday"`
		HitsLimit  int     `json:"hitsLimit"`
		Credits    int     `json:"credits"`
		Server     int     `json:"server"`
		OffsetRows int     `json:"offsetRows"`
		Cache      int     `json:"cache"`
		TotalRows  int     `json:"totalRows"`
		QueryTime  float64 `json:"queryTime"`
		S          int     `json:"s"`
	} `json:"info"`
}
