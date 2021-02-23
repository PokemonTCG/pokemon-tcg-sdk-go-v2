package request

import "net/url"

// A Request represents a request to the API.
type Request struct {
	endpoint string
	options  map[string]string
}

// New constructs a request with reasonable defaults and applies any
// options provided.
func New(endpoint string, options ...Option) *Request {
	r := &Request{endpoint, make(map[string]string)}
	for _, o := range options {
		o(r)
	}
	return r
}

// GetURL creates the URL from the internal endpoint and options.
func (r *Request) GetURL() (string, error) {
	u, err := url.Parse(r.endpoint)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for k, v := range r.options {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}
