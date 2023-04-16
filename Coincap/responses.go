package Coincap

import "fmt"

type assetsResponse struct {
	Assets []Asset  `json:"data"`
	Timestamp int64 `json:"timestamp"`
}
type Asset struct{
	Id string `json:"id"`
	Rank string `json:"rank"`
	Symbol string `json:"symbol"`
	Name  string `json:"name"`
	Supply string `json:"supply"`
	MaxSupply string `json:"maxSupply"`
	MarketCapUsd string `json:"marketCapUsd"`
	VolumeUsd24Hr string `json:"volumeUsd24Hr"`
	PriceUsd	string `json:"priceUsd"`
	ChangePercent24Hr string	`json:"changePercent24Hr"`
	Vwap24H 	string	`json:"vwap24Hr"`

}

func (d Asset) Info() string{
	return fmt.Sprintf("[ID] %s | [RANK] %s | [SYMBOL] %s | [NAME] %s | [SUPPLY] %s | [MAXSUPPLY] %s | [MARKETCAPUSD] %s | [VOLUMEUSD24HR] %s | [PRICEUSD] %s | [CHANGEPERCENT24HR] %s | [VWAP24HR] %s" ,
		d.Id,d.Rank,d.Symbol,d.Name,d.Supply,d.MaxSupply,d.MarketCapUsd,d.VolumeUsd24Hr,d.PriceUsd,d.ChangePercent24Hr,d.Vwap24H)
}

type assetResponse struct {
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

