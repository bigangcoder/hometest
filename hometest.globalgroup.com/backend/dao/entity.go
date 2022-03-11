package dao

type BaseConfigType struct {
	DBHost        string `json:"dbHost"`
	DBName        string `json:"dbName"`
	DBPassword    string `json:"dbPassword"`
	DBUser        string `json:"dbUser"`
	DBPort        int    `json:"dbPort"`
	JiyunUpperUrl string `json:"jiyunUpperUrl"`
	Dialect       string `json:"dialect"`
}

type CommonRequestType struct {
	Token string `json:"token" doc:"a token to HTTP requests in order to authenticate them required"`
}

type CommmonResponseType struct {
	RespCode int    `json:"respCode"`
	RespInfo string `json:"respInfo"`
}
type AccountListType struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Id       int    `json:"id"`
	Account  string `json:"account"`
	Balance  int    `json:"balance" doc:"unit:us dollar cents , one dolar is equal to 100 cents "`
	Currency string `json:"currency"`
}

type PaymentListType struct {
	Payments []PaymentType `json:"payments"`
}

type PaymentType struct {
	Token       string `json:"token"  gorm:"-" doc:"a token to HTTP requests in order to authenticate them required"`
	Id          int    `json:"id"`
	Account     string `json:"account"`
	Amount      int    `json:"amount" doc:"unit:us dollar cents , one dolar is equal to 100 cents "`
	ToAccount   string `json:"to_account"`
	FromAccount string `json:"from_account"`
	Direction   string `json:"direction"`
	CreatedAt   string `json:"created_at"`
}

type SendPaymentRequestType struct {
	Account   string `json:"account" doc:"user account (outgoing) required"`
	Amount    int    `json:"amount" doc:"the amount to be tranfering (should be in the same currency),unit:us dollar cents , one dolar is equal 100 cents required"`
	ToAccount string `json:"to_account" doc:"user account (incomming) required"`
}

type SendPaymentResponseType struct {
	AccountBalance   int    `json:"accountBalance" doc:"unit:us dollar cents , one dolar is equal to 100 cents " `
	Amount           int    `json:"amount" doc:"unit:us dollar cents , one dolar is equal to 100 cents "`
	ToAccountBalance int    `json:"toAccountBalance" doc:"unit:us dollar cents , one dolar is equal to 100 cents "`
	RespCode         int    `json:"respCode" doc:"resonse code, 0 for transcation ok, other values for wrong in server side"`
	RespInfo         string `json:"respMsg" doc:"response information from the server side"`
}
