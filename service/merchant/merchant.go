package merchant

import (
	"context"
	"merchant-report/schema"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
)

type merchantService struct {
	DB rel.Repository
}

func NewMerchantService(db rel.Repository) MerchantService {
	return &merchantService{DB: db}
}

type MerchantService interface {
	CheckMerchantByUserId(c context.Context, merchantId int, userId int) bool
}

func (a *merchantService) CheckMerchantByUserId(ctx context.Context, merchantId int, userId int) bool {
	merchant := schema.Merchant{}
	err := a.DB.Find(ctx, &merchant, where.Eq("user_id", userId), where.Eq("id", merchantId))
	if err != nil {
		return false
	}

	return true
}
