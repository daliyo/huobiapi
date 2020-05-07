package huobiapi

import (
	"net/http"
	"strconv"
)

// KLineResponse K线数据响应模型
type KLineResponse struct {
	BaseResponse
	Data []KLineItem `json:"data,omitempty"`
}

// KLineItem K线数据项模型
type KLineItem struct {
	// Id 调整为新加坡时间的时间戳，单位秒，并以此作为此K线柱的id
	ID int64 `json:"id,omitempty"`
	// Amount 以基础币种计量的交易量
	Amount float32 `json:"amount,omitempty"`
	// Count 交易次数
	Count int `json:"count,omitempty"`
	// Open 本阶段开盘价
	Open float32 `json:"open,omitempty"`
	// Close 本阶段收盘价
	Close float32 `json:"close,omitempty"`
	// Low 本阶段最低价
	Low float32 `json:"low,omitempty"`
	// High 本阶段最高价
	High float32 `json:"high,omitempty"`
	// Vol 以报价币种计量的交易量
	Vol float32 `json:"vol,omitempty"`
}

// PeriodOption 数据时间粒度
type PeriodOption string

const (
	// Period1min 1min
	Period1min PeriodOption = "1min"
	// Period5min 5min
	Period5min PeriodOption = "5min"
	// Period15min 15min
	Period15min PeriodOption = "15min"
	// Period30min 30min
	Period30min PeriodOption = "30min"
	// Period60min 60min
	Period60min PeriodOption = "60min"
	// Period4hour 4hour
	Period4hour PeriodOption = "4hour"
	// Period1day 1day
	Period1day PeriodOption = "1day"
	// Period1mon 1mon
	Period1mon PeriodOption = "1mon"
	// Period1week 1week
	Period1week PeriodOption = "1week"
	// Period1year 1year
	Period1year PeriodOption = "1year"
)

// KLineOptionalParam K 线数据（蜡烛图）可选参数
type KLineOptionalParam struct {
	Size int
}

// KLine K 线数据（蜡烛图）
//
// 此接口返回历史K线数据。
//
// symbol 交易对
//
// period 返回数据时间粒度，也就是每根蜡烛的时间区间
//
// size 返回 K 线数据条数
func KLine(symbol string, period PeriodOption, opt KLineOptionalParam) (KLineResponse, error) {
	req := new(requestStruct)
	req.Method = http.MethodGet
	req.Path = "/market/history/kline"
	req.addURLParam("symbol", symbol)
	req.addURLParam("period", string(period))
	if opt.Size > 0 {
		req.addURLParam("size", strconv.Itoa(opt.Size))
	}

	res, err := sendRequest(req)
	if err != nil {
		return KLineResponse{}, err
	}
	o := new(KLineResponse)
	err = parseRespBody(res, o)
	return *o, err
}
