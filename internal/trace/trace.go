package trace

import "sync"

var _ T = (*Trace)(nil)

type T interface {
	TraceId() string
	WithRequest(req *Request) *Trace
	WithResponse(rsp *Response) *Trace
	AppendDialog(dialog *Dialog) *Trace
	AppendDebug(debug *Debug) *Trace
	AppendSQL(sql *SQL) *Trace
	AppendRedis(redis *Redis) *Trace
}

// Trace 记录的参数
type Trace struct {
	mux                sync.Mutex
	TranceId           string    `json:"trace_id"`             // 链路 ID
	Request            *Request  `json:"request"`              // 请求信息
	Response           *Response `json:"response"`             // 响应信息
	ThirdPartyRequests []*Dialog `json:"third_party_requests"` // 调用第三方接口的信息
	Debugs             []*Debug  `json:"debugs"`               // 调试信息
	SQLs               []*SQL    `json:"sqls"`                 // 执行的 SQL 信息
	Redis              []*Redis  `json:"redis"`                // 执行的 Redis 信息
	Success            bool      `json:"success"`              // 请求结果 true or false
	CostSeconds        float64   `json:"cost_seconds"`         // 执行时长(单位秒)
}

// Request 请求信息
type Request struct {
	TTL        string      `json:"ttl"`         // 请求超时时间
	Method     string      `json:"method"`      // 请求方式
	DecodedURL string      `json:"decoded_url"` // 请求地址
	Header     interface{} `json:"header"`      // 请求 Header 信息
	Body       interface{} `json:"body"`        // 请求 Body 信息
}

// Response 相应信息
type Response struct {
	Header          interface{} `json:"header"`                      // Header 信息
	Body            interface{} `json:"body"`                        // Body 信息
	BusinessCode    int         `json:"business_code,omitempty"`     // 业务码
	BusinessCodeMsg string      `json:"business_code_msg,omitempty"` // 提示信息
	HttpCode        int         `json:"http_code"`                   // HTTP 状态码
	HttpCodeMsg     string      `json:"http_code_msg"`               // HTTP 状态码信息
	CostSeconds     float64     `json:"cost_seconds"`                // 执行时间(单位秒)
}

// Dialog 调用第三方接口信息
type Dialog struct {
	mux         sync.Mutex
	Request     *Request    `json:"request"`      // 请求信息
	Responses   []*Response `json:"responses"`    // 返回信息
	Success     bool        `json:"success"`      // 是否成功，true 或 false
	CostSeconds float64     `json:"cost_seconds"` // 执行时长(单位秒)
}

// Debug 调试信息
type Debug struct {
	Key         string      `json:"key"`          // 标示
	Value       interface{} `json:"value"`        // 值
	CostSeconds float64     `json:"cost_seconds"` // 执行时间(单位秒)
}

// SQL SQL信息
type SQL struct {
	Timestamp   string  `json:"timestamp"`     // 时间，格式：2006-01-02 15:04:05
	Stack       string  `json:"stack"`         // 文件地址和行号
	SQL         string  `json:"sql"`           // SQL 语句
	Rows        int64   `json:"rows_affected"` // 影响行数
	CostSeconds float64 `json:"cost_seconds"`  // 执行时长(单位秒)
}

// Redis Redis信息
type Redis struct {
	Timestamp   string  `json:"timestamp"`       // 时间，格式：2006-01-02 15:04:05
	Handle      string  `json:"handle"`          // 操作，SET/GET 等
	Key         string  `json:"key"`             // Key
	Value       string  `json:"value,omitempty"` // Value
	TTL         float64 `json:"ttl,omitempty"`   // 超时时长(单位分)
	CostSeconds float64 `json:"cost_seconds"`    // 执行时间(单位秒)
}

func NewTrance(id string) *Trace {
	instance := &Trace{
		TranceId: id,
	}
	return instance
}

func (t *Trace) TraceId() string {
	return t.TranceId
}

func (t *Trace) WithRequest(req *Request) *Trace {
	t.Request = req
	return t
}

func (t *Trace) WithResponse(rsp *Response) *Trace {
	t.Response = rsp
	return t
}
func (t *Trace) AppendDialog(dialog *Dialog) *Trace {
	t.mux.Lock()
	defer t.mux.Unlock()
	t.ThirdPartyRequests = append(t.ThirdPartyRequests, dialog)
	return t
}

func (t *Trace) AppendDebug(debug *Debug) *Trace {
	t.mux.Lock()
	defer t.mux.Unlock()
	t.Debugs = append(t.Debugs, debug)
	return t
}

func (t *Trace) AppendSQL(sql *SQL) *Trace {
	t.mux.Lock()
	defer t.mux.Unlock()
	t.SQLs = append(t.SQLs, sql)
	return t
}

func (t *Trace) AppendRedis(redis *Redis) *Trace {
	t.mux.Lock()
	defer t.mux.Unlock()
	t.Redis = append(t.Redis, redis)
	return t
}
