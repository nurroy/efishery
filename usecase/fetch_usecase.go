package usecase

import (
	"belajar/efishery/configs"
	"belajar/efishery/models"
	"belajar/efishery/service/efishery"
	"belajar/efishery/service/exchangerate"
	"belajar/efishery/utils/converter"
	"github.com/jellydator/ttlcache/v2"
	"log"
	"sort"
	"strconv"
	"time"
)

type FetchService struct {
	config *configs.Config
	cache ttlcache.SimpleCache
}

type Key struct {
	Province string
	Week int
}

type FetchServiceInterface interface {
	GetStorages()(models.ResponseStorage,error)
	GetAggregate()(*models.ResponseAgregate,error)
}

func CreateFetchServiceImpl(config *configs.Config,cache ttlcache.SimpleCache) FetchServiceInterface {
	return &FetchService{config,cache}
}


//GetStorages is storage usecase
func (r *FetchService) GetStorages() (resp models.ResponseStorage,err error) {
	var row models.Storage
	var rate float64
	res,err := efishery.PostEfisheryStorage()
	if err != nil{
		log.Printf("[Usecase][FETCH][STORAGES][1001] error when validate user with error : %v",err)
		return resp, err
	}

	//get rate
	rate, err = getRate(r,rate)
	if err != nil {
		log.Printf("[Usecase][FETCH][STORAGES][1002] error when validate user with error : %v",err)
		return resp, err
	}

	//appending data to response
	var rows []models.Storage
	for _, val := range res {
		price, _ := strconv.Atoi(val.Price)
		size, _ := strconv.Atoi(val.Size)

		row = models.Storage{
			UUID:         val.UUID,
			Komoditas:    val.Komoditas,
			AreaProvinsi: val.AreaProvinsi,
			AreaKota:     val.AreaKota,
			Size:         float64(size),
			Price:        float64(price),
			TglParsed:    val.TglParsed,
			Timestamp:    val.Timestamp,
			PriceUSD:     float64(price) * rate,
		}
		rows = append(rows, row)
	}

	resp = models.ResponseStorage{RespStorage: rows}
	r.cache.Set("rate", rate)
	return resp, nil
}

//GetAggregate is Aggregate usecase
func (r *FetchService) GetAggregate() (resp *models.ResponseAgregate,err error) {
	var rate float64
	res,err := efishery.PostEfisheryStorage()
	if err != nil{
		log.Printf("[Usecase][FETCH][AGGREGATE][1001] error when validate user with error : %v",err)
		return resp, err
	}

	//get rate
	rate, err = getRate(r,rate)
	if err != nil {
		log.Printf("[Usecase][FETCH][AGGREGATE][1002] error when validate user with error : %v",err)
		return resp, err
	}

	// group response data to map
	data := groupResp(res)

	//appending data for response
	arrData := appendDataResponse(data,rate)

	// sort arrData
	sort.Slice(arrData, func(i, j int) bool {
		var sortByProvince, sortByWeek bool
		sortByProvince = arrData[i].AreaProvinsi < arrData[j].AreaProvinsi

		if arrData[i].AreaProvinsi == arrData[j].AreaProvinsi {
			sortByWeek = arrData[i].Week < arrData[j].Week
			return sortByWeek
		}
		return sortByProvince
	})

	resp = &models.ResponseAgregate{Data: arrData}
	r.cache.Set("rate", rate)
	return resp, nil
}

//getRate function for get rate
func getRate(r *FetchService, rate float64) (float64, error) {
	v, err := r.cache.Get("rate")
	if err == ttlcache.ErrNotFound {
		//get converter
		resConv, err := exchangerate.PostConvert()
		if err != nil {
			return 0, err
		}
		rate = resConv.Info.Rate
	}

	if v != nil {
		rate = converter.InterfaceToFloat64(v)
	}
	return rate, nil
}


//groupResp for make a response group data to map
func groupResp(res efishery.ResponseEfishery) map[Key][]models.Storage {
	data := make(map[Key][]models.Storage)
	var row models.Storage
	for _, val := range res {
		iTime, _ := strconv.Atoi(val.Timestamp)
		tn := time.Unix(int64(iTime), 0)
		price, _ := strconv.Atoi(val.Price)
		size, _ := strconv.Atoi(val.Size)
		_, week := tn.ISOWeek()
		k := Key{
			Province: val.AreaProvinsi,
			Week:     week,
		}
		row = models.Storage{
			UUID:         val.UUID,
			Komoditas:    val.Komoditas,
			AreaProvinsi: val.AreaProvinsi,
			AreaKota:     val.AreaKota,
			Size:         float64(size),
			Price:        float64(price),
			TglParsed:    val.TglParsed,
			Timestamp:    val.Timestamp,
		}
		if val.AreaProvinsi != "" {
			data[k] = append(data[k], row)
		}
	}
	return data
}


// appendDataResponse for appending data to response
func appendDataResponse(data map[Key][]models.Storage,rate float64) []models.RespData {
	var arrData []models.RespData
	for k, v := range data {
		var price []float64
		var size []float64

		for _, vD := range v {
			price = append(price, vD.Price)
			size = append(size, vD.Size)
		}

		sort.Sort(sort.Float64Slice(price))
		sort.Sort(sort.Float64Slice(size))
		medianPrice := Median(price)
		medianSize := Median(size)
		avgPrice := Avg(price)
		avgSize := Avg(size)

		aggregate := models.AggregateData{
			Price: models.AggregateValue{
				Min:    price[0],
				Max:    price[len(price)-1],
				Median: medianPrice,
				Avg:    avgPrice,
			},
			Size: models.AggregateValue{
				Min:    size[0],
				Max:    size[len(size)-1],
				Median: medianSize,
				Avg:    avgSize,
			},
			PricfeUSD: models.AggregateValue{
				Min:    price[0] * rate,
				Max:    price[len(price)-1]*rate,
				Median: medianPrice*rate,
				Avg:    avgPrice*rate,
			},
		}
		respData := models.RespData{
			AreaProvinsi: k.Province,
			Week:         int32(k.Week),
			Aggregate:    aggregate,
		}

		arrData = append(arrData, respData)

	}
	return arrData
}

func Median(data []float64)float64{
	var median float64
	if len(data)%2==0{
		median = (data[len(data)/2-1]+data[len(data)/2])/2
	}else{
		median = data[len(data)/2]
	}
	return median
}

func Avg(data []float64)float64{
	var avg float64
	var sum float64
	for _,v:=range data{
		sum += v
	}
	avg = sum/float64(len(data))
	return avg
}