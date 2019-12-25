package models

type DataRest struct {
	Status string      `json:"status"`
	Data   DataDetails `json:"data"`
}

type DataDetail struct {
	Judul string
	Harga string
}

type DataDetails []DataDetail

func SetDataDetailRest() *DataDetail {
	return &DataDetail{}
}
