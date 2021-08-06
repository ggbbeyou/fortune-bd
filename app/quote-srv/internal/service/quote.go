package service

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/empty"
	"time"
	quotepb "wq-fotune-backend/api/quote"
	"wq-fotune-backend/api/response"
	"wq-fotune-backend/app/quote-srv/cron"
	"wq-fotune-backend/libs/exchange"
	"wq-fotune-backend/libs/logger"
)

const (
	errID = "quote"
	USDT  = "usdt"
	BTC   = "btc"
)

func (q *QuoteService) GetTicksWithExchangeSymbol(ctx context.Context, req *quotepb.GetTicksSymbolReq, resp *quotepb.TickResp) error {
	var tickArrayAll = make([]cron.Ticker, 0)
	//var tickerAll = make(map[string]map[string]interface{})
	if req.Exchange == exchange.BINANCE {
		if req.Symbol == USDT {
			tickArrayAll = cron.BinanceTickArrayAll
		}
		if req.Symbol == BTC {
			tickArrayAll = cron.BinaceTickArrayBtc
		}
	}
	if req.Exchange == exchange.HUOBI {
		if req.Symbol == USDT {
			tickArrayAll = cron.HuobiTickArrayAll
		}
		if req.Symbol == BTC {
			tickArrayAll = cron.HuobiTickArrayBtc
		}
	}
	if req.Exchange == exchange.OKEX {
		if req.Symbol == USDT {
			tickArrayAll = cron.OkexTickArrayAll
		}
		if req.Symbol == BTC {
			tickArrayAll = cron.OkexTickArrayBtc
		}
	}

	ticks, err := json.Marshal(tickArrayAll)
	if err != nil {
		return response.NewInternalServerErrMsg(errID)
	}
	resp.Ticks = ticks
	return nil
}

func (q *QuoteService) GetTicksWithExchange(ctx context.Context, req *quotepb.GetTicksReq, resp *quotepb.TickResp) error {
	var tickerAll = make(map[string]map[string]interface{})
	tickerAll[exchange.BINANCE] = map[string]interface{}{
		"usdt": cron.BinanceTickMapAll,
	}
	ticks, err := json.Marshal(tickerAll)
	if err != nil {
		return response.NewInternalServerErrMsg(errID)
	}
	resp.Ticks = ticks
	return nil
}

func (q *QuoteService) GetRate(ctx context.Context, e *empty.Empty, rmb *quotepb.RateUsdRmb) error {
	bytes, err := q.dao.RedisCli.Get(cron.RateKey).Bytes()
	if err != nil {
		logger.Warnf("redis取汇率错误 %v", err)
		return response.NewInternalServerErrMsg(errID)
	}
	var rate cron.QuoteRate
	if err := json.Unmarshal(bytes, &rate); err != nil {
		logger.Warnf("解析redis获取的汇率数据失败 %v", err)
		return response.NewInternalServerErrMsg(errID)
	}
	rmb.InstrumentID = rate.InstrumentID
	rmb.Rate = rate.Rate
	rmb.Timestamp = rate.Timestamp
	return nil
}



//func (q *QuoteService) GetOkexTicks(ctx context.Context, req *empty.Empty, resp *fotune_srv_quote.OkexTickResp) error {
//}

func (q *QuoteService) StreamOkexTicks(ctx context.Context, req *quotepb.GetTicksReq, resp quotepb.QuoteService_StreamOkexTicksStream) error {
	for {
		tickArrayAll := cron.OkexTickArrayAll
		if req.Exchange == exchange.HUOBI {
			tickArrayAll = cron.HuobiTickArrayAll
		}
		if req.Exchange == exchange.BINANCE {
			tickArrayAll = cron.BinanceTickArrayAll
		}

		if len(tickArrayAll) == 0 {
			time.Sleep(2 * time.Second)
			continue
		}
		ticks, err := json.Marshal(tickArrayAll)
		if err != nil {
			logger.Infof("streamOkexTicks json Marshal err %v", err)
			return err
		}
		err = resp.Send(&quotepb.TickResp{Ticks: ticks})
		if err != nil {
			logger.Infof("streamOkexTicks sendMsg err %v", err)
			return err
		}
		time.Sleep(6 * time.Second)
		continue
	}
}

func (q *QuoteService) GetOkexTicks(ctx context.Context, req *quotepb.GetTicksReq) *quotepb.TickResp {
	tickArrayAll := cron.OkexTickArrayAll
	if req.Exchange == exchange.HUOBI {
		tickArrayAll = cron.HuobiTickArrayAll
	}
	if req.Exchange == exchange.BINANCE {
		tickArrayAll = cron.BinanceTickArrayAll
	}
	if len(tickArrayAll) == 0 {
		time.Sleep(2 * time.Second)
	}
	ticks, err := json.Marshal(tickArrayAll)
	if err != nil {
		logger.Infof("streamOkexTicks json Marshal err %v", err)
		return nil
	}
	return &quotepb.TickResp{Ticks: ticks}
}
