package schema

import (
	"encoding/json"
	"time"
)

type Transaction struct {
	ID         uint64    `json:"id" db:"id,primary"`
	MerchantID uint64    `json:"merchant_id" db:"merchant_id"`
	OutletID   uint64    `json:"outlet_id" db:"outlet_id"`
	BillTotal  float64   `json:"bill_total" db:"bill_total"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	CreatedBy  uint64    `json:"created_by" db:"created_by"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy  uint64    `json:"updated_by" db:"updated_by"`
}

func (Transaction) Table() string {
	return "Transactions"
}

type CountTransactionReport struct {
	Count int `json:"count" db:"count"`
}

type TransactionReportMerchant struct {
	MerchantName string    `json:"merchant_name" db:"merchant_name"`
	Date         time.Time `json:"date" db:"date"`
	Omzet        string    `json:"omzet" db:"omzet"`
}

type TransactionReportOutlet struct {
	MerchantName string    `json:"merchant_name" db:"merchant_name"`
	OutletName   string    `json:"outlet_name" db:"outlet_name"`
	Date         time.Time `json:"date" db:"date"`
	Omzet        string    `json:"omzet" db:"omzet"`
}

func (d *TransactionReportMerchant) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"merchant_name": d.MerchantName,
		"date":          d.Date.Format("2006-01-02"),
		"omzet":         d.Omzet,
	})
}

func (d *TransactionReportOutlet) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"merchant_name": d.MerchantName,
		"outlet_name":   d.OutletName,
		"date":          d.Date.Format("2006-01-02"),
		"omzet":         d.Omzet,
	})
}
