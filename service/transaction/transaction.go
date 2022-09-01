package transaction

import (
	"context"
	"log"
	"merchant-report/model"
	"merchant-report/schema"
	"time"

	"github.com/go-rel/rel"
)

type transactionService struct {
	DB rel.Repository
}

func NewTransactionService(db rel.Repository) TransactionService {
	return &transactionService{DB: db}
}

type TransactionService interface {
	ReportMerchant(c context.Context, req model.TrxRequest) (*model.TrxResponse, error)
	ReportOutlet(c context.Context, req model.TrxRequest) (*model.TrxResponse, error)
}

func (a *transactionService) ReportMerchant(c context.Context, req model.TrxRequest) (*model.TrxResponse, error) {
	var trx []schema.TransactionReportMerchant
	pagination := model.Pagination{}
	tDate, _ := time.Parse("2006-01", req.Date)
	if tDate.IsZero() {
		tDate = time.Now()
	}

	page := req.Page
	if req.Page <= 1 {
		page = 1
	}

	limit := req.Limit
	if req.Limit <= 1 {
		limit = 4
	}

	startDay := limit*(req.Page-1) + 1

	if req.Page <= 1 {
		startDay = 1
	}
	endDay := startDay + (limit - 1)

	startDate := DayOfMonth(tDate, startDay)
	endDate := DayOfMonth(tDate, endDay)
	eom := EndOfMonth(tDate)

	totalData := eom.Day()
	if endDay >= totalData {
		endDay = totalData
	}
	err := a.DB.FindAll(c, &trx, rel.SQL(`SELECT m.id, m.merchant_name, SUM(t.bill_total) AS omzet, CAST(t.created_at AS date) AS date
			FROM Transactions t 
			INNER JOIN Merchants m ON m.id = t.merchant_id
			WHERE t.merchant_id = ?
			AND (DATE(t.created_at) BETWEEN ? AND ?)
			GROUP BY CAST(t.created_at AS date), m.id;`, req.MerchantID, startDate, endDate))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var trxRes []schema.TransactionReportMerchant
	for i := startDay; i <= endDay; i++ {
		temp := schema.TransactionReportMerchant{}
		thisDate := DayOfMonth(tDate, i)
		for _, s := range trx {
			if thisDate.Day() == s.Date.Day() {
				temp = s
			}
		}
		if temp.Omzet == "" {
			temp.Date = thisDate
			temp.Omzet = "0"
		}

		trxRes = append(trxRes, temp)

	}

	pagination.Page = page
	pagination.TotalPage = PageCount(totalData, limit)
	pagination.TotalData = totalData
	pagination.Limit = limit

	res := model.TrxResponse{Pagination: pagination, Data: &trxRes}
	return &res, nil
}

func (a *transactionService) ReportOutlet(c context.Context, req model.TrxRequest) (*model.TrxResponse, error) {
	var trx []schema.TransactionReportOutlet
	tDate, _ := time.Parse("2006-01", req.Date)
	pagination := model.Pagination{}
	if tDate.IsZero() {
		tDate = time.Now()
	}

	page := req.Page
	if req.Page <= 1 {
		page = 1
	}

	limit := req.Limit
	if req.Limit <= 1 {
		limit = 4
	}

	startDay := limit*(page-1) + 1

	if req.Page <= 1 {
		startDay = 1
	}
	endDay := startDay + (limit - 1)

	startDate := DayOfMonth(tDate, startDay)
	endDate := DayOfMonth(tDate, endDay)
	eom := EndOfMonth(tDate)

	totalData := eom.Day()
	if endDay >= totalData {
		endDay = totalData
	}
	err := a.DB.FindAll(c, &trx, rel.SQL(`SELECT m.id, m.merchant_name, o.id, o.outlet_name, m.merchant_name, SUM(t.bill_total) AS omzet, CAST(t.created_at AS date) AS date
			FROM Transactions t 
			INNER JOIN Merchants m ON m.id = t.merchant_id
			INNER JOIN Outlets o ON o.id = t.outlet_id
			WHERE t.outlet_id = ?
			AND (DATE(t.created_at) BETWEEN ? AND ?)
			GROUP BY CAST(t.created_at AS date), m.id, o.id;`, req.OutletID, startDate, endDate))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var trxRes []schema.TransactionReportOutlet
	for i := startDay; i <= endDay; i++ {
		temp := schema.TransactionReportOutlet{}
		thisDate := DayOfMonth(tDate, i)
		for _, s := range trx {
			if thisDate.Day() == s.Date.Day() {
				temp = s
			}
		}
		if temp.Omzet == "" {
			temp.Date = thisDate
			temp.Omzet = "0"
		}

		trxRes = append(trxRes, temp)

	}

	pagination.Page = page
	pagination.TotalPage = PageCount(totalData, limit)
	pagination.TotalData = totalData
	pagination.Limit = limit

	res := model.TrxResponse{Pagination: pagination, Data: &trxRes}
	log.Println(res)
	return &res, nil
}

func BeginningOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 0, -date.Day()+1)
}
func DayOfMonth(date time.Time, day int) time.Time {
	return date.AddDate(0, 0, -date.Day()+day)
}

func EndOfMonth(date time.Time) time.Time {
	return date.AddDate(0, 1, -date.Day())
}

func PageCount(totalData int, limit int) int {
	totalPage := totalData / limit

	if totalData%limit > 0 {
		totalPage++
	}

	return totalPage

}
