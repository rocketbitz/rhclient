package rhclient

import (
	"encoding/json"
	"sync"

	"github.com/valyala/fasthttp"
)

var once sync.Once
var rh *Robinhood

var baseURL = "https://api.robinhood.com"

type Robinhood struct {
	request func(req *fasthttp.Request, resp *fasthttp.Response)
	token   string
}

func Client() *Robinhood {
	once.Do(func() {
		rh = &Robinhood{}
		rh.request = func(req *fasthttp.Request, resp *fasthttp.Response) {
			if len(rh.token) > 0 {
				req.Header.Set("Authorization", rh.token)
			}
			req.Header.SetContentType("application/json")
			req.Header.Set("Accept", "application/json")
			(&fasthttp.Client{}).Do(req, resp)
		}
	})
	return rh
}

func (rh *Robinhood) post(uri string, payload interface{}) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(uri)
	req.Header.SetMethod("POST")
	if payload != nil {
		if body, err := json.Marshal(payload); err != nil {
			return nil, err
		} else {
			req.SetBody(body)
		}
	}
	resp := fasthttp.AcquireResponse()
	rh.request(req, resp)
	return resp, nil
}

func (rh *Robinhood) get(uri string) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(uri)
	req.Header.SetMethod("GET")
	resp := fasthttp.AcquireResponse()
	rh.request(req, resp)
	return resp, nil
}
