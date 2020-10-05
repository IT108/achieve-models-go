package achieve_models_go

const (
	AUTH_REGISTER_KEY = "register"
)

type RegisterRequest struct {
	Request
	Email    string
	Username string
	Password string
}

type RegisterResponse struct {
	Request
	ResponseCode int
	Error        string
}

type AuthenticateRequest struct {
	Request
	Username string
	Password string
}

type AuthenticateResponse struct {
	Request
	ResponseCode int
	Error        string
}

type AuthorizeRequest struct {
	Request
	Username string
}

type AuthorizeResponse struct {
	Request
	ResponseCode int
	roles        []string
	Error        string
}

//func (receiver *RegisterRequest) AcceptReq(kafkaMsg *kafka.Message) {
//	err := json.Unmarshal(kafkaMsg.Value, receiver)
//	if err != nil {
//		panic(err)
//	}
//}
