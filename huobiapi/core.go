package huobiapi

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/daliyo/log4go/log4go"
)

var (
	accessKey = ""
	secretKey = ""
	host      = ""
)

const (
	// StatusOK StatusOK
	StatusOK = "ok"
)

// APIError 火币API错误
type APIError struct {
	API string
	Err error
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api %s, %s", e.API, e.Err)
}

// Init Init
func Init(accessKeyValue, secretKeyValue, hostValue string) {
	accessKey = accessKeyValue
	secretKey = secretKeyValue
	host = hostValue
}

type requestStruct struct {
	Path       string
	Method     string
	Sign       bool
	URLParams  url.Values
	FormParams url.Values
}

func (req *requestStruct) addURLParam(name, value string) {
	if req.URLParams == nil {
		req.URLParams = url.Values{}
	}
	req.URLParams.Add(name, value)
}

func (req *requestStruct) addFormParam(name, value string) {
	if req.FormParams == nil {
		req.FormParams = url.Values{}
	}
	req.FormParams.Add(name, value)
}

func buildRequestURL(path string, params url.Values) string {
	log4go.Debug("Building request URL")
	buff := strings.Builder{}
	buff.WriteString("https://")
	buff.WriteString(host)
	buff.WriteString(path)

	s := params.Encode()
	if s != "" {
		buff.WriteString("?")
		buff.WriteString(s)
	}

	URL := buff.String()
	log4go.Debug("Request URL successfully constructed", URL)

	return URL
}

func parseRespBody(res *http.Response, v interface{}) error {
	log4go.Debug("Request sent successfully")
	log4go.Debug("Parsing the response content")

	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log4go.Debug("Parsing failed! Unable to read response content")
		log4go.Error(&APIError{res.Request.URL.Path, err})
		return err
	}

	err = json.Unmarshal(buff, &v)
	if err != nil {
		log4go.Debug("Parsing failed! Can't deserialize response content")
		log4go.Error(&APIError{res.Request.URL.Path, err})
		return err
	}

	log4go.Debug("Response content parsed successfully")
	return nil
}

func sign(req *requestStruct) {
	log4go.Debug("Sign the request")
	buff := bytes.NewBufferString("")
	buff.WriteString(req.Method)
	buff.WriteString("\n")
	buff.WriteString(host)
	buff.WriteString("\n")
	buff.WriteString(req.Path)
	buff.WriteString("\n")

	req.addURLParam("AccessKeyId", accessKey)
	req.addURLParam("SignatureMethod", "HmacSHA256")
	req.addURLParam("SignatureVersion", "2")
	req.addURLParam("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05"))

	buff.WriteString(req.URLParams.Encode())

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Reset()
	h.Write(buff.Bytes())

	req.addURLParam("Signature", base64.StdEncoding.EncodeToString(h.Sum(nil)))
}

func sendRequest(req *requestStruct) (*http.Response, error) {

	if req.Sign {
		sign(req)
	}

	switch req.Method {
	case http.MethodGet:
		URL := buildRequestURL(req.Path, req.URLParams)
		return http.Get(URL)

	case http.MethodPost:
		URL := buildRequestURL(req.Path, req.URLParams)
		return http.PostForm(URL, req.FormParams)

	}
	return nil, errors.New("request method unsupported")
}
