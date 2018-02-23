package rhclient

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Fundamental struct {
	Open          string `json:"open"`
	High          string `json:"high"`
	Low           string `json:"low"`
	Volume        string `json:"volume"`
	AverageVolume string `json:"average_volume"`
	High52Weeks   string `json:"high_52_weeks"`
	DividendYield string `json:"dividend_yield"`
	Low52Weeks    string `json:"low_52_weeks"`
	MarketCap     string `json:"market_cap"`
	PeRatio       string `json:"pe_ratio"`
	Description   string `json:"description"`
	Instrument    string `json:"instrument"`
}

type FundamentalsResponse struct {
	Results []Fundamental `json:"results"`
}

func (rh *Robinhood) GetFundamental(symbol string) (*Fundamental, error) {
	uri := fmt.Sprintf("%v/fundamentals/%v/", baseURL, symbol)
	resp, err := rh.get(uri)
	if err != nil {
		return nil, err
	}
	f := Fundamental{}
	if err := json.Unmarshal(resp.Body(), &f); err != nil {
		return nil, err
	}
	return &f, nil
}

func (rh *Robinhood) GetFundamentals(symbols []string) ([]Fundamental, error) {
	uri := fmt.Sprintf("%v/fundamentals?symbols=%v", baseURL, strings.Join(symbols, ","))
	resp, err := rh.get(uri)
	if err != nil {
		return nil, err
	}
	fResp := FundamentalsResponse{}
	if err := json.Unmarshal(resp.Body(), &fResp); err != nil {
		return nil, err
	}
	return fResp.Results, nil
}
