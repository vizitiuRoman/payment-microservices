package bitcoin

type Body struct {
	Id       int64    `json:"id"`
	JsonRpc  string   `json:"jsonrpc"`
	Method   string   `json:"method"`
	Params   []string `json:"params"`
	Username string   `json:"username"`
	Password string   `json:"password"`
}
