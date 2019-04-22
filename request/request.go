package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/diskordanz/darksky/config"
	"github.com/pkg/errors"
	"github.com/xedinaska/int-weather-sdk/request"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RequestClient struct {
	HTTPClient HTTPClient
}

type apiRequest struct {
	ctx    context.Context
	method string
	path   string
	params url.Values
	data   interface{}
	target interface{}
}

func (r *RequestClient) Get(ctx context.Context, path string, params url.Values, target interface{}) (*request.Call, error) {
	return r.doRequest(&apiRequest{
		ctx:    ctx,
		method: http.MethodGet,
		path:   path,
		params: params,
		target: target,
	})
}

func (r *RequestClient) Post(ctx context.Context, path string, data interface{}, target interface{}) (*request.Call, error) {
	return r.doRequest(&apiRequest{
		ctx:    ctx,
		method: http.MethodPost,
		path:   path,
		data:   data,
		target: target,
	})
}

func (r *RequestClient) Delete(ctx context.Context, path string, params url.Values, target interface{}) (*request.Call, error) {
	return r.doRequest(&apiRequest{
		ctx:    ctx,
		method: http.MethodDelete,
		path:   path,
		target: target,
		params: params,
	})
}

func (r *RequestClient) doRequest(req *apiRequest) (*request.Call, error) {

	querystring := ""
	if req.params != nil {
		querystring = fmt.Sprintf("?%s", req.params.Encode())
	}

	// build full url
	baseURL := fmt.Sprintf("%s", os.Getenv(config.BaseURL))
	url := fmt.Sprintf("%s/%s%s", baseURL, req.path, querystring)

	// create json payload
	payload, err := json.Marshal(req.data)
	if err != nil {
		return nil, errors.Wrap(errors.WithStack(err), "failed to unmarshal request payload")
	}

	// create request object
	httpRequest, err := http.NewRequest(req.method, url, bytes.NewReader(payload))
	if err != nil {
		return nil, errors.Wrap(errors.WithStack(err), "failed to prepare request object")
	}

	//httpRequest.Header.Add("Authorization", fmt.Sprintf("TOKEN %s", req.ctx.Value(config.Token)))
	if req.data != nil {
		httpRequest.Header.Add("Content-Type", "application/json")
	}

	// time now
	start := time.Now()

	// send request
	response, err := r.HTTPClient.Do(httpRequest)
	if err != nil {
		return nil, errors.Wrap(errors.WithStack(err), "failed to make HTTP call")
	}

	var responseBody []byte
	if response.Body != nil {
		responseBody, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, errors.Wrap(errors.WithStack(err), "failed to read request body")
		}
	}

	defer response.Body.Close()

	call := &request.Call{
		RequestMethod:   response.Request.Method,
		RequestURL:      response.Request.URL.String(),
		RequestHeaders:  response.Request.Header,
		RequestBody:     string(payload),
		ResponseStatus:  response.StatusCode,
		ResponseHeaders: response.Header,
		ResponseBody:    string(responseBody),
		Took:            time.Since(start),
	}

	// decode response
	if response.StatusCode != 204 {
		err := json.Unmarshal([]byte(call.ResponseBody), req.target)
		if err != nil {
			return nil, errors.Wrap(errors.WithStack(err), "failed to unmarshal response body")
		}
	}

	return call, nil
}
