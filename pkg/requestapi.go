package pkg

import (
	"belajar/efishery/utils/constants"
	utils "belajar/efishery/utils/response"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Info is the http req info
type ReqInfo struct {
	URL         string
	Method      string
	HeaderInfo  HeaderInfoSchema
	HeadersInfo map[string]interface{}
	Body        []byte
}

type HeaderInfoSchema struct {
	ContentType   string
	Auth          string
	XBRISignature string
	XBRITimestamp string
}

type ResInfo struct {
	StatusCode int
	Header     http.Header
	Body       []byte
}

// POST API request
func Invoke(reqinf *ReqInfo, timeout time.Duration) (*ResInfo, error) {
	var req *http.Request
	var err error

	switch reqinf.Method {
	case constants.MethodGET:
		req, err = http.NewRequest(constants.MethodGET, reqinf.URL, nil)
	case constants.MethodPOST:
		req, err = http.NewRequest(constants.MethodPOST, reqinf.URL, bytes.NewReader(reqinf.Body))
	}
	if err != nil {
		return nil, err
	}

	// set header
	for key, value := range reqinf.HeadersInfo {
		req.Header.Add(key, value.(string))
	}

	// execute
	cl := &http.Client{
		Timeout: timeout,
	}

	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var errornot200 error

	if res.StatusCode != http.StatusOK {
		errornot200 = utils.NewErr(fmt.Sprintf("%v for %v", res.StatusCode, reqinf.URL))
	}
	// read body
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &ResInfo{
		StatusCode: res.StatusCode,
		Body:       buf,
	}, errornot200
}
