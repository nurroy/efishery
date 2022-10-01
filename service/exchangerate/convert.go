package exchangerate

import (
	"belajar/efishery/pkg"
	"belajar/efishery/utils/constants"
	"encoding/json"
	"time"
)


type ResQuery struct {
	Amount int    `json:"amount"`
	From   string `json:"from"`
	To     string `json:"to"`
}

type ResInfo struct {
	Rate      float64 `json:"rate"`
	Timestamp int     `json:"timestamp"`
}

type ResConverter struct {
	Date string `json:"date"`
	Info ResInfo  `json:"info"`
	Query ResQuery`json:"query"`
	Result  float64 `json:"result"`
	Success bool    `json:"success"`
}


// PostConvert is request to api
func PostConvert() (*ResConverter, error) {
	res, err := pkg.Invoke(&pkg.ReqInfo{
		URL:    "https://api.apilayer.com/exchangerates_data/convert?to=USD&from=IDR&amount=10000",
		Method: constants.MethodGET,
		HeadersInfo: map[string]interface{}{
			constants.HeaderContentType: constants.JSONType,
			constants.HeaderApiKey: "MkOIS0PeoWldhGDtRNGXKVbzr34GRQcI",
		},
	}, 60*time.Second)
	if err != nil {
		return nil, err
	}

	var resp ResConverter

	err = json.Unmarshal(res.Body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}


