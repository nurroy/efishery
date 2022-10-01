package efishery

import (
	"belajar/efishery/pkg"
	"belajar/efishery/utils/constants"
	"encoding/json"
	"log"
	"time"
)

type ResponseEfishery []struct {
	UUID         string `json:"uuid,omitempty"`
	Komoditas    string `json:"komoditas,omitempty"`
	AreaProvinsi string `json:"area_provinsi,omitempty"`
	AreaKota     string `json:"area_kota,omitempty"`
	Size         string `json:"size,omitempty"`
	Price        string `json:"price,omitempty"`
	TglParsed    string `json:"tgl_parsed,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
}


// PostEfisheryStorage is request to api
func PostEfisheryStorage() (ResponseEfishery, error) {
	res, err := pkg.Invoke(&pkg.ReqInfo{
		URL:    "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list",
		Method: constants.MethodGET,
		HeadersInfo: map[string]interface{}{
			constants.HeaderContentType: constants.JSONType,
		},
	}, 60*time.Second)
	if err != nil {
		return nil, err
	}

	log.Println(string(res.Body))
	var data ResponseEfishery

	err = json.Unmarshal(res.Body, &data)
	if err != nil {
		return nil, err
	}


	return data, nil
}

