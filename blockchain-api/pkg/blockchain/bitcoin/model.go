package bitcoin

type Body struct {
	ID       int64    `json:"id"`
	JsonRPC  string   `json:"jsonrpc"`
	Method   string   `json:"method"`
	Params   []string `json:"params"`
	Username string   `json:"username"`
	Password string   `json:"password"`
}
