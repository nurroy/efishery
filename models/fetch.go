package models


type Storage struct {
	UUID         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         float64 `json:"size"`
	Price        float64 `json:"price"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
	PriceUSD     float64 `json:"price_usd"`
}

type ResponseStorage struct {
	RespStorage []Storage
}

type AggregateValue struct {
	Min		float64	`json:"min"`
	Max		float64	`json:"max"`
	Median	float64	`json:"median"`
	Avg		float64	`json:"avg"`
}

type AggregateData struct {
	Price AggregateValue `json:"price"`
	Size  AggregateValue `json:"size"`
	PricfeUSD AggregateValue `json:"price_usd"`
}

type RespData struct {
	AreaProvinsi string  `json:"area_provinsi"`
	Week 		int32		`json:"week"`
	Aggregate    AggregateData `json:"aggregate"`
}

type ResponseAgregate struct {
	Data []RespData `json:"data"`
}


