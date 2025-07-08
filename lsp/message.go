package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	Id     int    `json:"id"`
	Method string `json:"method"`

	//Params later....
}

type Response struct {
	RPC string `json:"jsonrpc"`
	Id  *int   `json:"id|omitempty"`

	//Result
	//Error
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`

	//Params
}
