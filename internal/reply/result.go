package reply

// Result ...
type Result struct {
	Code int16       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
