package message

type Message struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Object  interface{} `json:"object,omitempty"`
}

func Success() Message {
	return Message{
		Code:    0,
		Message: "ok",
	}
}

func SuccessWithObject(object interface{}) Message {
	return Message{
		Code:    0,
		Message: "ok",
		Object:  object,
	}
}

func Failed() Message {
	return Message{
		Code:    -1,
		Message: "failed",
	}
}

func FailedWithObject(object interface{}) Message {
	return Message{
		Code:    -1,
		Message: "failed",
		Object:  object,
	}
}
