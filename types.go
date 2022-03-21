package nomics

type CurrenciesTickerRequest struct {
	IDs                string
	Interval           string
	Convert            string
	Status             string
	Filter             string
	PlatformCurrency   string
	Sort               string
	IncludeTransparncy string
	PerPage            string
	Page               string
}

type CurrenciesTickerResponse struct {
	Currency             string `json:"currency"`
	ID                   string `json:"id"`
	Status               string `json:"status"`
	Price                string `json:"price"`
	PriceDate            string `json:"price_date"`
	PriceTimestamp       string `json:"price_timestamp"`
	Symbol               string `json:"symbol"`
	CirculatingSupply    string `json:"circulating_supply"`
	MaxSupply            string `json:"max_supply"`
	Name                 string `json:"name"`
	LogoUrl              string `json:"logo_url"`
	MarketCap            string `json:"market_cap"`
	MarketCapDominance   string `json:"market_cap_dominance"`
	TransparentMarketCap string `json:"transparent_market_cap"`
	NumExchanges         string `json:"num_exchanges"`
	NumPairs             string `json:"num_pairs"`
	NumPairsUnmapped     string `json:"num_pairs_unmapped"`
	FirstCandle          string `json:"first_candle"`
	FirstTrade           string `json:"first_trade"`
	FirstOrderBook       string `json:"first_order_book"`
	FirstPricedAt        string `json:"first_priced_at"`
	Rank                 string `json:"rank"`
	RankDelta            string `json:"rank_delta"`
	High                 string `json:"high"`
	HighTimestamp        string `json:"high_timestamp"`
}
