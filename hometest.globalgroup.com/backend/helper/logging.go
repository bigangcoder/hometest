package helper

import (
	"time"

	"Testprj03/hometest.globalgroup.com/backend/dao"

	"Testprj03/hometest.globalgroup.com/backend/service"

	"github.com/go-kit/kit/log"
)

type LoggingMiddlewareType struct {
	Logger log.Logger
	Next   service.PrettyServiceInterface
}

func (mw LoggingMiddlewareType) GetAccounts() (output dao.AccountListType, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "uppercase",
			"input", "none",
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.GetAccounts()
	return
}
func (mw LoggingMiddlewareType) GetAllPayments() (output dao.PaymentListType, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "uppercase",
			"input", "none",
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.GetAllPayments()
	return
}

func (mw LoggingMiddlewareType) SendPayment(req dao.SendPaymentRequestType) (output dao.SendPaymentResponseType, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "SendPayment",
			"input", "none",
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.SendPayment(req)
	return
}
