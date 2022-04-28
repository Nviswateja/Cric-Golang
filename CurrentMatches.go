package main

type CurrentMatchList struct {
	Apikey string `json:"apikey"`
	Data   []struct {
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		MatchType   string   `json:"matchType"`
		Status      string   `json:"status"`
		Venue       string   `json:"venue"`
		Date        string   `json:"date"`
		DateTimeGMT string   `json:"dateTimeGMT"`
		Teams       []string `json:"teams"`
		Score       []struct {
			R      int     `json:"r"`
			W      int     `json:"w"`
			O      float64 `json:"o"`
			Inning string  `json:"inning"`
		} `json:"score"`
		SeriesID       string `json:"series_id"`
		FantasyEnabled bool   `json:"fantasyEnabled"`
		HasSquad       bool   `json:"hasSquad"`
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