# RESTFul api docs
> Version：v1.0.0<br>
> BaseUrl: http://127.0.0.1:8080


### 1. see available accounts api
> Author：bigang

#### request docs
> method：GET<br>
request URL ：[/api/v1/account](#)


#### request
```json
 {
	"token": "this a sample value for token"
} 
```
#### request params

| field | type   | required | comment                                                |
| ----- | ------ | -------- | ------------------------------------------------------ |
| token | string | Y        | a token to HTTP requests in order to authenticate them |


#### response
```json
 {
	"accounts": [
		{
			"id": 0,
			"account": "",
			"balance": 0,
			"currency": ""
		}
	]
} 
```
#### response params

| field    | type         | required             | comment |
| -------- | ------------ | -------------------- | ------- |
| accounts | struct slice | [Click](#1.accounts) |

<a id="1.accounts"></a> 
##### accounts 
 
| field    | type   | required                                               | comment |
| -------- | ------ | ------------------------------------------------------ | ------- |
| id       | int    |                                                        |
| account  | string |                                                        |
| balance  | int    | unit:us dollar cents , one dolar is equal to 100 cents |
| currency | string |                                                        |
 



### 2. see all payments api
> Author：bigang

#### request docs
> method：GET<br>
request URL ：[/api/v1/payment](#)


#### request
```json
 {
	"token": "this a sample value for token"
} 
```
#### request params

| field | type   | required | comment                                                |
| ----- | ------ | -------- | ------------------------------------------------------ |
| token | string | Y        | a token to HTTP requests in order to authenticate them |


#### response
```json
 {
	"payments": [
		{
			"id": 0,
			"account": "",
			"amount": 0,
			"to_account": "",
			"from_account": "",
			"direction": "",
			"created_at": ""
		}
	]
} 
```
#### response params

| field    | type         | required             | comment |
| -------- | ------------ | -------------------- | ------- |
| payments | struct slice | [Click](#2.payments) |

<a id="2.payments"></a> 
##### payments 
 
| field       | type   | required                                               | comment |
| ----------- | ------ | ------------------------------------------------------ | ------- |
| id          | int    |                                                        |
| account     | string |                                                        |
| amount      | int    | unit:us dollar cents , one dolar is equal to 100 cents |
| toAccount   | string |                                                        |
| fromAccount | string |                                                        |
| direction   | string |                                                        |
| createdAt   | string |                                                        |
 



### 3. sending payment api from one account to another account
> Author：bigang

#### request docs
> method：POST<br>
request URL ：[/api/v1/account/0](#)


#### request
```json
 {
	"token": "",
	"account": "",
	"amount": 0,
	"to_account": ""
} 
```
#### request params

| field     | type   | required | comment                                                                                                          |
| --------- | ------ | -------- | ---------------------------------------------------------------------------------------------------------------- |
| token     | string | Y        | a token to HTTP requests in order to authenticate them                                                           |
| account   | string | Y        | user account (outgoing)                                                                                          |
| amount    | int    | Y        | the amount to be tranfering (should be in the same currency),unit:us dollar cents , one dolar is equal 100 cents |
| toAccount | string | Y        | user account (incomming)                                                                                         |


#### response
```json
 {
	"accountBalance": 0,
	"amount": 0,
	"toAccountBalance": 0,
	"respCode": 0,
	"respMsg": ""
} 
```
#### response params

| field            | type   | required                                                               | comment |
| ---------------- | ------ | ---------------------------------------------------------------------- | ------- |
| accountBalance   | int    | unit:us dollar cents , one dolar is equal to 100 cents                 |
| amount           | int    | unit:us dollar cents , one dolar is equal to 100 cents                 |
| toAccountBalance | int    | unit:us dollar cents , one dolar is equal to 100 cents                 |
| respCode         | int    | resonse code, 0 for transcation ok, other values for wrong in server s |
| respInfo         | string | response information from the server s                                 |



