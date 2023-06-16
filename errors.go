package qingflowapi

type ApiError struct {
	Code           string
	Message        string
	DefaultMessage string
}

func (e ApiError) Error() string {
	return e.Code + ":" + e.Message
}

func newApiError(code string, message string, defaultMessage string) ApiError {
	if len(message) == 0 {
		message = defaultMessage
	}
	var err ApiError
	err.Code = code
	err.Message = message
	err.DefaultMessage = defaultMessage
	return err
}

type ClientApiError struct {
	ApiError
}

type NotFoundApiError struct {
	ClientApiError
}

func notfound(defaultMessage string) func(string, string) error {
	return func(code string, message string) error {
		var err NotFoundApiError
		err.ApiError = newApiError(code, message, defaultMessage)
		return err
	}
}

type InvalidArgumentApiError struct {
	ClientApiError
}

func argumentInvalid(defaultMessage string) func(string, string) error {
	return func(code string, message string) error {
		var err InvalidArgumentApiError
		err.ApiError = newApiError(code, message, defaultMessage)
		return err
	}
}

type ServerApiError struct {
	ApiError
}

type InternalApiError struct {
	ServerApiError
}

func internal(defaultMessage string) func(string, string) error {
	return func(code string, message string) error {
		var err InvalidArgumentApiError
		err.ApiError = newApiError(code, message, defaultMessage)
		return err
	}
}

func translateError(code string, message string) error {
	convert, ok := codeMessage[code]
	if ok {
		return convert(code, message)
	}
	return ApiError{Code: code, Message: message}
}

type CodeMessage map[string]func(string, string) error

var codeMessage = CodeMessage{
	"40001": internal("登录信息失效，请重新登录"),
	"40011": notfound("用户不存在"),
}
