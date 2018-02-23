package rhclient

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

type TimeInForce string

const (
	GoodForDay        TimeInForce = "gfd"
	GoodToClose       TimeInForce = "gtc"
	FillOrKill        TimeInForce = "fok"
	ImmediateOrCancel TimeInForce = "ioc"
	OPG               TimeInForce = "opg"
)

type Trigger string

const (
	TriggerImmediate Trigger = "immediate"
	TriggerStop      Trigger = "stop"
)

type Side string

const (
	Buy  Side = "buy"
	Sell Side = "sell"
)

type RHOrder struct {
	Account                string          `json:"account"`
	Instrument             string          `json:"instrument"`
	Symbol                 string          `json:"symbol"`
	TimeInForce            TimeInForce     `json:"time_in_force"`
	Trigger                Trigger         `json:"trigger"`
	Price                  decimal.Decimal `json:"price,omitempty"`
	StopPrice              decimal.Decimal `json:"stop_price,omitempty"`
	Quantity               uint            `json:"quantity"`
	Side                   Side            `json:"side"`
	ClientID               string          `json:"client_id,omitempty"`
	ExtendedHours          bool            `json:"extended_hours,omitempty"`
	OverrideDayTradeChecks bool            `json:"override_day_trade_checks,omitempty"`
	OverrideDTBPChecks     bool            `json:"override_dtbp_checks"`
}

type OrderState string

const (
	OrderQueued          OrderState = "queued"
	OrderUnconfirmed     OrderState = "unconfirmed"
	OrderConfirmed       OrderState = "confirmed"
	OrderPartiallyFilled OrderState = "partially_filled"
	OrderFilled          OrderState = "filled"
	OrderRejected        OrderState = "rejected"
	OrderCanceled        OrderState = "canceled"
)

type RHOrderResponse struct {
	ID                     string           `json:"id"`
	Executions             []string         `json:"executions"`
	Fees                   decimal.Decimal  `json:"fees"`
	Cancel                 string           `json:"cancel"`
	CumulativeQuantity     decimal.Decimal  `json:"cumulative_quantity"`
	RejectReason           string           `json:"reject_reason"`
	State                  OrderState       `json:"state"`
	ClientID               *string          `json:"client_id"`
	URL                    string           `json:"url"`
	Position               string           `json:"position"`
	AveragePrice           *decimal.Decimal `json:"average_price"`
	ExtendedHours          bool             `json:"extended_hours"`
	OverrideDayTradeChecks bool             `json:"override_day_trade_checks"`
	OverrideDTBPChecks     bool             `json:"override_dtbp_checks"`
	UpdatedAt              time.Time        `json:"updated_at"`
	CreatedAt              time.Time        `json:"created_at"`
	LastTranscationAt      time.Time        `json:"last_transaction_at"`
}

func (o *RHOrderResponse) fromMap(m map[string]interface{}) (err error) {
	o.ID = m["id"].(string)
	o.Executions = m["executions"].([]string)
	if fees, err := decimal.NewFromString(m["fees"].(string)); err != nil {
		return err
	} else {
		o.Fees = fees
	}
	o.Cancel = m["cancel"].(string)
	if cumQty, err := decimal.NewFromString(m["cumulative_quantity"].(string)); err != nil {
		return err
	} else {
		o.CumulativeQuantity = cumQty
	}
	o.RejectReason = m["reject_reason"].(string)
	o.State = OrderState(m["state"].(string))
	if m["client_id"] != nil {
		clientID := m["client_id"].(string)
		o.ClientID = &clientID
	}
	o.URL = m["url"].(string)
	o.Position = m["position"].(string)
	if m["average_price"] != nil {
		if avgPx, err := decimal.NewFromString(m["average_price"].(string)); err != nil {
			return err
		} else {
			o.AveragePrice = &avgPx
		}
	}
	o.ExtendedHours = m["extended_hours"].(bool)
	o.OverrideDayTradeChecks = m["override_day_trade_checks"].(bool)
	o.OverrideDTBPChecks = m["override_dtbp_checks"].(bool)
	// o.UpdatedAt, err = time.Parse(time.RFC3339, )
	return err
}

func (rh *Robinhood) Order(o RHOrder) (*RHOrderResponse, error) {
	uri := fmt.Sprintf("%v/orders/", baseURL)
	resp, err := rh.post(uri, o)
	if err != nil {
		return nil, err
	}
	oResp := &RHOrderResponse{}
	if err = json.Unmarshal(resp.Body(), oResp); err != nil {
		return nil, err
	}
	return oResp, nil
}

func (rh *Robinhood) GetOrder(id string) (*RHOrderResponse, error) {
	uri := fmt.Sprintf("%v/orders/%v", baseURL, id)
	resp, err := rh.get(uri)
	if err != nil {
		return nil, err
	}
	oResp := &RHOrderResponse{}
	if err = json.Unmarshal(resp.Body(), oResp); err != nil {
		return nil, err
	}
	return oResp, nil
}

type RHOrderQuery struct {
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	Instrument string     `json:"instrument,omitempty"`
	Cursor     string     `json:"cursor,omitempty"`
}

// func (rh *Robinhood) ListOrders(q RHOrderQuery) ([]RHOrderResponse, error) {
// 	uri := fmt.Sprintf("%v/orders/", baseURL)
// 	resp, err := rh.get(uri)
// 	if err != nil {
// 		return nil, err
// 	}
// 	ret := map[string]interface{}{}
// 	if err = json.Unmarshal(resp.Body(), &ret); err != nil {
// 		return nil, err
// 	}
// 	orders := []RHOrderResponse{}
// }
