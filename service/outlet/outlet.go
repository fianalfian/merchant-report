package outlet

import (
	"context"
	"log"
	"merchant-report/schema"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
)

type outletService struct {
	DB rel.Repository
}

func NewOutletService(db rel.Repository) OutletService {
	return &outletService{DB: db}
}

type OutletService interface {
	CheckOutletByUserId(c context.Context, merchantId int, userId int) bool
}

func (a *outletService) CheckOutletByUserId(ctx context.Context, outlet_id int, userId int) bool {
	merchant := []schema.Merchant{}
	err := a.DB.FindAll(ctx, &merchant, rel.Select("id"), where.Eq("user_id", userId))
	if err != nil {
		return false
	}

	values := []any{}
	for _, s := range merchant {
		values = append(values, s.ID)
	}

	outlet := schema.Outlet{}
	err = a.DB.Find(ctx, &outlet, rel.Select("id", "outlet_name"), where.Eq("id", outlet_id), where.In("merchant_id", values...))
	if err != nil {
		return false
	}

	log.Println(outlet)

	return true
}
