package service

import (
	"Testprj03/hometest.globalgroup.com/backend/dao"
	"errors"
)

// StringService provides operations on strings.
type PrettyServiceInterface interface {
	GetAccounts() (dao.AccountListType, error)
	GetAllPayments() (dao.PaymentListType, error)
	SendPayment(req dao.SendPaymentRequestType) (dao.SendPaymentResponseType, error)
}

type PrettyServiceType struct{}

func (PrettyServiceType) GetAccounts() (dao.AccountListType, error) {
	response, err := getAccounts()
	return response, err
}

func (PrettyServiceType) GetAllPayments() (dao.PaymentListType, error) {
	response := getPayments()
	return response, nil
}

func (PrettyServiceType) SendPayment(req dao.SendPaymentRequestType) (dao.SendPaymentResponseType, error) {
	response, _ := sendPayment(req)
	return response, nil
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
