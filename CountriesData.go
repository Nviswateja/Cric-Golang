package main

type CountriesData struct {
	Apikey string `json:"apikey"`
	Data   []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		GenericFlag string `json:"genericFlag"`
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
