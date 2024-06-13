package response

const (
	ErrCodeSuccess       = 20001
	ErrCodeParamsInvalid = 20003
)

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamsInvalid: "field is invalid",
}
