package response

type Result struct {
	Data interface{}
	Code int64
	Msg  string
}
