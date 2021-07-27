package fireblockssdk

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type FireblocksSDK struct {
	options *options
	request request
}

type request struct {
	options *options
}

type options struct {
	privateKey string
	apiKey     string
	apiUrl     string
	hc         *http.Client
	ctx        context.Context
	timeout    time.Duration
}

func NewClient(optionFuncs ...func(*options)) FireblocksSDK {
	options := &options{
		apiUrl:  "https://api.fireblocks.io",
		hc:      &http.Client{},
		timeout: 5 * time.Second,
		ctx:     context.Background(),
	}

	for _, optionFunc := range optionFuncs {
		optionFunc(options)
	}
	return FireblocksSDK{
		options: options,
		request: request{
			options: options,
		},
	}
}

func WithPrivateKey(privateKey string) func(*options) {
	return func(o *options) {
		o.privateKey = privateKey
	}
}

func WithApiKey(apiKey string) func(*options) {
	return func(o *options) {
		o.apiKey = apiKey
	}
}

func WithApiUrl(apiUrl string) func(*options) {
	return func(o *options) {
		o.apiUrl = apiUrl
	}
}

func WithClient(hc *http.Client) func(*options) {
	return func(o *options) {
		o.hc = hc
	}
}

func WithTimeout(timeout time.Duration) func(*options) {
	return func(o *options) {
		o.timeout = timeout
	}
}

func WithContext(ctx context.Context) func(*options) {
	return func(o *options) {
		o.ctx = ctx
	}
}

func (r request) get(path string) ([]byte, error) {
	client := resty.NewWithClient(r.options.hc)
	client.SetTimeout(r.options.timeout)

	resp, err := client.R().
		SetContext(r.options.ctx).
		SetHeader("X-API-Key", r.options.apiKey).
		SetHeader("Authorization", "Bearer "+r.signJwt(path, "")).
		Get(r.options.apiUrl + path)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errResp map[string]interface{}
		if err = json.Unmarshal(resp.Body(), &errResp); err != nil {
			return nil, errors.New(string(resp.Body()))
		}

		return nil, errors.New(errResp["message"].(string))
	}

	return resp.Body(), nil
}

func (r request) post(path string, body interface{}, requestOptions *RequestOptions) ([]byte, error) {
	client := resty.NewWithClient(r.options.hc)
	client.SetTimeout(r.options.timeout)

	var idempotencyKey string
	if requestOptions != nil && requestOptions.IdempotencyKey != "" {
		idempotencyKey = requestOptions.IdempotencyKey
	}
	
	resp, err := client.R().
		SetContext(r.options.ctx).
		SetHeader("X-API-Key", r.options.apiKey).
		SetHeader("Authorization", "Bearer "+r.signJwt(path, body)).
		SetHeader("Idempotency-Key", idempotencyKey).
		SetBody(body).
		Post(r.options.apiUrl + path)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errResp map[string]interface{}
		if err = json.Unmarshal(resp.Body(), &errResp); err != nil {
			return nil, errors.New(string(resp.Body()))
		}

		return nil, errors.New(errResp["message"].(string))
	}

	return resp.Body(), nil
}

func (r request) put(path string, body interface{}) ([]byte, error) {
	client := resty.NewWithClient(r.options.hc)
	client.SetTimeout(r.options.timeout)

	resp, err := client.R().
		SetContext(r.options.ctx).
		SetHeader("X-API-Key", r.options.apiKey).
		SetHeader("Authorization", "Bearer "+r.signJwt(path, body)).
		SetBody(body).
		Put(r.options.apiUrl + path)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		var errResp map[string]interface{}
		if err = json.Unmarshal(resp.Body(), &errResp); err != nil {
			return nil, errors.New(string(resp.Body()))
		}

		return nil, errors.New(errResp["message"].(string))
	}

	return resp.Body(), nil
}
