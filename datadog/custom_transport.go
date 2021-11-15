package datadog

// Credit to https://github.com/turbot/steampipe-plugin-datadog/blob/add-tables/internal/transport/custom_transport.go

import (
	"bytes"
	"context"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"time"
)

var (
	defaultHTTPRetryDuration = 5 * time.Second
	defaultHTTPRetryTimeout  = 60 * time.Second
	rateLimitResetHeader     = "X-Ratelimit-Reset"
)

// CustomTransport holds DefaultTransport configuration and is used to for custom http error handling
type CustomTransport struct {
	defaultTransport  http.RoundTripper
	httpRetryDuration time.Duration
	httpRetryTimeout  time.Duration
}

// CustomTransportOptions Set options for CustomTransport
type CustomTransportOptions struct {
	Timeout *time.Duration
}

// RoundTrip method used to retry http errors
func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var ccancel context.CancelFunc
	ctx := req.Context()
	if _, set := ctx.Deadline(); !set {
		ctx, ccancel = context.WithTimeout(ctx, t.httpRetryTimeout)
		defer ccancel()
	}

	retryCount := 0
	for {
		newRequest := t.copyRequest(req)
		resp, respErr := t.defaultTransport.RoundTrip(newRequest)
		// Close the body so connection can be re-used
		if resp != nil {
			localVarBody, _ := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
		}
		if respErr != nil {
			return resp, respErr
		}

		// Check if request should be retried and get retry time
		retryDuration, retry := t.retryRequest(resp)
		if !retry {
			return resp, respErr
		}

		// Calculate retryDuration if nil
		if retryDuration == nil {
			newRetryDurationVal := time.Duration(retryCount) * t.httpRetryDuration
			retryDuration = &newRetryDurationVal
		}

		select {
		case <-ctx.Done():
			return resp, respErr
		case <-time.After(*retryDuration):
			retryCount++
			continue
		}
	}
}

func (t *CustomTransport) copyRequest(r *http.Request) *http.Request {
	newRequest := *r

	if r.Body == nil || r.Body == http.NoBody {
		return &newRequest
	}

	body, _ := r.GetBody()
	newRequest.Body = body

	return &newRequest
}

func (t *CustomTransport) retryRequest(response *http.Response) (*time.Duration, bool) {
	if v := response.Header.Get(rateLimitResetHeader); v != "" && response.StatusCode == 429 {
		vInt, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, true
		}
		retryDuration := time.Duration(vInt) * time.Second
		return &retryDuration, true
	}

	if response.StatusCode >= 500 {
		return nil, true
	}

	return nil, false
}

// NewCustomTransport returns new CustomTransport struct
func NewCustomTransport(t http.RoundTripper, opt CustomTransportOptions) *CustomTransport {
	// Use default transport if one provided is nil
	if t == nil {
		t = http.DefaultTransport
	}

	ct := CustomTransport{
		defaultTransport:  t,
		httpRetryDuration: defaultHTTPRetryDuration,
	}

	if opt.Timeout != nil {
		ct.httpRetryTimeout = *opt.Timeout
	} else {
		ct.httpRetryTimeout = defaultHTTPRetryTimeout
	}

	return &ct
}

// It also tries to parse Retry-After response header when a http.StatusTooManyRequests
// (HTTP Code 429) is found in the resp parameter. Hence it will return the number of
// seconds the server states it may be ready to process more requests from this client.
func (t *CustomTransport) DefaultBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	if resp != nil {
		if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusServiceUnavailable {
			if s, ok := resp.Header["Retry-After"]; ok {
				if sleep, err := strconv.ParseInt(s[0], 10, 64); err == nil {
					return time.Second * time.Duration(sleep)
				}
			}
		}
	}

	mult := math.Pow(2, float64(attemptNum)) * float64(min)
	sleep := time.Duration(mult)
	if float64(sleep) != mult || sleep > max {
		sleep = max
	}
	return sleep
}
