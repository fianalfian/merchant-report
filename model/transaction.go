package model

import "time"

type TrxRequest struct {
	MerchantID   uint64
	MerchantName string
	OutletID     uint64
	OutletName   string
	Date         string
	Limit        int
	Page         int
}

type ReportFilter struct {
	MerchantID uint64
	OutletID   uint64
	Date       time.Time
	Limit      int
	Page       int
}

type Pagination struct {
	Limit     int `json:"limit"`
	Page      int `json:"page"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
}

type TrxResponse struct {
	Pagination Pagination `json:"pagination"`
	Data       any        `json:"data"`
}
