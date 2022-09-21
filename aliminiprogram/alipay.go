package aliminiprogram

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"git.myarena7.com/arena/platform/conf"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	defaultBaseURL = "https://openapi.alipay.com/gateway.do"
	userAgent      = ""
	timeLayout     = "2006-01-02 15:04:05"
)

// Options 公共请求参数
type Options struct {
	AppID      string // 支付宝分配给开发者的应用ID
	Format     string // 仅支持JSON 非必传
	Charset    string // 请求使用的编码格式，如utf-8,gbk,gb2312等
	SignType   string // 商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2
	Version    string // 调用的接口版本，固定为：1.0
	BizContent string // 请求参数的集合
	RootCertSN string // 根证书SN
	AppCertSN  string // 应用证书SN
}

// CommonResp 公共返回参数
type CommonResp struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

// Option 参数配置方法
type Option func(*Options)

// AppID 支付宝分配给开发者的应用ID
func AppID(appID string) Option {
	return func(o *Options) {
		o.AppID = appID
	}
}

// RootCertSN 根证书SN
func RootCertSN(rootCertSN string) Option {
	return func(o *Options) {
		o.RootCertSN = rootCertSN
	}
}

// AppCertSN 应用证书SN
func AppCertSN(appCertSN string) Option {
	return func(o *Options) {
		o.AppCertSN = appCertSN
	}
}

// ValueOptions 可选配置参数
type ValueOptions func(values url.Values)

// NotifyURL 支付宝服务器主动通知商户服务器里指定的页面http/https路径
func NotifyURL(notifyURL string) ValueOptions {
	return func(v url.Values) {
		v.Set("notify_url", notifyURL)
	}
}

// EncryptType 设置额外加密方式
func EncryptType(encrypt string) ValueOptions {
	return func(o url.Values) {
		o.Set("encrypt_type", encrypt)
	}
}

func EncryptKey(key string) ValueOptions {
	return func(o url.Values) {
		o.Set("encrypt_key", key)
	}
}

// AuthToken 针对用户授权接口，获取用户相关数据时，用于标识用户授权关系
func AuthToken(authToken string) ValueOptions {
	return func(o url.Values) {
		o.Set("auth_token", authToken)
	}
}

// AppAuthToken 第三方应用授权
func AppAuthToken(AppAuthToken string) ValueOptions {
	return func(o url.Values) {
		o.Set("app_auth_token", AppAuthToken)
	}
}

// Format 仅支持JSON
func Format(format string) Option {
	return func(o *Options) {
		o.Format = format
	}
}

// Charset 请求使用的编码格式，如utf-8,gbk,gb2312等
func Charset(charset string) Option {
	return func(o *Options) {
		o.Charset = charset
	}
}

// SignType 商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA
func SignType(signType string) Option {
	return func(o *Options) {
		o.SignType = signType
	}
}

// A Client manages communication with the Alipay API.
type Client struct {
	clientMu sync.Mutex   // clientMu protects the client during calls that modify the CheckRedirect func.
	client   *http.Client // HTTP client used to communicate with the API.

	// Base URL for API requests. Defaults to the public Alipay API, but can be
	// set to a domain endpoint to use with GitHub Enterprise. BaseURL should
	// always be specified with a trailing slash.
	BaseURL *url.URL

	o *Options

	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey

	// User agent used when communicating with the Alipay API.
	UserAgent string

	RootCertSN string
	AppCertSN  string
	AppAesKey  string

	common Service // Reuse a single struct instead of allocating one for each service on the heap.

	App                *AppService         // 小程序
	Mini               *MiniService        // 应用服务
	CultureSportCenter *CultureSportCenter // 文体中心
	OAuth              *OAuth              // auth
}

type Service struct {
	Client *Client
}

// NewClient returns a new Alipay API client. If a nil httpClient is
// provided, a new http.Client will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, setters ...Option) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	options := &Options{
		Format:   "JSON",
		Charset:  "utf-8",
		SignType: "RSA2",
		Version:  "1.0",
	}
	for _, setter := range setters {
		setter(options)
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:     httpClient,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		RootCertSN: options.RootCertSN,
		AppCertSN:  options.AppCertSN,
		BaseURL:    baseURL,
		UserAgent:  userAgent,
		o:          options,
	}
	c.common.Client = c
	c.App = (*AppService)(&c.common)
	c.Mini = (*MiniService)(&c.common)
	c.CultureSportCenter = (*CultureSportCenter)(&c.common)
	c.OAuth = (*OAuth)(&c.common)
	return c
}

// File wrapped file content
type File struct {
	Name    string
	Content io.Reader
}

// Read proxy Content Read
func (f File) Read(p []byte) (n int, err error) {
	return f.Content.Read(p)
}

func (c *Client) NewGetRequest(method string, v url.Values, setters ...ValueOptions) (*http.Request, error) {
	var (
		sign        string
		contentType = "application/x-www-form-urlencoded"
		req         *http.Request
		err         error
	)
	v.Set("app_id", c.o.AppID)
	v.Set("method", method)
	v.Set("format", c.o.Format)
	v.Set("charset", c.o.Charset)
	v.Set("sign_type", c.o.SignType)
	v.Set("timestamp", time.Now().Format(timeLayout))
	v.Set("version", c.o.Version)
	v.Set("app_cert_sn", c.o.AppCertSN)
	v.Set("alipay_root_cert_sn", c.o.RootCertSN)
	for _, setter := range setters {
		setter(v)
	}
	sign, err = c.Sign(v)
	if err != nil {
		return nil, err
	}
	v.Set("sign", sign)
	c.BaseURL.RawQuery = v.Encode()

	req, err = http.NewRequest("GET", c.BaseURL.String(), nil)
	if err != nil {
		return nil, err
	}
	v = req.URL.Query()
	v.Set("charset", c.o.Charset)
	req.URL.RawQuery = v.Encode()
	req.Header.Set("Content-Type", contentType)

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method string, bizContent interface{}, setters ...ValueOptions) (*http.Request, error) {
	var (
		sign        string
		contentType = "application/x-www-form-urlencoded"
		buf         *bytes.Buffer
		req         *http.Request
		reader      io.Reader
		err         error
	)
	v := url.Values{}
	v.Set("app_id", c.o.AppID)
	v.Set("method", method)
	v.Set("format", c.o.Format)
	v.Set("charset", c.o.Charset)
	v.Set("sign_type", c.o.SignType)
	v.Set("timestamp", time.Now().Format(timeLayout))
	v.Set("version", c.o.Version)
	v.Set("app_cert_sn", c.o.AppCertSN)
	v.Set("alipay_root_cert_sn", c.o.RootCertSN)
	for _, setter := range setters {
		setter(v)
	}

	if bizContent != nil {
		render, ok := bizContent.(MultiRender)
		if ok {
			var b bytes.Buffer
			w := multipart.NewWriter(&b)
			for key, r := range render.MultipartParams() {
				var fw io.Writer
				if x, ok := r.(*File); ok {
					if fw, err = w.CreateFormFile(key, x.Name); err != nil {
						return nil, err
					}
				}
				if _, err = io.Copy(fw, r); err != nil {
					return nil, err
				}
			}
			params := render.Params()
			for key, val := range params {
				v.Set(key, val)
			}
			sign, err = c.Sign(v)
			if err != nil {
				return nil, err
			}
			v.Set("sign", sign)
			for k := range v {
				_ = w.WriteField(k, v.Get(k))
			}
			err = w.Close()
			if err != nil {
				return nil, err
			}
			reader = &b
			contentType = w.FormDataContentType()
		} else {
			bizNewContent := bizContent
			if v.Get("encrypt_type") != "" && v.Get("encrypt_key") != "" {
				bizContentByte, err := json.Marshal(bizContent)
				if err != nil {
					return nil, err
				}
				keyByte, err := base64.StdEncoding.DecodeString(v.Get("encrypt_key"))
				if err != nil {
					return nil, err
				}
				aesBizContent, err := EncryptByAes(bizContentByte, keyByte)
				if err != nil {
					return nil, err
				}
				bizNewContent = aesBizContent
			}
			buf = &bytes.Buffer{}
			enc := json.NewEncoder(buf)
			enc.SetEscapeHTML(false)
			err = enc.Encode(bizNewContent)
			if err != nil {
				return nil, err
			}
			v.Set("biz_content", buf.String())
			sign, err = c.Sign(v)
			if err != nil {
				return nil, err
			}
			v.Set("sign", sign)
			reader = strings.NewReader(v.Encode())
		}
	} else {
		sign, err = c.Sign(v)
		if err != nil {
			return nil, err
		}
		v.Set("sign", sign)
		reader = strings.NewReader(v.Encode())
	}

	req, err = http.NewRequest("POST", c.BaseURL.String(), reader)
	if err != nil {
		return nil, err
	}
	v = req.URL.Query()
	v.Set("charset", c.o.Charset)
	req.URL.RawQuery = v.Encode()
	req.Header.Set("Content-Type", contentType)

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Sign 参数签名
func (c *Client) Sign(values url.Values) (string, error) {
	if c.PrivateKey == nil {
		return "", nil
	}
	var buf strings.Builder
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := values[k]
		for _, v := range vs {
			if v == "" {
				continue
			}
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	valuesStr := buf.String()
	signType := crypto.SHA256
	if c.o.SignType == "RSA" {
		signType = crypto.SHA1
	}
	h := crypto.Hash.New(signType)
	h.Write([]byte(valuesStr))

	signature, err := rsa.SignPKCS1v15(rand.Reader, c.PrivateKey, signType, h.Sum(nil))
	if err != nil {
		return "", err
	}
	sign := base64.StdEncoding.EncodeToString(signature)
	return sign, nil
}

// ErrorResponse is common error response.
type ErrorResponse struct {
	*http.Response
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%d %v %+v, %v %+v",
		r.Response.StatusCode, r.Msg, r.Code, r.SubCode, r.SubMsg)
}

func withContext(ctx context.Context, req *http.Request) *http.Request {
	return req.WithContext(ctx)
}

// Response is a Alipay API response.
type Response struct {
	*http.Response
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it. If rate limit is exceeded and reset time is in the future,
// Do returns *RateLimitError immediately without making a network API call.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}
	req = withContext(ctx, req)
	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{resp}

	if strings.ToUpper(req.Method) == "GET" {
		err = c.CheckResponse(resp, req.URL.Query().Get("method"))
		if err != nil {
			return response, err
		}
	} else {
		err = c.CheckResponse(resp, req.URL.Query().Get("method"))
		if err != nil {
			return response, err
		}
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}
	return response, err
}

// CheckResponse 检查返回内容
func (c *Client) CheckResponse(r *http.Response, requestMethod string) error {
	errorResponse := &ErrorResponse{Response: r}
	var backData []byte
	data, err := ioutil.ReadAll(r.Body)
	var resp, sign []byte
	if err == nil && data != nil {
		obj := make(map[string]json.RawMessage)
		if err = json.Unmarshal(data, &obj); err != nil {
			return err
		}

		var responseKey string
		for k, v := range obj {
			if strings.Contains(k, "response") {
				resp = v
				responseKey = k
				break
			}
		}
		sign = obj["sign"]
		if len(sign) > 0 {
			var signStr string
			if err = json.Unmarshal(sign, &signStr); err != nil {
				return fmt.Errorf("反序列化签名失败: %w", err)
			}
			if err = c.VerifySign(resp, signStr); err != nil {
				return fmt.Errorf("支付宝同步请求签名验证不通过: %w", err)
			}
		}
		if !strings.HasPrefix(string(resp), "{") { // 加密返回
			// TODO 目前就一个三方应用直接写死 多个的时候直接初始化client好
			if conf.GetAliMiniProgramConf().AppAesKey == "" {
				return fmt.Errorf("AESKEY未配置")
			}
			keyByte, err := base64.StdEncoding.DecodeString(conf.GetAliMiniProgramConf().AppAesKey)
			if err != nil {
				return err
			}
			dataSlice, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(string(resp), "\"", ""))
			if err != nil {
				return err
			}
			backSlice, err := AesCBCDecrypt(dataSlice, keyByte)
			if err != nil {
				return err
			}
			if err = json.Unmarshal(backSlice, &errorResponse); err != nil {
				return fmt.Errorf("解析支付宝返回结构失败: %w", err)
			}
			obj[responseKey] = backSlice
			byteData, err := json.Marshal(obj)
			if err != nil {
				return err
			}
			backData = byteData
		} else {
			if err = json.Unmarshal(resp, &errorResponse); err != nil {
				return fmt.Errorf("解析支付宝返回结构失败: %w", err)
			}
		}
	}
	if len(backData) == 0 {
		backData = data
	}
	buf := bytes.NewBuffer(backData)
	if requestMethod == "alipay.system.oauth.token" {
		if errorResponse.Code == "" && errorResponse.SubCode == "" && errorResponse.SubMsg == "" {
			r.Body = ioutil.NopCloser(buf)
			return nil
		} else {
			return errorResponse
		}
	} else {
		if errorResponse.Code == "10000" {
			r.Body = ioutil.NopCloser(buf)
			return nil
		}
	}
	return errorResponse
}

// VerifySign 校验同步请求返回参数
func (c *Client) VerifySign(content []byte, sign string) error {
	signData, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	signType := crypto.SHA256
	if c.o.SignType == "RSA" {
		signType = crypto.SHA1
	}
	h := crypto.Hash.New(signType)
	h.Write(content)
	return rsa.VerifyPKCS1v15(c.PublicKey, signType, h.Sum(nil), signData)
}
