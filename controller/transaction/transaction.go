package transaction

import (
	"log"
	"merchant-report/model"
	"merchant-report/service/merchant"
	"merchant-report/service/outlet"
	"merchant-report/service/transaction"
	"merchant-report/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type transactionController struct {
	TransactionService transaction.TransactionService
	MerchantService    merchant.MerchantService
	OutletService      outlet.OutletService
}

func NewTransactionController(ts *transaction.TransactionService, ms *merchant.MerchantService, outlets *outlet.OutletService) transactionController {
	return transactionController{
		TransactionService: *ts,
		MerchantService:    *ms,
		OutletService:      *outlets,
	}
}

func (a *transactionController) Route(e *echo.Echo, isLogin echo.MiddlewareFunc) {
	trxGroup := e.Group("api/transactions")
	trxGroup.Use(isLogin)
	trxGroup.GET("/merchant/:id", a.ReportMerchant)
	trxGroup.GET("/outlet/:id", a.ReportOutlet)
}

func (a *transactionController) ReportMerchant(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	claims := utils.GetJWTPayload(c)
	ok := a.MerchantService.CheckMerchantByUserId(c.Request().Context(), id, int(claims.UserID))
	if !ok {
		return c.JSON(http.StatusOK, model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "BAD_REQUEST",
			Data:   "Not Found Merchant",
		})
	}

	outletId, _ := strconv.ParseUint(c.QueryParam("outlet_id"), 10, 64)
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	req := model.TrxRequest{
		MerchantID: claims.MerchantID,
		OutletID:   outletId,
		Date:       c.QueryParam("date"),
		OutletName: c.QueryParam("outlet_name"),
		Limit:      limit,
		Page:       page,
	}
	res, err := a.TransactionService.ReportMerchant(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusOK, model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})

}

func (a *transactionController) ReportOutlet(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	claims := utils.GetJWTPayload(c)
	ok := a.OutletService.CheckOutletByUserId(c.Request().Context(), id, int(claims.UserID))
	log.Println(ok)
	if !ok {
		return c.JSON(http.StatusOK, model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "BAD_REQUEST",
			Data:   "Not Found Outlet",
		})
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	page, _ := strconv.Atoi(c.QueryParam("page"))

	req := model.TrxRequest{
		OutletID:   uint64(id),
		Date:       c.QueryParam("date"),
		OutletName: c.QueryParam("outlet_name"),
		Limit:      limit,
		Page:       page,
	}
	res, err := a.TransactionService.ReportOutlet(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusOK, model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   res,
	})

}
