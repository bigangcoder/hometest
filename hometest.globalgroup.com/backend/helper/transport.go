package helper

import (
	"context"
	"encoding/json"
	"net/http"

	"Testprj03/hometest.globalgroup.com/backend/dao"
	"Testprj03/hometest.globalgroup.com/backend/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeAccountEndpoint(svc service.PrettyServiceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		response, err := svc.GetAccounts()
		if err != nil {
			return dao.CommmonResponseType{RespCode: -1, RespInfo: err.Error()}, nil
		}
		return response, nil
	}
}

func MakePaymentEndpoint(svc service.PrettyServiceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		response, err := svc.GetAllPayments()
		if err != nil {
			return dao.CommmonResponseType{RespCode: -1, RespInfo: err.Error()}, nil
		}
		return response, nil
	}
}

func DecodePaymentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func MakeSendPaymentEndpoint(svc service.PrettyServiceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(dao.SendPaymentRequestType)
		response, err := svc.SendPayment(req)
		if err != nil {
			return dao.CommmonResponseType{RespCode: -1, RespInfo: err.Error()}, nil
		}
		return response, nil
	}
}

func DecodeSendPaymentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request dao.SendPaymentRequestType
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeAccountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,GET,POST,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,X-Requested-With")
	return json.NewEncoder(w).Encode(response)
}
