package backend

import (
	"Testprj03/hometest.globalgroup.com/backend/helper"
	"Testprj03/hometest.globalgroup.com/backend/service"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func DoMainCtl() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc service.PrettyServiceInterface

	svc = service.PrettyServiceType{}
	svc = helper.LoggingMiddlewareType{Logger: logger, Next: svc}

	accountHandler := httptransport.NewServer(
		helper.MakeAccountEndpoint(svc),
		helper.DecodeAccountRequest,
		helper.EncodeResponse,
	)

	paymentHandler := httptransport.NewServer(
		helper.MakePaymentEndpoint(svc),
		helper.DecodePaymentRequest,
		helper.EncodeResponse,
	)

	sendPaymentHandler := httptransport.NewServer(
		helper.MakeSendPaymentEndpoint(svc),
		helper.DecodeSendPaymentRequest,
		helper.EncodeResponse,
	)
	http.Handle("/api/v1/account", accountHandler)
	http.Handle("/api/v1/payment", paymentHandler)
	http.Handle("/api/v1/account/0", sendPaymentHandler)

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
