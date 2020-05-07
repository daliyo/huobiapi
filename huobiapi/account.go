package huobiapi

import (
	"fmt"
	"net/http"
)

// AccountBalance 账户余额
//
// 查询指定账户的余额，支持以下账户
//
// spot：现货账户， margin：逐仓杠杆账户，otc：OTC 账户，point：点卡账户，super-margin：全仓杠杆账户
func AccountBalance(accountID string) (AccountBalanceResponse, error) {
	req := new(requestStruct)
	req.Method = http.MethodGet
	req.Path = fmt.Sprintf("/v1/account/accounts/%s/balance", accountID)
	req.Sign = true

	res, err := sendRequest(req)
	if err != nil {
		return AccountBalanceResponse{}, err
	}
	o := new(AccountBalanceResponse)
	err = parseRespBody(res, o)
	return *o, err
}

// Accounts 账户信息
//
// 查询当前用户的所有账户 ID account-id 及其相关信息
func Accounts() (AccountInfoResponse, error) {
	req := new(requestStruct)
	req.Method = http.MethodGet
	req.Path = "/v1/account/accounts"
	req.Sign = true

	res, err := sendRequest(req)
	if err != nil {
		return AccountInfoResponse{}, err
	}
	o := new(AccountInfoResponse)
	err = parseRespBody(res, o)
	return *o, nil
}
