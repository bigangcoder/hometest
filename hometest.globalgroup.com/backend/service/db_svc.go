package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"Testprj03/hometest.globalgroup.com/backend/config"
	"Testprj03/hometest.globalgroup.com/backend/dao"

	"github.com/go-orm/gorm"
	_ "github.com/go-orm/gorm/dialects/postgres"
)

var baseConfigType dao.BaseConfigType
var connectionUrl string

func init() {
	loadJsonConfig()
	connectionUrl = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", baseConfigType.DBHost, baseConfigType.DBPort, baseConfigType.DBUser, baseConfigType.DBName, baseConfigType.DBPassword)
}
func getConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open(baseConfigType.Dialect, connectionUrl)
	return
}
func getAccounts() (dao.AccountListType, error) {
	accountList := dao.AccountListType{}
	accountList.Accounts = make([]dao.Account, 0)
	db, err := getConnection()
	defer db.Close()
	if err != nil {
		return accountList, err
	}

	db.Raw("SELECT * FROM user_account WHERE disabled=0").Scan(&accountList.Accounts)
	return accountList, nil
}
func getPayments() (paymentList dao.PaymentListType) {
	paymentList.Payments = make([]dao.PaymentType, 0)
	db, err := getConnection()
	if err != nil {
		panic(fmt.Sprintf("database connection error:%#v", err))
	}
	defer db.Close()
	db.Raw("SELECT * FROM user_payment WHERE disabled=0").Scan(&paymentList.Payments)
	return
}

func sendPayment(req dao.SendPaymentRequestType) (dao.SendPaymentResponseType, error) {
	db, err := getConnection()
	if err != nil {
		panic(fmt.Sprintf("database connection error:%#v", err))
	}
	defer db.Close()
	account := dao.Account{}
	toAccount := dao.Account{}
	txRollback := true
	tx := db.Begin()
	defer func() {
		if txRollback {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err := tx.Table(dao.TABLE_ACCOUNT).Find(&account, "account=?", req.Account).Error; err != nil {
		return dao.SendPaymentResponseType{RespCode: 500, RespInfo: "connect db error happened"}, err
	}
	if err := tx.Table(dao.TABLE_ACCOUNT).Find(&toAccount, "account=?", req.ToAccount).Error; err != nil {
		return dao.SendPaymentResponseType{RespCode: 500, RespInfo: "connect db error happened"}, err
	}
	if req.Amount <= 0 {
		return dao.SendPaymentResponseType{RespCode: 499, RespInfo: "error: the sending payment amount could not less or equal than 0"}, errors.New("amount error")
	}
	if account.Currency != toAccount.Currency {
		return dao.SendPaymentResponseType{RespCode: 498, RespInfo: "error: should be the same currency"}, errors.New("currency error")
	}
	if account.Balance <= 0 || account.Balance-req.Amount < 0 {
		return dao.SendPaymentResponseType{RespCode: 497, RespInfo: "error: balance is insufficient"}, errors.New("balance error")
	}
	account.Balance -= req.Amount
	toAccount.Balance += req.Amount

	if err := tx.Table(dao.TABLE_ACCOUNT).Save(&account).Error; err != nil {
		return dao.SendPaymentResponseType{RespCode: 496, RespInfo: "error: update error"}, err
	}
	if err := tx.Table(dao.TABLE_ACCOUNT).Save(&toAccount).Error; err != nil {
		return dao.SendPaymentResponseType{RespCode: 496, RespInfo: "error: update error"}, err
	}
	createdAt := fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"))
	payment := dao.PaymentType{Account: account.Account, Amount: req.Amount, ToAccount: toAccount.Account, FromAccount: "", Direction: dao.OUTGOING, CreatedAt: createdAt}
	if err := tx.Table(dao.TABLE_PAYMENT).Save(&payment).Error; err != nil {
		return dao.SendPaymentResponseType{RespCode: 496, RespInfo: "error: update error"}, err
	}
	payment = dao.PaymentType{Account: toAccount.Account, Amount: req.Amount, ToAccount: "", FromAccount: account.Account, Direction: dao.INCOMING, CreatedAt: createdAt}
	if err := tx.Table(dao.TABLE_PAYMENT).Save(&payment).Error; err != nil {
		return dao.SendPaymentResponseType{RespCode: 496, RespInfo: "error: update error"}, err
	}
	txRollback = false
	return dao.SendPaymentResponseType{AccountBalance: account.Balance, Amount: req.Amount, ToAccountBalance: toAccount.Balance, RespCode: 0, RespInfo: ""}, nil
}

func loadJsonConfig() {
	data, err := ioutil.ReadFile(config.GetConfigFileName())
	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}
	json.Unmarshal(data, &baseConfigType)
}
