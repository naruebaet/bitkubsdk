package bitkubsdk

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/naruebaet/bitkubsdk/src/model"
	"github.com/naruebaet/bitkubsdk/src/pkg/curl"
	"github.com/naruebaet/bitkubsdk/src/pkg/socket"
	"net/http"
	"strings"
)

const ApiHost = "https://api.bitkub.com"

type BitkubAccess struct {
	ApiKey    string
	ApiSecret string
}

type BitkubRepository interface {
	// non-secure zone
	// rest api
	GetServerStatus() (response []model.ServerStatusResponse, err error)
	GetServerTime() (response string, err error)
	GetSymbols() (response model.SymbolResponse, err error)
	GetTicker(sym string) (response map[string]model.TickerResponseResult, err error)
	GetTrades(sym string, lmt int) (response model.TradeResponse, err error)
	GetBids(sym string, lmt int) (response model.BidsAskResponse, err error)
	GetAsks(sym string, lmt int) (response model.BidsAskResponse, err error)
	GetBooks(sym string, lmt int) (response model.BooksResponse, err error)
	GetTradingview(sym string, frame, frm, lmt int) (response model.TradingviewResponse, err error)
	GetDepth(sym string, lmt int) (response model.BooksResponse, err error)
	// socket
	WatchTicker(ctx context.Context, operations func(conn *websocket.Conn))
	WatchTrade(ctx context.Context, operations func(conn *websocket.Conn))

	// secure zone
	// POST /api/market/wallet
	// POST /api/market/balances
	// POST /api/market/place-bid
	// POST /api/market/place-ask
	// POST /api/market/place-bid/test
	// POST /api/market/place-ask/test
	// POST /api/market/place-ask-by-fiat
	// POST /api/market/cancel-order
	// POST /api/market/my-open-orders
	// POST /api/market/my-order-history
	// POST /api/market/order-info
	// POST /api/crypto/addresses
	// POST /api/crypto/withdraw
	// POST /api/crypto/deposit-history
	// POST /api/crypto/withdraw-history
	// POST /api/crypto/generate-address
	// POST /api/fiat/accounts
	// POST /api/fiat/withdraw
	// POST /api/fiat/deposit-history
	// POST /api/fiat/withdraw-history
	// POST /api/market/wstoken
	// POST /api/user/limits
	// POST /api/user/trading-credits
}

func NewBitkub(apiKey string, apiSecret string) BitkubRepository {
	// go interface is awesome!
	// why you do that work!
	var bk BitkubRepository
	bk = &BitkubAccess{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
	return bk
}

func (bk *BitkubAccess) GetServerStatus() (response []model.ServerStatusResponse, err error) {
	statusCode, err := curl.HttpGet(ApiHost+"/api/status", &response)
	if err != nil || statusCode != http.StatusOK {
		return nil, err
	}

	return response, nil
}

func (bk *BitkubAccess) GetServerTime() (response string, err error) {
	raw, statusCode, err := curl.HttpGetRaw(ApiHost + "/api/servertime")
	if err != nil || statusCode != http.StatusOK {
		return "", err
	}

	response = string(raw)

	return response, nil
}

func (bk *BitkubAccess) GetSymbols() (response model.SymbolResponse, err error) {
	statusCode, err := curl.HttpGet(ApiHost+"/api/market/symbols", &response)
	if err != nil || statusCode != http.StatusOK {
		return model.SymbolResponse{}, err
	}

	return response, nil
}

func (bk *BitkubAccess) GetTicker(sym string) (response map[string]model.TickerResponseResult, err error) {
	statusCode, err := curl.HttpGet(ApiHost+"/api/market/ticker?sym="+sym, &response)
	if err != nil || statusCode != http.StatusOK {
		return map[string]model.TickerResponseResult{}, err
	}

	return response, nil
}

func (bk *BitkubAccess) GetTrades(sym string, lmt int) (response model.TradeResponse, err error) {
	statusCode, err := curl.HttpGet(fmt.Sprintf("%s/api/market/trades?sym=%s&lmt=%d", ApiHost, sym, lmt), &response)
	if err != nil || statusCode != http.StatusOK {
		return model.TradeResponse{}, err
	}

	return response, nil
}

func (bk *BitkubAccess) GetBids(sym string, lmt int) (response model.BidsAskResponse, err error) {
	statusCode, err := curl.HttpGet(fmt.Sprintf("%s/api/market/bids?sym=%s&lmt=%d", ApiHost, sym, lmt), &response)
	if err != nil || statusCode != http.StatusOK {
		return model.BidsAskResponse{}, err
	}

	return response, nil
}

func (bk *BitkubAccess) GetAsks(sym string, lmt int) (response model.BidsAskResponse, err error) {
	statusCode, err := curl.HttpGet(fmt.Sprintf("%s/api/market/asks?sym=%s&lmt=%d", ApiHost, sym, lmt), &response)
	if err != nil || statusCode != http.StatusOK {
		return model.BidsAskResponse{}, err
	}

	return response, nil
}

func (bk *BitkubAccess) GetBooks(sym string, lmt int) (response model.BooksResponse, err error) {
	statusCode, err := curl.HttpGet(fmt.Sprintf("%s/api/market/books?sym=%s&lmt=%d", ApiHost, sym, lmt), &response)
	if err != nil || statusCode != http.StatusOK {
		return model.BooksResponse{}, err
	}

	return response, nil
}

func (bk *BitkubAccess) GetTradingview(sym string, interval, frm, lmt int) (response model.TradingviewResponse, err error) {
	query := fmt.Sprintf("?sym=%s&int=%d&frm=%d&lmt=%d", sym, interval, frm, lmt)
	statusCode, err := curl.HttpGet(ApiHost+"/api/market/tradingview"+query, &response)
	if err != nil || statusCode != http.StatusOK {
		return model.TradingviewResponse{}, err
	}

	return response, nil
}

func (bk *BitkubAccess) GetDepth(sym string, lmt int) (response model.BooksResponse, err error) {
	statusCode, err := curl.HttpGet(fmt.Sprintf("%s/api/market/depth?sym=%s&lmt=%d", ApiHost, sym, lmt), &response)
	if err != nil || statusCode != http.StatusOK {
		return model.BooksResponse{}, err
	}

	return response, nil
}

func (bk *BitkubAccess) WatchTrade(ctx context.Context, operations func(conn *websocket.Conn)) {
	var streamData []string
	syms, _ := bk.GetSymbols()
	for _, sym := range syms.Result {
		symbol := strings.ToLower(sym.Symbol)
		streamData = append(streamData, "market.trade."+symbol)
	}

	wsConn := socket.GetWSSession("/websocket-api", strings.Join(streamData, ","))
	defer wsConn.Close()

	go operations(wsConn)

	<-ctx.Done()
	wsConn.Close()
	return
}

func (bk *BitkubAccess) WatchTicker(ctx context.Context, operations func(conn *websocket.Conn)) {
	var streamData []string
	syms, _ := bk.GetSymbols()
	for _, sym := range syms.Result {
		symbol := strings.ToLower(sym.Symbol)
		streamData = append(streamData, "market.ticker."+symbol)
	}

	wsConn := socket.GetWSSession("/websocket-api", strings.Join(streamData, ","))
	defer wsConn.Close()

	go operations(wsConn)

	<-ctx.Done()
	wsConn.Close()
	return
}
