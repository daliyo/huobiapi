package huobiapi

import (
	"net/http"
)

// Currencys 获取所有币种
//
// 此接口返回所有火币全球站支持的币种。
func Currencys() (CurrencysResponse, error) {
	req := new(requestStruct)
	req.Method = http.MethodGet
	req.Path = "/v1/common/currencys"

	res, err := sendRequest(req)
	if err != nil {
		return CurrencysResponse{}, err
	}
	o := new(CurrencysResponse)
	err = parseRespBody(res, o)
	return *o, err
}

// Timestamp 获取当前系统时间戳
//
// 此接口返回当前的系统时间戳，即从 UTC 1970年1月1日0时0分0秒0毫秒到现在的总毫秒数。
func Timestamp() (TimestampResponse, error) {
	req := new(requestStruct)
	req.Method = http.MethodGet
	req.Path = "/v1/common/timestamp"

	res, err := sendRequest(req)
	if err != nil {
		return TimestampResponse{}, err
	}
	o := new(TimestampResponse)
	err = parseRespBody(res, o)
	return *o, err
}

// Symbols 获取所有交易对
//
// 此接口返回所有火币全球站支持的交易对。
func Symbols() (SymbolResponse, error) {
	req := new(requestStruct)
	req.Method = http.MethodGet
	req.Path = "/v1/common/symbols"

	res, err := sendRequest(req)
	if err != nil {
		return SymbolResponse{}, err
	}
	o := new(SymbolResponse)
	err = parseRespBody(res, o)
	return *o, err
}
