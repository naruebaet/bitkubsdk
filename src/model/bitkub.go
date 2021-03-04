package model

type ServerStatusResponse struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ServerTimeResponse struct {
	Error  int                        `json:"error"`
	Result []ServerTimeResponseResult `json:"result"`
}

type ServerTimeResponseResult struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Info   string `json:"info"`
}

type TickerResponseResult struct {
	ID            int     `json:"id"`
	Last          float64 `json:"last"`
	LowestAsk     float64 `json:"lowestAsk"`
	HighestBid    float64 `json:"highestBid"`
	PercentChange float64 `json:"percentChange"`
	BaseVolume    float64 `json:"baseVolume"`
	QuoteVolume   float64 `json:"quoteVolume"`
	IsFrozen      int     `json:"isFrozen"`
	High24Hr      float64 `json:"high24hr"`
	Low24Hr       float64 `json:"low24hr"`
	Change        float64 `json:"change"`
	PrevClose     float64 `json:"prevClose"`
	PrevOpen      float64 `json:"prevOpen"`
}

type SymbolResponse struct {
	Error  int `json:"error"`
	Result []struct {
		ID     int    `json:"id"`
		Symbol string `json:"symbol"`
		Info   string `json:"info"`
	} `json:"result"`
}

type TradeResponse struct {
	Error  int             `json:"error"`
	Result [][]interface{} `json:"result"`
	// timestamp
	// rate
	// amount
	// side
}

type BidsAskResponse struct {
	Error  int         `json:"error"`
	Result [][]float64 `json:"result"`
	// order id
	// timestamp
	// volume
	// rate
	// amount
}

type BooksResponse struct {
	Error  int `json:"error"`
	Result struct {
		Bids [][]float64 `json:"bids"`
		// order_id
		// timestamp
		// volume
		// rate
		// amount
		Asks [][]float64 `json:"asks"`
		// order_id
		// timestamp
		// volume
		// rate
		// amount
	} `json:"result"`
}

type TradingviewResponse struct {
	Open      []int     `json:"o"`
	Close     []int     `json:"c"`
	High     []int     `json:"h"`
	Last      []int     `json:"l"`
	Symbol    string    `json:"s"`
	Vol       []float64 `json:"v"`
	Timestamp []int     `json:"t"`
}

type DepthResponse struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}