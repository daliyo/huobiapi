package huobiapi

import (
	"time"
)

// CurrencysResponse 币种响应模型
type CurrencysResponse struct {
	BaseResponse
	Data []string `json:"data,omitempty"`
}

// TimestampResponse 时间戳响应模型
type TimestampResponse struct {
	BaseResponse
	Data int64 `json:"data,omitempty"`
}

// Time 获取响应中的时间
func (o TimestampResponse) Time() time.Time {
	return time.Unix(0, o.Data*int64(time.Millisecond))
}

// SymbolResponse 交易对响应模型
type SymbolResponse struct {
	BaseResponse
	Data []SymbolItem `json:"data,omitempty"`
}

// SymbolItem 交易对数据项模型
type SymbolItem struct {
	// BaseCurrency 交易对中的基础币种
	BaseCurrency string `json:"base-currency,omitempty"`
	// QuoteCurrency 交易对中的报价币种
	QuoteCurrency string `json:"quote-currency,omitempty"`
	// PricePrecision 交易对报价的精度（小数点后位数）
	PricePrecision int `json:"price-precision,omitempty"`
	// AmountPrecision 交易对基础币种计数精度（小数点后位数）
	AmountPrecision int `json:"amount-precision,omitempty"`
	// SymbolPartition 交易区
	//
	// 可能值: [main，innovation]
	SymbolPartition string `json:"symbol-partition,omitempty"`
	// Symbol 交易对
	Symbol string `json:"symbol,omitempty"`
	// State 交易对状态
	//
	// 可能值: [online，offline,suspend] online - 已上线；offline - 交易对已下线，不可交易；suspend -- 交易暂停
	State string `json:"state,omitempty"`
	// ValuePrecision 交易对交易金额的精度（小数点后位数）
	ValuePrecision int `json:"value-precision,omitempty"`
	// MinOrderAmt 交易对最小下单量
	//
	// 下单量指当订单类型为限价单或sell-market时，下单接口传的'amount'
	MinOrderAmt float32 `json:"min-order-amt,omitempty"`
	// MaxOrderAmt 交易对最大下单量
	MaxOrderAmt float32 `json:"max-order-amt,omitempty"`
	// MinOrderValue 最小下单金额
	//
	// 下单金额指当订单类型为限价单时，下单接口传入的(amount * price)。当订单类型为buy-market时，下单接口传的'amount'
	MinOrderValue float32 `json:"min-order-value,omitempty"`
	// LeverageRatio 交易对杠杆最大倍数
	LeverageRatio float32 `json:"leverage-ratio,omitempty"`
}

// AccountBalanceResponse 账户余额响应模型
type AccountBalanceResponse struct {
	BaseResponse
	Data AccountBalanceItem `json:"data,omitempty"`
}

// AccountBalanceItem 账户余额项模型
type AccountBalanceItem struct {
	// ID 账户 ID
	ID int64 `json:"id,omitempty"`
	// State 账户状态
	//
	// working：正常 lock：账户被锁定
	State string `json:"state,omitempty"`
	// Type 账户类型
	//
	// spot：现货账户， margin：逐仓杠杆账户，otc：OTC 账户，point：点卡账户，super-margin：全仓杠杆账户
	Type string                      `json:"type,omitempty"`
	List []AccountBalanceItemDetails `json:"list,omitempty"`
}

// AccountBalanceItemDetails 账户余额明细
type AccountBalanceItemDetails struct {
	// Balance 余额
	Balance string `json:"balance,omitempty"`
	// Currency 币种
	Currency string `json:"currency,omitempty"`
	// Type 类型	trade: 交易余额，frozen: 冻结余额
	Type string `json:"type,omitempty"`
}

// AccountInfoResponse 账户信息响应模型
type AccountInfoResponse struct {
	BaseResponse
	Data []AccountInfoItem `json:"data,omitempty"`
}

// AccountInfoItem 账户数据项模型
type AccountInfoItem struct {
	// ID account-id
	ID int64 `json:"id,omitempty"`
	// State 账户状态
	State string `json:"state,omitempty"`
	// Type 账户类型
	//
	// spot：现货账户， margin：逐仓杠杆账户，otc：OTC 账户，point：点卡账户，super-margin：全仓杠杆账户
	Type string `json:"type,omitempty"`
	// SubType 子账户类型（仅对逐仓杠杆账户有效）
	SubType string `json:"sub_type,omitempty"`
}

// BaseResponse BaseResponse
type BaseResponse struct {
	Status     string `json:"status,omitempty"`
	Timestamp  int64  `json:"ts,omitempty"`
	ErrCode    string `json:"err-code,omitempty"`
	ErrMessage string `json:"err-msg,omitempty"`
}

// IsOK 判断响应内容是否OK
func (o *BaseResponse) IsOK() bool {
return o.Status == "ok"
}