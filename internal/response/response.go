package response

type Response struct {
	Message string `json:"message,omitempty"`
	Status  string `json:"status"`
}

func Error(message string) Response {
	return Response{Message: message, Status: STATUS_ERROR}
}

func Ok() Response {
	return Response{Status: STATUS_OK}
}

const STATUS_OK = "ok"
const STATUS_ERROR = "error"
