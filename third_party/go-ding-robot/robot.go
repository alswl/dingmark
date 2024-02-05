package robot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/alswl/dingmark/third_party/go-ding-robot/response"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/alswl/dingmark/third_party/go-ding-robot/message"
)

const (
	Timeout = 5 * time.Second
	origin  = "https://oapi.dingtalk.com"
)

// Webhook writable for CORS proxy
var Webhook = "https://oapi.dingtalk.com/robot/send"
var ExtendHeaders = map[string]string{}

type Robot struct {
	Token, Secret string
	Client        *http.Client
}

func New(token, secret string) *Robot {
	return &Robot{
		Token:  token,
		Secret: secret,
	}
}

func (bot *Robot) Send(message message.IMessage) (*response.Response, error) {
	var (
		data []byte
		err  error
	)
	if data, err = json.Marshal(message); err != nil {
		return nil, err
	}

	var result []byte
	if result, err = bot.send(data); err != nil {
		return nil, err
	}
	var resp = new(response.Response)
	if err = json.Unmarshal(result, &resp); err != nil {
		return nil, err
	}
	if resp.Status > 0 {
		resp.Code = resp.Status
		resp.Message = fmt.Sprintf("punish[%s] wait[%d]", resp.Punish, resp.Wait)
	}

	return resp, nil
}

// http request
func (bot *Robot) send(data []byte) ([]byte, error) {
	var (
		req *http.Request
		err error
	)
	if req, err = bot.buildRequest(data, http.Header{}); err != nil {
		return nil, err
	}

	if bot.Client == nil {
		bot.setDefaultClient()
	}
	var resp *http.Response
	if resp, err = bot.Client.Do(req); err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error http code: %d", resp.StatusCode)
	}

	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	return body, err
}

// build http post request
func (bot *Robot) buildRequest(data []byte, header http.Header) (*http.Request, error) {
	var (
		url *url.URL
		err error
	)
	if url, err = bot.buildWebHook(); err != nil {
		return nil, err
	}

	var req *http.Request
	if req, err = http.NewRequest("POST", url.String(), bytes.NewReader(data)); err != nil {
		return nil, err
	}

	header.Set("Content-Type", "application/json;charset=utf-8")
	for k, v := range ExtendHeaders {
		header.Set(k, v)
	}
	req.Header = header

	return req, nil
}

// build robot webhook
func (bot *Robot) buildWebHook() (*url.URL, error) {
	var (
		url *url.URL
		err error
	)
	if url, err = url.Parse(Webhook); err != nil {
		return nil, err
	}

	query := url.Query()
	query.Set("access_token", bot.Token)
	if bot.Secret != "" {
		timestamp, sign := bot.sign()
		query.Add("timestamp", strconv.FormatInt(timestamp, 10))
		query.Add("sign", sign)
	}
	url.RawQuery = query.Encode()

	return url, nil
}

// default http client
func (bot *Robot) setDefaultClient() *Robot {
	bot.Client = &http.Client{Timeout: Timeout}

	return bot
}

// signature
func (bot *Robot) sign() (int64, string) {
	timestamp := time.Now().Unix() * 1000
	src := fmt.Sprintf("%d\n%s", timestamp, bot.Secret)

	sha256 := hmac.New(sha256.New, []byte(bot.Secret))
	sha256.Write([]byte(src))

	return timestamp, base64.StdEncoding.EncodeToString(sha256.Sum(nil))
}
